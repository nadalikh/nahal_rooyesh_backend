package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func calculateShafts(land Land) float32 {
	shaft := ((land.Length / 3) + 1) * ((land.Width * 10 / 96) + 1)
	return shaft
}
func calculateArcAndChord(land Land) (float32, float32) {
	bow := ((land.Length / 3) + 1) * (land.Width * 10 / 96)
	chord := bow - ((land.Width * 10 / 96) * 2)
	return bow, chord
}
func calculate267cmElements(chord float32) float32 {
	const elementCategoryPerChord = 2
	return chord * elementCategoryPerChord
}
func calculate176cmElements(chord float32) float32 {
	const elementCategoryPerChord = 2
	return chord * elementCategoryPerChord
}
func calculate150cmElements(chord float32) float32 {
	const elementCategoryPerChord = 2
	return chord * elementCategoryPerChord
}
func calculateKhorshidi(chord float32) float32 {
	const numberOfKhorshidiPerChord = 2
	return chord * numberOfKhorshidiPerChord
}
func calculateCentralConnector(land Land) float32 {
	x := (land.Length / 6) * (land.Width * 10 / 96)
	return x
}
func calculateHardenerBeforeWindow(land Land) float32 {
	return (land.Length / 6) * (land.Width * 10 / 96)
}
func calculateWindBreaker(land Land) float32 {
	return ((land.Width * 10 / 96) + 1) * 2
}
func calculateSecondaryShaft(land Land) float32 {
	const secondaryShaftPerAperture = 4
	return (land.Width * 10 / 96) * secondaryShaftPerAperture
}
func calculateTheSideGutter(land Land) float32 {
	return land.Length / 6 * 2
}
func calculateCentralGutter(land Land) float32 {
	return land.Length / 6 * 9
}
func calculateTheSideHeadShaft(land Land) float32 {
	//return 2 * land.Length / 3
	return (land.Length/3 + 1) * 2
}
func calculateTheCentralHeadShaft(land Land, sideHeadShaft float32) float32 {
	return ((land.Length / 3) * ((land.Width * 10 / 96) + 1)) - sideHeadShaft
}
func calculateBindingGothic(bow float32) float32 {
	return bow
}
func calculateHooks(land Land) (float32, float32, float32) {
	firstBowToSecond := (land.Width * 10 / 96) * 4
	diagonal := (land.Width * 10 / 96) * 4
	secondToShaft := (land.Width * 10 / 96) * 4
	return firstBowToSecond, diagonal, secondToShaft
}
func calculateWindowPicket(land Land) float32 {
	return (land.Length/3 + 1) * (land.Width / 96 * 10)
}
func calculateRack(bow float32) float32 {
	return bow
}
func calculate6lamp(firstBowToSecond float32, secondToShaft float32, chord float32) float32 {
	return (firstBowToSecond * 2) + secondToShaft + (chord * 4)
}
func calculate4lamp(khorshidi float32, chord float32) float32 {
	return (chord * 2) + chord + khorshidi
}
func calculate80X80OneWay4cm(windbreakerShafts float32, secondToShaft float32, sideShafts float32) float32 {
	//sideShaft := 2 * ((land.Length / 3) + 1)
	const sideRows = 2
	return 1 + (windbreakerShafts * 2) +
		secondToShaft +
		(sideShafts * 2) -
		(sideRows * 2)
}
func calculateOneWay80X803cmBushan(land Land, secondaryShafts float32, sideShafts float32) float32 {
	return (sideShafts * 2) + ((land.Width / 96 * 10) + 1) + secondaryShafts*2
}
func calculate80X80TowWay(land Land, shaft float32, sideShafts float32) float32 {

	centralRows := (land.Width / 96 * 10) - 1 // +1 -2
	return (shaft - sideShafts) - (centralRows * 2)
}
func calculateRowing(windowPicket float32) float32 {
	return windowPicket
}
func calculateLOF(windowPicket float32) float32 {
	return windowPicket
}
func calculateExcel(windowPicket float32) float32 {
	return windowPicket
}
func calculatePinion(windowPicket float32) float32 {
	return windowPicket
}
func calculateHeadOfWindowH(land Land) float32 {
	return land.Length / 6 * land.Width / 96 * 10
}
func calculateHardenerUnderTheWindow(land Land) float32 {
	return land.Length / 6 * land.Width / 96 * 10
}
func calculateShaftPipe(land Land) float32 {
	return land.Length / 6 * land.Width / 96 * 10
}
func calculateShaftPipeConnector(land Land) float32 {
	return ((land.Length / 6) - 1) * land.Width / 96 * 10
}
func calculateHardenerBushen(land Land, bows float32) float32 {
	return bows - (land.Width / 96 * 10 * 2)
}
func calculateSideHardener(land Land, bows float32) float32 {
	return land.Length / 6 * 2
}
func calculateHeadAndTailHardener(land Land, bows float32) float32 {
	return land.Length / 6 * 4
}
func calculate_H_InOutConnector(land Land) float32 {
	return ((land.Length / 6) - 1) * land.Width / 96 * 10
}
func calculateGolpich(rowing float32) float32 {
	return rowing * 2
}
func calculateLocking(land Land, sideGutter float32, centralGutter float32) float32 {
	return (6 * 2 * (sideGutter + centralGutter)) +
		((land.Width * 10 / 96) * 11 * 2) + 20
}
func calculateSpring(land Land, headOfWindowH float32, locking float32) float32 {
	return ((land.Width * 10 / 96) * 2 * 11) +
		(headOfWindowH * 6) + locking
}
func calculateHorseShoe(secondaryShaft float32) float32 {
	return secondaryShaft
}
func calculateGableScrew(sumOfAllFastening float32, locking float32) float32 {
	return sumOfAllFastening + (locking / 0.3)
}
func calculateBoltAndNut(bindingGathic float32,
	excel float32, rowing float32,
	golpich float32, H_InOutConnector float32,
	lamp4 float32, lamp6 float32,
	horseShoe float32, centralHeadShaft float32, sideHeadShaft float32,
	oneWay80X804cm float32, OneWay80X803cmBushan float32,
	towWay80X80 float32, windowPicket float32,
) (_2cm float32, _3cm float32, _4cm float32, _5cm float32, _6cm float32, _10cm float32, halfThread float32) {
	_2cm, _3cm, _4cm, _5cm, _6cm, _10cm, halfThread = 0, 0, 0, 0, 0, 0, 0
	//For Gathik
	_2cm += 2 * bindingGathic
	_3cm += 4 * bindingGathic
	_4cm += 1 * bindingGathic
	_6cm += 2 * bindingGathic
	//excel
	_5cm += 2 * excel
	//rowing
	_5cm += 2 * rowing
	//Golpich
	_4cm += 3 * golpich
	//H connector
	_4cm += 3 * H_InOutConnector
	//Lamp 4
	_6cm += lamp4
	//lamp 6
	_5cm += lamp6
	//horse shoe
	_10cm += 2 * horseShoe
	//central head shaft
	_3cm += 12 * centralHeadShaft
	_4cm += 4 * centralHeadShaft
	//side head shaft
	_3cm += 8 * sideHeadShaft
	_4cm += 3 * sideHeadShaft
	//one way connector
	_3cm += oneWay80X804cm
	_3cm += OneWay80X803cmBushan
	//tow-way connector
	_3cm += 2 * towWay80X80
	//window picket
	_5cm += 2 * windowPicket
	halfThread += 1 * windowPicket
	return
}
func getNumber(resultOfCalculates map[string]float32, key string) float32 {
	if val, ok := resultOfCalculates[key]; ok {
		return val
	} else {
		return 0
	}
}
func getPrice(config interface{}, slug string, quantity float32) float32 {

	if config != "" {
		var cnf map[string]interface{}
		switch config.(type) {
		case string:
			err := json.Unmarshal([]byte(config.(string)), &cnf)
			if err != nil {
				panic(err)
			}
			break
		default:

			temp, err := json.Marshal(config)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(temp, &cnf)
			if err != nil {
				panic(err)
			}
		}

		switch slug {
		case "pipe":
			switch cnf["galvanize"] {
			case "fabric":
				fmt.Println(slug)
				return quantity * getFabricPipePrice(cnf, slug)
			case "warm":
				return quantity * getPipeWarmPrice(cnf, slug)
			}
		}

	}
	return 0
}
func getPriceFromRequest(c *gin.Context) interface{} {
	var configDTO struct {
		ElementSlug string      `json:"element_slug"`
		Config      interface{} `json:"config"`
		Quantity    float32     `json:"quantity"`
	}
	err := c.BindJSON(&configDTO)
	if err != nil {
		panic(err)
	}
	price := getPrice(configDTO.Config, configDTO.ElementSlug, configDTO.Quantity)
	return struct {
		Price float32 `json:"price"`
	}{Price: price}
}
func elementFactory(resultOfCalculates map[string]float32, configs *map[string]string) []Category {
	khorshidiNumber := getNumber(resultOfCalculates, "khorshidi")
	pipeConfig := (*configs)["pipe"]
	var x = []Category{
		{"لوله ها و المان ها", []Element{
			{"ستون ها", getNumber(resultOfCalculates, "shaft"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "shaft")), pipeConfig, "shaft"},
			{"کمان ها", getNumber(resultOfCalculates, "bow"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "bow")), pipeConfig, "bow"},
			{"وتر ها", getNumber(resultOfCalculates, "chord"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "chord")), pipeConfig, "chord"},
			{"المان 267cm", getNumber(resultOfCalculates, "267cmElements"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "267cmElements")), pipeConfig, "267cmElements"},
			{"المان 176cm", getNumber(resultOfCalculates, "176cmElements"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "176cmElements")), pipeConfig, "176cmElements"},
			{"المان 150cm", getNumber(resultOfCalculates, "150cmElements"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "150cmElements")), pipeConfig, "150cmElements"},
			{"کانکتور مرکزی", getNumber(resultOfCalculates, "centralConnector"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "centralConnector")), pipeConfig, "centralConnector"},
			//{" المان شبکه ای", getNumber(resultOfCalculates,"shaft" ), "عدد"},
			{"خورشیدی", khorshidiNumber, "عدد", getPrice(pipeConfig, "pipe", khorshidiNumber), pipeConfig, "khorshidi"},
			{"بادبند های X", getNumber(resultOfCalculates, "windBreaker"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "windBreaker")), pipeConfig, "windBreaker"},
			{"سفت کن قبل پنجره", getNumber(resultOfCalculates, "hardenerBeforeWindow"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "hardenerBeforeWindow")), pipeConfig, "hardenerBeforeWindow"},
			{"ستون های فرعی", getNumber(resultOfCalculates, "secondaryShaft"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "secondaryShaft")), pipeConfig, "secondaryShaft"},
			{"سر ستون کناری", getNumber(resultOfCalculates, "sideHeadShaft"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "sideHeadShaft")), pipeConfig, "sideHeadShaft"},
			{"سرستون میانی", getNumber(resultOfCalculates, "centralHeadShaft"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "centralHeadShaft")), pipeConfig, "centralHeadShaft"},
			{"پروفیل لاگینگ", getNumber(resultOfCalculates, "locking"), "متر", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "locking")), pipeConfig, "locking"},
			{"فنر", getNumber(resultOfCalculates, "spring"), "متر", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "spring")), pipeConfig, "spring"},
			{"ناودان کناری", getNumber(resultOfCalculates, "sideGutter"), "تعداد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "sideGutter")), pipeConfig, "sideGutter"},
			{"ناودان میانی", getNumber(resultOfCalculates, "centralGutter"), "تعداد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "centralGutter")), pipeConfig, "centralGutter"},
		}},
		{"هوک ها", []Element{
			{"کمان دوم به ستون فرعی", getNumber(resultOfCalculates, "secondToShaft"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "secondToShaft")), pipeConfig, "secondToShaft"},
			{"هوک کمان مورب", getNumber(resultOfCalculates, "diagonal"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "diagonal")), pipeConfig, "diagonal"},
			{"هوک کمان اول به دوم", getNumber(resultOfCalculates, "firstBowToSecond"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "firstBowToSecond")), pipeConfig, "firstBowToSecond"},
		}},
		{"بست ها", []Element{
			{"بست گاتیک", getNumber(resultOfCalculates, "bindingGathic"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "bindingGathic")), pipeConfig, "bindingGathic"},
			{"بست 80x 80 یک طرفه 4cm", getNumber(resultOfCalculates, "oneWay80X804cm"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "oneWay80X804cm")), pipeConfig, "oneWay80X804cm"},
			{"بست 80x 80 یک طرفه 3cm+ بوشن", getNumber(resultOfCalculates, "OneWay80X803cmBushan"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "oneWay80X804cm")), pipeConfig, "OneWay80X803cmBushan"},
			{"بست 80x 80 دوطرفه", getNumber(resultOfCalculates, "towWay80X80"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "towWay80X80")), pipeConfig, "towWay80X80"},
			{"بست لامپی 4", getNumber(resultOfCalculates, "lamp4"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "lamp4")), pipeConfig, "lamp4"},
			{"بست لامپی 6", getNumber(resultOfCalculates, "lamp6"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "lamp6")), pipeConfig, "lamp6"},
			{"بست سفت کن و بوشن", getNumber(resultOfCalculates, "hardenerBushen"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "hardenerBushen")), pipeConfig, "hardenerBushen"},
			{"بست سفت کن کنار", getNumber(resultOfCalculates, "sideHardener"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "sideHardener")), pipeConfig, "sideHardener"},
			{"بست سفت کن سر و ته", getNumber(resultOfCalculates, "headAndTailHardener"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "headAndTailHardener")), pipeConfig, "headAndTailHardener"},
			{"پیچ نعل اسبی", getNumber(resultOfCalculates, "horseShoe"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "horseShoe")), pipeConfig, "horseShoe"},
			{"پیچ شیروانی", getNumber(resultOfCalculates, "gableScrew"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "gableScrew")), pipeConfig, "gableScrew"},
			{"بست رابط درونی وبیرونی H", getNumber(resultOfCalculates, "H_InOutConnector"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "H_InOutConnector")), pipeConfig, "H_InOutConnector"},
		}},
		{"پنجره", []Element{
			{"رک", getNumber(resultOfCalculates, "rack"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "rack")), pipeConfig, "rack"},
			{"پینیون", getNumber(resultOfCalculates, "pinion"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "pinion")), pipeConfig, "pinion"},
			{"سفت کن زیرپنجره", getNumber(resultOfCalculates, "hardenerUnderTheWindow"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "hardenerUnderTheWindow")), "", "hardenerUnderTheWindow"},
			{"لوله شفت", getNumber(resultOfCalculates, "shaftPipe"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "shaftPipe")), pipeConfig, "shaftPipe"},
			{"رابط لوله شفت", getNumber(resultOfCalculates, "shaftPipeConnector"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "shaftPipeConnector")), pipeConfig, "shaftPipeConnector"},
			{"رابط گلپیچ", getNumber(resultOfCalculates, "golpich"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "golpich")), pipeConfig, "golpich"},
			{" تعداد پروفیل H سرپنجره", getNumber(resultOfCalculates, "headOfWindowH"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "headOfWindowH")), pipeConfig, "headOfWindowH"},
			{"دستک پنجره", getNumber(resultOfCalculates, "windowPicket"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "windowPicket")), pipeConfig, "windowPicket"},
			{"بست LOF", getNumber(resultOfCalculates, "LOF"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "LOF")), pipeConfig, "LOF"},
			{"بست اکسل", getNumber(resultOfCalculates, "excel"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "excel")), pipeConfig, "excel"},
			{"بست پارویی", getNumber(resultOfCalculates, "rowing"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "rowing")), pipeConfig, "rowing"},
		}},
		{"پیچ و مهره", []Element{
			{"۲ سانتی", getNumber(resultOfCalculates, "bolt_2cm"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "bolt_2cm")), pipeConfig, "bolt_2cm"},
			{"۳ سانتی", getNumber(resultOfCalculates, "bolt_3cm"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "bolt_3cm")), pipeConfig, "bolt_3cm"},
			{"۴ سانتی", getNumber(resultOfCalculates, "bolt_4cm"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "bolt_4cm")), pipeConfig, "bolt_4cm"},
			{"۵ سانتی", getNumber(resultOfCalculates, "bolt_5cm"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "bolt_5cm")), pipeConfig, "bolt_5cm"},
			{"۶ سانتی", getNumber(resultOfCalculates, "bolt_6cm"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "bolt_6cm")), pipeConfig, "bolt_6cm"},
			{"۱۰ سانتی", getNumber(resultOfCalculates, "bolt_10cm"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "bolt_10cm")), pipeConfig, "bolt_10cm"},
			{"نیم رزوه با مهره کاسه نمدی", getNumber(resultOfCalculates, "bolt_halfThread"), "عدد", getPrice(pipeConfig, "pipe", getNumber(resultOfCalculates, "bolt_halfThread")), pipeConfig, "bolt_halfThread"},
		}},
	}
	return x
}
