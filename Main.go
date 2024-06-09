package main

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type genericArray[T any] []T

type Element struct {
	Name            string `json:"name"`
	EstimatedNumber uint16 `json:"estimatedNumber"`
	Unit            string `json:"unit"`
}
type Category struct {
	Name     string    `json:"name"`
	Elements []Element `json:"elements"`
}
type Response[T any] struct {
	Message string          `json:"message"`
	Data    genericArray[T] `json:"data"`
}
type Land struct {
	Length uint16 `json:"length"`
	Width  uint16 `json:"width"`
}

func main() {

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		Println("here")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Resource not found",
		})
	})
	r.Use(CORSMiddleware())
	r.GET("/categories", func(c *gin.Context) { c.IndentedJSON(200, elementFactory()) })
	r.POST("/calculate", LandValidationMiddleware(), func(c *gin.Context) { c.IndentedJSON(200, completeCalculate(c)) })

	r.Run() // listen and serve on 0.0.0.0:8080
}
func landValidation(land Land) []string {
	var errors []string

	if land.Length%3 != 0 {
		errors = append(errors, "مقدار طول باید ضریب ۳ باشد")
	}
	if land.Width%96 != 0 {
		errors = append(errors, "مقدار عرض باید ضریب 9.6 باشد")
	}
	return errors
}

func LandValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var land Land
		err := c.ShouldBindJSON(&land)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}
		validationErrors := landValidation(land)
		if len(validationErrors) > 0 {
			var errors []string
			for _, value := range validationErrors {
				errors = append(errors, value)
			}
			c.IndentedJSON(400, Response[string]{"خطای ولیدیشن", errors})
			c.Abort()
			return
		}
		c.Set("land", land)
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
func elementFactory() []Category {
	var x = []Category{
		{"لوله ها و المان ها", []Element{
			{"ستون ها", 0, "عدد"},
			{"کانکتور مرکزی", 0, "عدد"},
			{" المان شبکه ای", 0, "عدد"},
			{" خورشیدی", 0, "عدد"},
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
		{"هوک ها", []Element{
			{"کمان دوم به ستون فرعی", 0, "عدد"},
			{"هوک کمان مورب", 0, "عدد"},
			{"هوک کمان اول به دوم", 0, "عدد"},
		}},
		{"بست ها", []Element{
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
		{"پنجره", []Element{
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
func completeCalculate(c *gin.Context) Response[any] {
	var response = Response[any]{"ارسال اشتباه", []any{}}
	if land, existed := c.Get("land"); !existed || land == nil {
		response.Message = "خطای سیستم"
		return response
	} else {
		calculateElements(land.(Land))
		response.Data = []any{}
		response.Message = "درخواست موفق بود"
		return response
	}
}

func calculateElements(land Land) {
	shaft := ((land.Length / 3) + 1) * ((land.Width * 10 / 96) + 1)
	Println(shaft)
}
