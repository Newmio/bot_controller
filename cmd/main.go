package main

import "bot/internal/app"

/*
docker run -d \
  --name mongo-container \
  -e MONGO_INITDB_ROOT_USERNAME=admin \
  -e MONGO_INITDB_ROOT_PASSWORD=password \
  -p 27017:27017 \
  mongo
*/

func main() {
	app.Init()
}
