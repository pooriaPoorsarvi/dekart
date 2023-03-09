




gcloud auth login --cred-file=$GOOGLE_APPLICATION_CREDENTIALS
gcloud config set project ppt-personal-project

./cloud-sql-proxy ppt-personal-project:us-central1:dekart & ./server

















