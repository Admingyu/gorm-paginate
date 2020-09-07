package main

import (
	"log"
	"strconv"

	"github.com/Admingyu/gorm-paginate/paginate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"not null;size:100"`
	Email string `json:"email" gorm:"size:200"`
	Phone string `json:"phone" gorm:"size:11"`
}

func DefaultRoute(c *gin.Context) {
	db, err := gorm.Open("sqlite3", "example.db")
	if err != nil {
		log.Println(err)
	}

	pageIndex, _ := strconv.Atoi(c.DefaultQuery("page_index", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "2"))

	dbQuery := db.Model(&Users{}).Select("id, name, email, phone")

	params := paginate.PageParams{
		Serializer: Users{},
		Order:      []string{"id desc", "name desc"},
		PageIndex:  pageIndex,
		PageSize:   pageSize,
	}
	data := paginate.Pagenate(dbQuery, params)
	c.JSON(200, data)
}

func main() {

	router := gin.Default()
	router.GET("/", DefaultRoute)
	err := router.Run("0.0.0.0:8989")
	print(err)

}
