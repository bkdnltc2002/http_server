package handler

import (
	"github.com/gin-gonic/gin"
	"app/database"
	"app/model"
	"log"
	"net/http"
	"fmt"
)

func GetUserHandler(c *gin.Context){
	db := database.GetDatabase()
	var students []model.Student
	err := db.Preload("ClassInfo").Find(&students).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(students)

	c.JSON(http.StatusOK, students)
}

func SearchHandler(c *gin.Context) {
	db := database.GetDatabase()
	searchQuery := c.Query("name")

	var students []model.Student
	err :=db.Where("student_name LIKE ?", "%"+searchQuery+"%").Find(&students).Error
	if err != nil {
		c.String(http.StatusNotFound, "The name does not exist")
		return
	}
	c.JSON(http.StatusOK, students)
}