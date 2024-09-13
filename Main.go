package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type genericArray[T any] []T

type Element struct {
	Name            string  `json:"name"`
	EstimatedNumber float32 `json:"estimatedNumber"`
	Unit            string  `json:"unit"`
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
	Length float32 `json:"length"`
	Width  float32 `json:"width"`
}

func main() {
	defer db.Close()

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Resource not found",
		})
	})
	r.Use(CORSMiddleware())
	r.GET("/categories", func(c *gin.Context) { c.IndentedJSON(200, elementFactory(make(map[string]float32))) })
	r.POST("/calculate", LandValidationMiddleware(), func(c *gin.Context) { c.IndentedJSON(200, completeCalculate(c)) })
	r.GET("/khorshidi-properties", func(c *gin.Context) { c.IndentedJSON(200, getKhorishidiProperties()) })
	r.POST("/khorshidi-fabric", func(c *gin.Context) { c.IndentedJSON(200, addKhorshidiFabricPrice(c)) })

	r.Run() // listen and serve on 0.0.0.0:8080
}
func landValidation(land Land) []string {
	var errors []string

	if int(land.Length) <= 0 || int(land.Width) <= 0 {
		errors = append(errors, "مقدار طول یا عرض باید بزرگتر از 0 باشند.")
	}
	if int(land.Length)%3 != 0 {
		errors = append(errors, "مقدار طول باید مضرب ۳ باشد")
	}
	if int(land.Width)%96 != 0 {
		errors = append(errors, "مقدار عرض باید مضرب 9.6 باشد")
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

		data := make(map[string]float32)
		land_ := land.(Land)
		//done
		data["shaft"] = calculateShafts(land_)
		sideShafts := ((land_.Length / 3) + 1) * 2
		bow, chord := calculateArcAndChord(land_)
		//
		data["bow"] = bow
		//
		data["chord"] = chord
		//
		data["267cmElements"] = calculate267cmElements(chord)
		data["176cmElements"] = calculate176cmElements(chord)
		data["150cmElements"] = calculate150cmElements(chord)
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
		data["centralGutter"] = calculateCentralGutter(land_)
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
		data["windowPicket"] = calculateWindowPicket(land_)
		//
		data["rack"] = calculateRack(bow)
		//
		data["lamp6"] = calculate6lamp(firstBowToSecond, secondToShaft, chord)
		//
		data["lamp4"] = calculate4lamp(data["khorshidi"], chord)
		//
		data["oneWay80X804cm"] = calculate80X80OneWay4cm(
			data["windBreaker"]*2,
			secondToShaft,
			sideShafts,
		)
		data["OneWay80X803cmBushan"] = calculateOneWay80X803cmBushan(land_,
			data["secondaryShaft"],
			sideShafts)
		//
		data["towWay80X80"] = calculate80X80TowWay(land_, data["shaft"], sideShafts)
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
		data["hardenerBushen"] = calculateHardenerBushen(land_, bow)
		//
		data["sideHardener"] = calculateSideHardener(land_, bow)
		//
		data["headAndTailHardener"] = calculateHeadAndTailHardener(land_, bow)
		//
		data["H_InOutConnector"] = calculate_H_InOutConnector(land_)
		//
		data["golpich"] = calculateGolpich(data["rowing"])
		//done
		data["locking"] = calculateLocking(land_, data["sideGutter"], data["centralGutter"])
		data["horseShoe"] = calculateHorseShoe(data["secondaryShaft"])
		data["gableScrew"] = calculateGableScrew(
			data["bindingGathic"]+
				data["oneWay80X804cm"]+
				data["OneWay80X803cmBushan"]+
				data["towWay80X80"]+
				data["lamp4"]+
				data["lamp6"]+
				data["hardenerBushen"]+
				data["sideHardener"]+
				data["headAndTailHardener"]+
				data["H_InOutConnector"]+
				data["LOF"]+
				data["excel"], data["locking"])
		//done
		data["spring"] = calculateSpring(land_, data["headOfWindowH"], data["locking"])
		data["bolt_2cm"],
			data["bolt_3cm"],
			data["bolt_4cm"],
			data["bolt_5cm"],
			data["bolt_6cm"],
			data["bolt_10cm"],
			data["bolt_halfThread"] = calculateBoltAndNut(
			data["bindingGathic"], data["excel"],
			data["rowing"], data["golpich"],
			data["H_InOutConnector"], data["lamp4"],
			data["lamp6"], data["horseShoe"],
			data["centralHeadShaft"], data["sideHeadShaft"],
			data["oneWay80X804cm"], data["OneWay80X803cmBushan"],
			data["towWay80X80"], data["windowPicket"],
		)

		result := elementFactory(data)
		var response = Response[Category]{"ارسال موفق", result}

		return response
	}
}
