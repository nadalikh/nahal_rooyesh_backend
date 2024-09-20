package main

import (
	"database/sql"
	"fmt"
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

type warm struct {
	ID    int `json:"id"`
	Price int `json:"price"`
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
func addKhorshidiFabricPrice(c *gin.Context) {
	var priceDTO struct {
		Diagonal  int `json:"diagonal"`
		Thickness int `json:"thickness""`
		Price     int `json:"price"`
	}
	type khorshidiFabricPricDTO struct {
		ID          int `json:"id"`
		DiagonalId  int `json:"diagonal_id"`
		ThicknessId int `json:"thickness_id"`
		Price       int `json:"price"`
	}
	result, err := db.Query("select * from khorshidi_fabric")
	if err != nil {
		panic(err)
	}

	err = c.BindJSON(&priceDTO)
	//check for repetitive prices
	var data khorshidiFabricPricDTO
	for result.Next() {
		err = result.Scan(&data.ID, &data.DiagonalId, &data.ThicknessId, &data.Price)
		if data.DiagonalId == priceDTO.Diagonal && data.ThicknessId == priceDTO.Thickness {

			c.IndentedJSON(400, Response[string]{"خطای ولیدیشن", []string{"اطلاعات تکراری است"}})
			return
		}
	}
	if err != nil {
		panic(err)
	}
	_, err = db.Query("insert into khorshidi_fabric (diagonal_id, thickness_id, price) values (?, ?, ?);", priceDTO.Diagonal, priceDTO.Thickness, priceDTO.Price)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(200, Response[string]{Message: "قیمت با موفقیت اضافه شد", Data: []string{""}})
}
func addWarm(c *gin.Context) interface{} {
	var warmPriceDTO struct {
		ElementSlug string `json:"element_slug"`
		Price       int    `json:"price"`
	}

	err := c.BindJSON(&warmPriceDTO)
	if err != nil {
		panic(err)
	}
	fmt.Println(warmPriceDTO.ElementSlug)
	result, err := db.Query("select * from warm where element_slug = ?", warmPriceDTO.ElementSlug)
	if err != nil {
		panic(err)
	}
	found := false
	for result.Next() {
		found = true
		break
	}
	if found {
		_, err = db.Query("update warm set price = ? where element_slug = ? ;", warmPriceDTO.Price, warmPriceDTO.ElementSlug)

	} else {
		_, err = db.Query("insert into warm (element_slug, price) values (?, ?);", warmPriceDTO.ElementSlug, warmPriceDTO.Price)

	}
	if err != nil {
		panic(err)
	}
	if found {
		return Response[string]{Message: "قیمت با موفقیت تغییر کرد", Data: []string{""}}
	} else {
		return Response[string]{Message: "قیمت با موفقیت اضافه شد", Data: []string{""}}
	}
}
func getWarm(c *gin.Context) interface{} {
	params := c.Request.URL.Query()
	elementSlug := params["element_slug"][0]
	var warm warm
	result, err := db.Query("select id, price from warm where element_slug = ?", elementSlug)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&warm.ID, &warm.Price)
		break
	}
	if err != nil {
		panic(err)
	}
	return warm
}
func getKhorshidiFabric(c *gin.Context) interface{} {
	type khorshidiFabricPricDTO struct {
		ID          int `json:"id"`
		DiagonalId  int `json:"diagonal_id"`
		ThicknessId int `json:"thickness_id"`
		Price       int `json:"price"`
	}
	result, err := db.Query("select * from khorshidi_fabric")
	if err != nil {
		panic(err)
	}
	list := make([]khorshidiFabricPricDTO, 0)
	var data khorshidiFabricPricDTO
	for result.Next() {
		err = result.Scan(&data.ID, &data.DiagonalId, &data.ThicknessId, &data.Price)
		list = append(list, data)
	}
	return list
}
