package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

type properties struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
	Slug  string `json:"slug"`
}

func init() {
	db, err = sql.Open("mysql", "root:expecto-patronum1379@tcp(127.0.0.1:3306)/green_house")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
}
func getKhorishidiProperties() interface{} {
	result, err := db.Query("SELECT * FROM khorshidi_properties")
	if err != nil {
		panic(err)
	}
	pr := make([]properties, 0)
	var prop properties
	for result.Next() {
		result.Scan(&prop.ID, &prop.Slug, &prop.Value)
		pr = append(pr, prop)
	}
	return pr
}
func addKhorshidiFabricPrice(c *gin.Context) interface{} {
	var priceDTO struct {
		Diagonal  int `json:"diagonal"`
		Thickness int `json:"thickness""`
		Price     int `json:"price"`
	}
	err := c.BindJSON(&priceDTO)
	if err != nil {
		panic(err)
	}
	return Response[string]{Message: "قیمت با موفقیت اضافه شد", Data: []string{""}}
}
