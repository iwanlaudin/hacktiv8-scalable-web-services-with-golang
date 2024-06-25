package main

import "rest-api-assignment/routes"

func main() {
	PORT := ":4000"

	routes.StartServer().Run(PORT)
}
