package controllers

import (
	"strconv"
	database "webapi/databases"
	"webapi/models"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {

	// c.JSON(200, gin.H{
	// 	"value" : "OK",
	// })

	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Could not find Book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)

}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Could not bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Could not create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)

}

func ShowBooks(c *gin.Context) {

	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Could not list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)

}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Could not bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Could not update book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {

	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.Delete(&models.Book{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Could not delete book: " + err.Error(),
		})

		return
	}

	c.JSON(204, book)
}
