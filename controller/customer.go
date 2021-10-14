package controller

import (
	"customer-restapi/customer"
	"customer-restapi/mysqllib"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get customers
func GetCustomers(c *gin.Context) {
	var customers []customer.Customer

	query := `SELECT customer_id, firstname, lastname, mobile, email, address_id
	FROM customer;`

	err := mysqllib.DBClient.Select(&customers, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customers)
}

// get specific customer
func GetCustomer(c *gin.Context) {
	id := c.Param("id")

	var customer customer.Customer

	query := `SELECT customer_id, firstname, lastname, mobile, email, address_id
	FROM customer
    WHERE customer_id = ?;`

	err := mysqllib.DBClient.Get(&customer, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": sql.ErrNoRows.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// create customer

func Createcustomer(c *gin.Context) {
	var customer customer.Customer

	err := c.ShouldBindJSON(&customer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// inserting customer
	customer_query := `INSERT INTO customer (customer_id, firstname, lastname, mobile, email, address_id) VALUES (?, ?, ?, ?, ?, ?);`
	customer_res, err := mysqllib.DBClient.Exec(customer_query, customer.ID, customer.Firstname,
		customer.Lastname, customer.Mobile, customer.Email, customer.AddressID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//address_id, err := address_res.LastInsertId()
	customer_id, err := customer_res.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"customer_id": customer_id,
	})

}

func UpdateCustomerName(c *gin.Context) {
	type Name struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	var name Name

	id := c.Param("id")

	err := c.ShouldBindJSON(&name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `UPDATE customer SET firstname = ?, lastname = ? WHERE customer_id = ?;`

	res, err := mysqllib.DBClient.Exec(query, name.Firstname, name.Lastname, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	no, err := res.RowsAffected()

	c.JSON(http.StatusOK, gin.H{
		"rows_affected": no,
	})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM customer WHERE customer_id = ?`

	res, err := mysqllib.DBClient.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	no, err := res.RowsAffected()

	c.JSON(http.StatusOK, gin.H{
		"rows_affected": no,
	})
}
