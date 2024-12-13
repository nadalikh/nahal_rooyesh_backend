package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"math"
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
	ID          int `json:"id"`
	Price       int `json:"price"`
	ThicknessId int `json:"thickness_id"`
	DiagonalId  int `json:"diagonal_id"`
}
type Config struct {
	ID   int    `json:"id"`
	Slug string `json:"element_slug"`
	Cnf  string `json:"config_json"`
}

func init() {
	db, err = sql.Open("mysql", "root:expecto-patronum1379@tcp(127.0.0.1:3306)/green_house")
	//db, err = sql.Open("mysql", "root:brauvZtcAqc6UJf@tcp(127.0.0.1:3306)/green_house")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
}
func loadAllConfigs() *map[string]string {
	result, err := db.Query("select * from config")
	var cnf Config
	var configs map[string]string
	configs = make(map[string]string)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&cnf.ID, &cnf.Slug, &cnf.Cnf)
		if err != nil {
			panic(err)
		}
		configs[cnf.Slug] = cnf.Cnf
	}
	return &configs

}
func getPipeProperties() interface{} {
	result, err := db.Query("SELECT * FROM properties")
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
func updateIronProperties(c *gin.Context) {
	var ironProps struct {
		ID    int `json:"id"`
		Value int `json:"value"`
	}
	err := c.BindJSON(&ironProps)
	if err != nil {
		panic(err)
	}
	_, err = db.Query("UPDATE iron_properties SET value = ? WHERE id = ?", ironProps.Value, ironProps.ID)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(200, Response[string]{Message: "مشخصه آهن با موفقیت تغییر کرد", Data: []string{""}})

}
func getIronProperties(c *gin.Context) {
	params := c.Request.URL.Query()
	slug := params["slug"][0]
	fmt.Println(slug)
	var ironProp struct {
		Slug  string `json:"slug"`
		Value string `json:"value"`
	}

	if err != nil {
		panic(err)
	}
	result, err := db.Query("select slug, value from iron_properties where slug=?", slug)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		err := result.Scan(&ironProp.Slug, &ironProp.Value)
		if err != nil {
			panic(err)
		}
	}
	c.IndentedJSON(200, ironProp)

}
func addFabric(c *gin.Context) {
	var priceDTO struct {
		Diagonal    int    `json:"diagonal"`
		Thickness   int    `json:"thickness""`
		Price       int    `json:"price"`
		ElementSlug string `json:"element_slug"`
	}
	type khorshidiFabricPricDTO struct {
		ID          int    `json:"id"`
		DiagonalId  int    `json:"diagonal_id"`
		ThicknessId int    `json:"thickness_id"`
		Price       int    `json:"price"`
		ElementSlug string `json:"element_slug"`
	}
	result, err := db.Query("select * from fabric")
	if err != nil {
		panic(err)
	}

	err = c.BindJSON(&priceDTO)
	//check for repetitive prices
	var data khorshidiFabricPricDTO
	for result.Next() {
		err = result.Scan(&data.ID, &data.ElementSlug, &data.Price, &data.ThicknessId, &data.DiagonalId)
		if data.DiagonalId == priceDTO.Diagonal && data.ThicknessId == priceDTO.Thickness && data.ElementSlug == priceDTO.ElementSlug {

			c.IndentedJSON(400, Response[string]{"خطای ولیدیشن", []string{"اطلاعات تکراری است"}})
			return
		}
	}
	if err != nil {
		panic(err)
	}
	_, err = db.Query("insert into fabric (diagonal_id, thickness_id, price, element_slug) values (?, ?, ?, ?);", priceDTO.Diagonal, priceDTO.Thickness, priceDTO.Price, priceDTO.ElementSlug)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(200, Response[string]{Message: "قیمت با موفقیت اضافه شد", Data: []string{""}})
}
func addWarm(c *gin.Context) {
	var priceDTO struct {
		Diagonal    int    `json:"diagonal"`
		Thickness   int    `json:"thickness""`
		Price       int    `json:"price"`
		ElementSlug string `json:"element_slug"`
	}
	type warmPricDTO struct {
		ID          int    `json:"id"`
		DiagonalId  int    `json:"diagonal_id"`
		ThicknessId int    `json:"thickness_id"`
		Price       int    `json:"price"`
		ElementSlug string `json:"element_slug"`
	}
	result, err := db.Query("select * from warm")
	if err != nil {
		panic(err)
	}

	err = c.BindJSON(&priceDTO)
	//check for repetitive prices
	var data warmPricDTO
	for result.Next() {
		err = result.Scan(&data.ID, &data.DiagonalId, &data.ThicknessId, &data.Price, &data.ElementSlug)
		if data.DiagonalId == priceDTO.Diagonal && data.ThicknessId == priceDTO.Thickness && data.ElementSlug == priceDTO.ElementSlug {

			c.IndentedJSON(400, Response[string]{"خطای ولیدیشن", []string{"اطلاعات تکراری است"}})
			return
		}
	}
	if err != nil {
		panic(err)
	}
	_, err = db.Query("insert into warm (diagonal_id, thickness_id, price, element_slug) values (?, ?, ?, ?);", priceDTO.Diagonal, priceDTO.Thickness, priceDTO.Price, priceDTO.ElementSlug)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(200, Response[string]{Message: "قیمت با موفقیت اضافه شد", Data: []string{""}})
}

