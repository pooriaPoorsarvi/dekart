package dekart

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/bigquery/v2"
	GcpOauth "google.golang.org/api/oauth2/v2"
	"os"
)

var OauthConf = &oauth2.Config{
	ClientID:     os.Getenv("OauthClientId"),
	ClientSecret: os.Getenv("OauthClientSecret"),
	Scopes:       []string{bigquery.BigqueryScope, GcpOauth.UserinfoProfileScope, GcpOauth.UserinfoEmailScope},
	Endpoint: google.Endpoint,
	RedirectURL: os.Getenv("OauthRedirectUrl"),
}






