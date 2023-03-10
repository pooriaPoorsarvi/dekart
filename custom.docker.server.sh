



/dekart/cloud-sql-proxy ${INSTANCE_CONNECTION_NAME} > output.txt &
timeout=10
while ! grep -q "The proxy has started" output.txt; do
  sleep 1
  timeout=$((timeout - 1))
  if [ $timeout -eq 0 ]; then
    echo "Error: postgres could not be reached"
    exit 1
  fi
done
/dekart/server



















