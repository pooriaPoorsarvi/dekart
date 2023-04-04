package user

import (
	"context"
	"dekart/src/server/utils"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	GcpOauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2/google"
)


func getEmail() (string, error) {
	keyFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if keyFile == "" {
		fmt.Println("GOOGLE_APPLICATION_CREDENTIALS environment variable not set")
		return "", nil
	}

	keyData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Debug().Msg("Error reading key file")
		log.Err(err)
		return "", err
	}


	ctx := context.Background()
	creds, err := google.CredentialsFromJSON(ctx, keyData)
	if err != nil {
		log.Debug().Msg("Error creating credentials from JSON")
		log.Err(err)
		return "", err
	}

	tok, err := creds.TokenSource.Token()
	if !tok.Valid() {
		err = errors.New("token is not valid")
		log.Debug().Msg("Error validating the token")
		log.Err(err)
		return "", err
	}


	client := utils.OauthConf.Client(ctx, tok)
	service, err := GcpOauth.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Debug().Msg("Error while getting config")
		log.Err(err)
		return "", err
	}
	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		log.Debug().Msg("Error while getting user info")
		log.Err(err)
		return "", err
	}

	return userInfo.Email, nil

}