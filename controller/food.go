package controller

import (
	"customer-restapi/food"
	"customer-restapi/mysqllib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFoods(c *gin.Context) {
	var foods []food.Food

	query := `SELECT food_id, food_name, short_desc, long_desc, price, is_availability FROM food`

	err := mysqllib.DBClient.Select(&foods, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, foods)
}

func GetFood(c *gin.Context) {
	id := c.Param("id")

	query := `SELECT food_id, food_name, short_desc, long_desc, price, is_availability FROM food WHERE food_id = ?`

	var food food.Food

	err := mysqllib.DBClient.Get(&food, query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, food)
}

func CreateFood(c *gin.Context) {
	var food food.Food

	err := c.ShouldBindJSON(&food)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `INSERT INTO food (food_id, food_name, short_desc, long_desc, price, is_availability) VALUES (?, ?, ?, ?, ?, ?);`

	res, err := mysqllib.DBClient.Exec(query, food.ID, food.FoodName, food.ShortDesc, food.LongDesc,
		food.Price, food.IsAvailable)

	id, err := res.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"inserted_id": id,
	})
}

func UpdateFoodPrice(c *gin.Context) {
	type ProductPrice struct {
		Price int `json:"price"`
	}

	var price ProductPrice

	err := c.ShouldBindJSON(&price)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")

	query := `UPDATE food SET price = ? WHERE food_id = ?`

	res, err := mysqllib.DBClient.Exec(query, price.Price, id)

	affected, err := res.RowsAffected()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"affected": affected,
	})
}

func DeleteFood(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM food WHERE food_id = ?`

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
