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
type Response[T interface{}] struct {
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
	r.GET("/categories", func(c *gin.Context) { c.IndentedJSON(200, elementFactory(make(map[string]uint16))) })
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
func completeCalculate(c *gin.Context) interface{} {
	if land, existed := c.Get("land"); !existed || land == nil {
		var response = Response[any]{"ارسال اشتباه", []any{}}
		response.Message = "خطای سیستم"
		return response
	} else {

		data := make(map[string]uint16)
		land_ := land.(Land)
		//done
		data["shaft"] = calculateShafts(land_)
		bow, chord := calculateArcAndChord(land_)
		//
		data["bow"] = bow
		//
		data["chord"] = chord
		//
		data["elementsOfchord"] = calculateElementsOfChord(chord)
		//done

		data["khorshidi"] = calculateKhorshidi(chord)
		//done
		data["centralConnector"] = calculateCentralConnector(land_)
		//done
		data["hardenerBeforeWindow"] = calculateHardenerBeforeWindow(land_)
		//done
		data["windBreaker"] = calculateWindBreaker(land_)
		//done
		data["secondaryShaft"] = calculateSecondaryShaft(land_)
		//
		data["sideGutter"] = calculateTheSideGutter(land_)
		//
		data["centralGutter"] = calculateCentralGutter(land_, data["sideGutter"])
		//done
		data["sideHeadShaft"] = calculateTheSideHeadShaft(land_)
		//done
		data["centralHeadShaft"] = calculateTheCentralHeadShaft(land_, data["sideHeadShaft"])
		//
		data["bindingGathic"] = calculateBindingGothic(bow)
		//
		firstBowToSecond, diagonal, secondToShaft := calculateHooks(land_)
		//
		data["firstBowToSecond"] = firstBowToSecond
		//
		data["diagonal"] = diagonal
		//
		data["secondToShaft"] = secondToShaft
		//
		data["windowPicket"] = calculateWindowPicket(bow)
		//
		data["rack"] = calculateRack(bow)
		//
		data["lamp6"] = calculate6lamp(firstBowToSecond, secondToShaft, chord)
		//
		data["lamp4"] = calculate4lamp(data["khorshidi"], chord)
		//
		data["oneWay80X80"] = calculate80X80OneWay(land_,
			data["windBreaker"]*2,
			secondToShaft,
		)
		//
		data["towWay80X80"] = calculate80X80TowWay(land_, data["shaft"])
		//
		data["rowing"] = calculateRowing(data["windowPicket"])
		//
		data["LOF"] = calculateLOF(data["windowPicket"])
		//
		data["excel"] = calculateExcel(data["windowPicket"])
		//
		data["pinion"] = calculatePinion(data["windowPicket"])
		//
		data["headOfWindowH"] = calculateHeadOfWindowH(land_)
		//
		data["hardenerUnderTheWindow"] = calculateHardenerUnderTheWindow(land_)
		//
		data["shaftPipe"] = calculateShaftPipe(land_)
		//
		data["shaftPipeConnector"] = calculateShaftPipeConnector(land_)
		//
		data["hardenerBushen"] = calculateHardenerBushen(data["windowPicket"])
		//
		data["H_InOutConnector"] = calculate_H_InOutConnector(data["headOfWindowH"])
		//
		data["kapage"] = calculateKapage(data["rowing"])
		//done
		data["locking"] = calculateLocking(land_, data["sideGutter"], data["centralGutter"])
		//done
		data["spring"] = calculateSpring(land_, data["headOfWindowH"], data["locking"])
		result := elementFactory(data)
		var response = Response[Category]{"ارسال موفق", result}

		return response
	}
}
