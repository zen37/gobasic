package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type payment struct {
	ID       string  `json:"id"`
	Invoice  string  `json:"invoice"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

var payments = []payment{}

/* // albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
*/
func main() {

	router := gin.Default()

	router.GET("/payments", getPayments)
	router.POST("/payments", postPayment)
	/* associate GET and /albums/:id path with a handler.
	in Gin, the colon preceding an item in the path signifies that the item is a path parameter. */
	router.GET("/payments/:id", getPaymentByID)

	router.Run("localhost:8080") // attaches the router to an http.Server and start server

}

/*
gin.Context is the most important part of Gin; it carries request details, validates and serializes JSON, and more.
Despite the similar name, this is different from Go’s built-in context package.
*/

func getPayments(c *gin.Context) {
	//Context.IndentedJSON can be replaced with a call to Context.JSON to send more compact JSON.
	//In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
	c.IndentedJSON(http.StatusOK, payments)
	//c.JSON(http.StatusOK, albums)
}

//postAlbum adds an album from JSON received in request body
func postPayment(c *gin.Context) {

	var newPayment payment

	//binds the request body to newAlbum
	if err := c.BindJSON(&newPayment); err != nil {
		return
	}

	payments = append(payments, newPayment)
	c.IndentedJSON(http.StatusCreated, newPayment)

}

func getPaymentByID(c *gin.Context) {

	id := c.Param("id")

	for _, a := range payments {
		if a.ID == id {
			c.IndentedJSON(http.StatusFound, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "payment not found"})
}
