package dekart

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/bigquery/v2"
	GcpOauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"time"
)

func (s Server) SaveToken(ctx context.Context, code string, state string) error{

	conf := &oauth2.Config{
		ClientID:     "197398309945-ostmrt571il6vtgvd0mdceaccmhdmji8.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{bigquery.BigqueryScope, GcpOauth.UserinfoProfileScope, GcpOauth.UserinfoEmailScope},
		Endpoint: google.Endpoint,
		RedirectURL: "http://localhost:8080/api/v1//callback-authenticate-oauth2",

	}


	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Info().Err(err)
		return err
	} else {


		client := conf.Client(ctx, tok)
		service, err := GcpOauth.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Info().Msgf("err 3")
			log.Info().Msgf(err.Error())
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

	}

	return nil
}

func (s Server) RetrieveToken(userEmail string) (*oauth2.Token, error) {
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


	token := new(oauth2.Token)
	token.AccessToken = userToken.AccessToken
	token.RefreshToken = userToken.RefreshToken
	token.Expiry = userToken.Expiry
	token.TokenType = userToken.TokenType

	return token, err


}