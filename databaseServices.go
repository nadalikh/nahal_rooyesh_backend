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
	db, err = sql.Open("mysql", "root:123456_nahal@tcp(127.0.0.1:3306)/green_house")
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
func getKhorishidiProperties() interface{} {
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
func getKhorshidiFabric(c *gin.Context) interface{} {
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
func removeWarmPrice(c *gin.Context) {
	id := c.Param("id")

	_, err = db.Query("delete from warm where id = ?  ", id)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(200, Response[string]{Message: "قیمت با موفقیت حذف شد ", Data: []string{""}})
}
func getFabricPrice(cnf map[string]interface{}, slug string) float32 {
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
	if err != nil {
		panic(err)
	}
	switch slug {
	case "khorshidi":
		return price.price * KHORSHIDI_LENGTH * multipled * 3.14
	default:
		return 0
	}
}
func getKhorshidiWarmPrice(cnf map[string]interface{}, slug string) float32 {
	warmConfig := cnf["props"].(map[string]interface{})
	result, err := db.Query("select price, slug, value from ( select * from warm  where element_slug=?)kf inner join  properties kp on kf.diagonal_id = kp.id or kf.thickness_id = kp.id where thickness_id =?  and diagonal_id=?", slug, warmConfig["thickness_id"], warmConfig["diagonal_id"])
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
	if err != nil {
		panic(err)
	}

	switch slug {
	case "khorshidi":
		return price.price * KHORSHIDI_LENGTH * multipled * 3.14
	default:
		return 0
	}
	//result, err := db.Query("select price from khorshidi_fabric where digonal_id = ? and thickness_id = ?", DTOConfig.fabric.)
}
