package main

import (
	"customer-restapi/mysqllib"
	"customer-restapi/router"
)

func main() {
	mysqllib.DBInitializer()

	router.RouteInitializer()
}
