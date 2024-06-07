package main

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type element struct {
	Name            string `json:"name"`
	EstimatedNumber uint16 `json:"estimatedNumber"`
	Unit            string `json:"unit"`
}
type category struct {
	Name     string    `json:"name"`
	Elements []element `json:"elements"`
}
type response struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
	Data    []any  `json:"data"`
}
type land struct {
	Length uint16 `json:"length"`
	Width  uint16 `json:"width"`
}

func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/categories", func(c *gin.Context) { c.IndentedJSON(200, elementFactory()) })
	r.Use(LandValidationMiddleware()).POST("/calculate", func(c *gin.Context) { c.IndentedJSON(200, calculateElements(c)) }).Use(LandValidationMiddleware())
	r.Run() // listen and serve on 0.0.0.0:8080
}

func LandValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var land land
		err := c.ShouldBindJSON(&land)
		if err != nil || land.Width%96 != 0 || land.Length%33 != 0 {
			c.AbortWithStatus(400)
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func elementFactory() []category {
	var x = []category{
		{"لوله ها و المان ها", []element{
			{"ستون ها", 0, "عدد"},
			{"کانکتور مرکزی", 0, "عدد"},
			{" المان شبکه ای", 0, "عدد"},
			{" خورشیدی", 0, "عدد"},
			{" کانکتور مرکزی", 0, "عدد"},
			{" بادبند های X", 0, "عدد"},
			{" سفت کن قبل پنجره", 0, "عدد"},
			{" ستون های فرعی", 0, "عدد"},
			{"سر ستون کناری", 0, "عدد"},
			{"سرستون میانی", 0, "عدد"},
			{" پروفیل لاگینگ", 0, "متر"},
			{" تعداد پروفیل H", 0, "عدد"},
			{" فنر", 0, "متر"},
			{" ناودان کناری", 0, "متر"},
			{" ناودان میانی", 0, "متر"},
		}},
		{"هوک ها", []element{
			{"کمان دوم به ستون فرعی", 0, "عدد"},
			{"هوک کمان مورب", 0, "عدد"},
			{"هوک کمان اول به دوم", 0, "عدد"},
		}},
		{"بست ها", []element{
			{"بست گاتیک", 0, "عدد"},
			{"بست 80x 80 یک طرفه", 0, "عدد"},
			{"بست 80x 80 دوطرفه", 0, "عدد"},
			{"بست پارویی", 0, "عدد"},
			{"بست لامپی 4", 0, "عدد"},
			{"بست لامپی 6", 0, "عدد"},
			{"بست LOF", 0, "عدد"},
			{"بست اکسل", 0, "عدد"},
			{"بست سفت کن", 0, "عدد"},
			{"بسط رابط H", 0, "عدد"},
		}},
		{"پنجره", []element{
			{"رک و پینیون", 0, "عدد"},
			{"سفت کن زیرپنجره", 0, "عدد"},
			{"لوله شفت", 0, "عدد"},
			{"رابط لوله شفت", 0, "عدد"},
			{"رابط درونی و بیرونی H", 0, "عدد"},
			{"رابط کاپیج", 0, "عدد"},
		}},
	}
	return x
}
func calculateElements(c *gin.Context) response {
	var response = response{http.StatusBadRequest, "ارسال اشتباه", []any{}}
	var land land
	err := c.BindJSON(&land)
	if err != nil {
		Println(err)

	}
	Println(land.Width % 96)
	Println(land.Length % 33)
	response.Data = []any{}
	response.Status = 200
	response.Message = "درخواست موفق بود"
	return response
}
