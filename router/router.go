package router

import (
	"customer-restapi/controller"

	"github.com/gin-gonic/gin"
)

func RouteInitializer() {
	r := gin.Default()

	// customers

	r.GET("/api/customers", controller.GetCustomers)
	r.GET("/api/customers/:id", controller.GetCustomer)
	r.POST("/api/customers", controller.Createcustomer)
	r.PUT("/api/customers/:id", controller.UpdateCustomerName)
	r.DELETE("/api/customers/:id", controller.DeleteCustomer)

	// foods
	r.GET("/api/foods", controller.GetFoods)
	r.GET("/api/foods/:id", controller.GetFood)
	r.POST("/api/foods", controller.CreateFood)
	r.PUT("/api/foods/:id", controller.UpdateFoodPrice)
	r.DELETE("/api/foods/:id", controller.DeleteFood)

	r.Run(":3000")
}
