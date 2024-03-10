package main

import (
	"go_restapi_assignment2/routers"
)

func main() {
	const PORT = ":8090"

	routers.StartServer().Run(PORT)
}
