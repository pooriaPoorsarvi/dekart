

docker tag ${CUSTOM_DEKART_IMAGE_NAME} ${CUSTOM_DEKART_IMAGE_NAME}
docker push ${CUSTOM_DEKART_IMAGE_NAME}


./custom.sentry.deploy.sh
./custom.app.deploy.sh

