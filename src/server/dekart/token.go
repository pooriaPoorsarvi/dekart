package dekart

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/bigquery/v2"
	GcpOauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"strings"
	"time"
)


func (s Server) checkTokenScopes(conf *oauth2.Config, service *GcpOauth.Service, tok *oauth2.Token) (error){
	tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Do()
	if err != nil {
		log.Info().Msgf("Failed to check token info: %v", err)
		return err
	}
	grantedScopes := strings.Split(tokenInfo.Scope, " ")
	for _, scope := range conf.Scopes {
		if !contains(grantedScopes, scope) {
			err := fmt.Errorf("scope '%s' not granted", scope)
			log.Info().Err(err)
			return err
		}
	}
	return nil
}

func (s Server) SaveToken(code string, state string) error {

	conf := &oauth2.Config{
		ClientID:     "197398309945-ostmrt571il6vtgvd0mdceaccmhdmji8.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{bigquery.BigqueryScope, GcpOauth.UserinfoProfileScope, GcpOauth.UserinfoEmailScope},
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/api/v1/callback-authenticate-oauth2",
	}
	ctx := context.Background()
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Info().Err(err)
		return err
	}

	// Check if all the requested scopes have been granted
	client := conf.Client(ctx, tok)
	service, err := GcpOauth.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Info().Msgf("err 3")
		log.Info().Msgf(err.Error())
		return err
	}


	err = s.checkTokenScopes(conf, service, tok)
	if err != nil{
		return err
	}


	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		log.Print("got user error", err.Error())
		return err
	}


	sqlStatement := `INSERT INTO user_token (id, access_token, refresh_token, expiry, token_type)
                     VALUES ($1, $2, $3, $4, $5)`
	_, err = s.db.Exec(sqlStatement, userInfo.Email, tok.AccessToken, tok.RefreshToken, tok.Expiry, tok.TokenType)
	if err != nil {
		log.Print("error saving token")
		log.Info().Msgf(err.Error())
	}

	log.Info().Msgf("Fully saved token")

	return nil
}


func contains(scopes []string, scope string) bool {
	for _, s := range scopes {
		if s == scope {
			return true
		}
	}
	return false
}
func (s Server) RetrieveToken(userEmail string) (*oauth2.Token, error) {


	ctx := context.Background()
	// Create an OAuth2 configuration with the same values as the one used to obtain the access token
	conf := &oauth2.Config{
		ClientID:     "197398309945-ostmrt571il6vtgvd0mdceaccmhdmji8.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{bigquery.BigqueryScope, GcpOauth.UserinfoProfileScope, GcpOauth.UserinfoEmailScope},
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/api/v1//callback-authenticate-oauth2",
	}

	type UserToken struct {
		ID           string
		AccessToken  string
		RefreshToken string
		Expiry       time.Time
		TokenType    string
	}
	var userToken UserToken

	sqlStatement := `SELECT * FROM user_token WHERE id=$1`
	row := s.db.QueryRow(sqlStatement, userEmail)
	err := row.Scan(&userToken.ID, &userToken.AccessToken, &userToken.RefreshToken, &userToken.Expiry, &userToken.TokenType)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info().Msgf("UserToken not found")
			return nil, err
		}
		log.Info().Msgf(err.Error())
	}



	// Use the refresh token to obtain a new access token
	var newToken * oauth2.Token
	remainingTime := userToken.Expiry.Sub(time.Now())
	if remainingTime < 1*time.Hour {
		tokenSource := conf.TokenSource(ctx, &oauth2.Token{RefreshToken: userToken.RefreshToken})
		newToken, err = tokenSource.Token()
		if err != nil {
			log.Warn().Msgf("Failed to refresh token: %v", err)
			s.deleteToken(userEmail) // remove the invalid token from the database
			return nil, err
		}

		sqlStatement = `UPDATE user_token SET access_token=$1, expiry=$2 WHERE id=$3`
		_, err = s.db.Exec(sqlStatement, newToken.AccessToken, newToken.Expiry, userEmail)
		if err != nil {
			log.Info().Msgf("Failed to update token in database: %v", err)
			s.deleteToken(userEmail) // remove the invalid token from the database
			return nil, err
		}

	}else{
		newToken = &oauth2.Token{
			AccessToken:  userToken.AccessToken,
			RefreshToken: userToken.RefreshToken,
			Expiry:       userToken.Expiry,
			TokenType:    userToken.TokenType,
		}
	}

	// Update the access token and expiry time in the database

	client := conf.Client(ctx, newToken)
	service, err := GcpOauth.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Info().Msgf("Failed to check token info: %v", err)
		s.deleteToken(userEmail) // remove the invalid token from the database
		return nil, err
	}
	err = s.checkTokenScopes(conf, service, newToken)
	if err != nil{
		log.Err(err)
		s.deleteToken(userEmail) // remove the invalid token from the database
		return nil, err
	}

	return newToken, nil
	// If the saved token is valid and contains all requested scopes, return it
}

func (s Server) deleteToken(userEmail string) {
	sqlStatement := `DELETE FROM user_token WHERE id=$1`
	_, err := s.db.Exec(sqlStatement, userEmail)
	if err != nil {
		log.Info().Msgf("Failed to delete token from database: %v", err)
	}
}
