docker run -it --rm  \
		-p 8080:8080 \
		-v ${GOOGLE_APPLICATION_CREDENTIALS_FOLDER}:/root/.config/gcloud \
		-e GOOGLE_APPLICATION_CREDENTIALS=/root/.config/gcloud/application_default_credentials.json \
		-e DEKART_POSTGRES_DB=${DEKART_POSTGRES_DB} \
		-e DEKART_POSTGRES_USER=${DEKART_POSTGRES_USER} \
		-e DEKART_POSTGRES_PASSWORD=${DEKART_POSTGRES_PASSWORD} \
		-e DEKART_POSTGRES_PORT=${DEKART_POSTGRES_PORT} \
		-e DEKART_POSTGRES_HOST=localhost \
		-e DEKART_CLOUD_STORAGE_BUCKET=${DEKART_CLOUD_STORAGE_BUCKET} \
		-e DEKART_BIGQUERY_PROJECT_ID=${DEKART_BIGQUERY_PROJECT_ID} \
		-e DEKART_BIGQUERY_MAX_BYTES_BILLED=53687091200 \
		-e DEKART_MAPBOX_TOKEN=${DEKART_MAPBOX_TOKEN} \
        -e INSTANCE_CONNECTION_NAME=${INSTANCE_CONNECTION_NAME} \
		-e DEKART_STATIC_FILES=${DEKART_STATIC_FILES} \
		-e DEKART_PORT=${DEKART_PORT} \
		${CUSTOM_DEKART_IMAGE_NAME}