func getWarm(c *gin.Context) interface{} {
	params := c.Request.URL.Query()
	elementSlug := params["element_slug"][0]

	result, err := db.Query("select id, diagonal_id, thickness_id, price from warm where element_slug=?", elementSlug)
	if err != nil {
		panic(err)
	}
	list := make([]warm, 0)
	var data warm
	for result.Next() {
		err = result.Scan(&data.ID, &data.DiagonalId, &data.ThicknessId, &data.Price)
		list = append(list, data)
	}
	return list
}
func getFabric(c *gin.Context) interface{} {
	type khorshidiFabricPricDTO struct {
		ID          int `json:"id"`
		DiagonalId  int `json:"diagonal_id"`
		ThicknessId int `json:"thickness_id"`
		Price       int `json:"price"`
	}
	params := c.Request.URL.Query()
	elementSlug := params["element_slug"][0]
	result, err := db.Query("select id, price ,thickness_id, diagonal_id from fabric where element_slug = ?", elementSlug)
	if err != nil {
		panic(err)
	}
	list := make([]khorshidiFabricPricDTO, 0)
	var data khorshidiFabricPricDTO
	for result.Next() {
		err = result.Scan(&data.ID, &data.Price, &data.ThicknessId, &data.DiagonalId)
		list = append(list, data)
	}
	return list
}
func removeFabricPrice(c *gin.Context) {
	id := c.Param("id")

	_, err = db.Query("delete from fabric where id = ?  ", id)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(200, Response[string]{Message: "قیمت با موفقیت حذف شد ", Data: []string{""}})
}

func getFabricPipePrice(cnf map[string]interface{}, slug string) float32 {
	warmConfig := cnf["props"].(map[string]interface{})
	result, err := db.Query("select price, slug, value from ( select * from fabric  where element_slug=?)kf inner join  properties kp on kf.diagonal_id = kp.id or kf.thickness_id = kp.id where thickness_id =?  and diagonal_id=?", slug, warmConfig["thickness_id"], warmConfig["diagonal_id"])
	var price struct {
		price float32
		slug  string
		value float32
	}
	var multipled float32
	multipled = 1
	for result.Next() {
		err := result.Scan(&price.price, &price.slug, &price.value)
		if err != nil {
			panic(err)
		}
		multipled *= price.value
	}
	multipled = 1
	if err != nil {
		panic(err)
	}
	switch slug {
	case "pipe":
		//return price.price * KHORSHIDI_LENGTH
		return price.price * float32(warmConfig["length"].(float64))

	default:
		return 0
	}
}
func getPipeWarmPrice(cnf map[string]interface{}, slug string) float32 {
	warmConfig := cnf["props"].(map[string]interface{})
	//result, err := db.Query("select price, slug, value from ( select * from warm  where element_slug=?)kf inner join  properties kp on kf.diagonal_id = kp.id or kf.thickness_id = kp.id where thickness_id =?  and diagonal_id=?", slug, warmConfig["thickness_id"], warmConfig["diagonal_id"])
	resultIronPrice, err := db.Query("select value from iron_properties  where id =1")
	var ironPrice struct {
		Value float32 `json:"value"`
	}
	for resultIronPrice.Next() {
		err := resultIronPrice.Scan(&ironPrice.Value)
		if err != nil {
			panic(err)
		}
	}
	err = resultIronPrice.Close()
	if err != nil {
		panic(err)
	}
	result, err := db.Query("select value, slug from properties  where id =?  or id=?", warmConfig["thickness_id"], warmConfig["diagonal_id"])

	var price struct {
		//price float32
		slug  string
		value float64
	}
	var space float64
	space = 1

	D := float64(1.0)
	T := float64(1.0)
	for result.Next() {
		//err := result.Scan(&price.price, &price.slug, &price.value)
		err := result.Scan(&price.value, &price.slug)
		if err != nil {
			panic(err)
		}
		//multipled *= price.value / 10 //convert to cm
		if price.slug == "diagonal" {
			D = price.value / 10
		} else if price.slug == "thickness" {
			T = price.value / 10
		}
	}
	space = (math.Pi * math.Pow(D, 2) / 4) - (math.Pi * math.Pow(D-(2*T), 2) / 4)
	fmt.Println(space)
	if err != nil {
		panic(err)
	}

	switch slug {
	case "pipe":
		//return price.price * KHORSHIDI_LENGTH * multipled * 3.14 * IRON_DENSITY
		//			Gram            CM				  CM				CM^3
		return float32(ironPrice.Value) * float32(warmConfig["length"].(float64)) * float32(space) * IRON_DENSITY
	default:
		return 0
	}
	//result, err := db.Query("select price from khorshidi_fabric where digonal_id = ? and thickness_id = ?", DTOConfig.fabric.)
}
