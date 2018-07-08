package controllers

import "github.com/gin-gonic/gin"
import (
	"gogo/service"
	"net/http"
	"strconv"
	"gogo/models"
	"time"
	"github.com/gin-gonic/gin/binding"
	"log"
	"strings"
)

func Index(c *gin.Context)  {
	//where := map[string]string{"name":"zhanglong"}
	a := service.GetBook()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    0,
		"msg":     a})
}

func Update(c *gin.Context)  {
	rid := c.DefaultQuery("id", "1")
	id, _ := strconv.Atoi(rid)
	bookModel := models.Book{
		Id:          0,
		Name:        c.DefaultQuery("name", "name"),
		Author:      "",
		Price:       c.DefaultQuery("price", "3.0"),
		Create_time: time.Time{},
		Update_time: time.Time{},
	}
	ret := service.UpdateBookById(id, bookModel)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    0,
		"msg": ret,
	})
}

func Delete(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	book_id, _:= strconv.Atoi(id)
	ret, _ := service.DeleteBook(book_id)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    0,
		"msg": ret,
	})
}

func Insert(c *gin.Context) {
	form := &models.BookAdd{}
	if validate := c.ShouldBindWith(form, binding.Form); validate != nil {
		log.Fatalln(c, -1, validate.Error())
		return
	}

	id, err := service.InsertBook("books", form)
	if err != nil {
		log.Fatalln(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": strings.Trim(strconv.FormatInt(id, 10) + " ", " " ),
		"code": 0,
		"msg": form,
	})
}