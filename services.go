package main

import (
	"encoding/json"
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
func getPrice(config, slug string, quantity float32) float32 {
	if config != "" {
		var cnf map[string]interface{}
		err := json.Unmarshal([]byte(config), &cnf)
		if err != nil {
			panic(err)
		}
		switch cnf["galvanize"] {
		case "warm":
			return quantity * getWarmPrice(slug)
		}
	}
	return 0
}
func elementFactory(resultOfCalculates map[string]float32, configs *map[string]string) []Category {
	khorshidiNumber := getNumber(resultOfCalculates, "khorshidi")
	khorshidiCnf := (*configs)["khorshidi"]
	var x = []Category{
		{"لوله ها و المان ها", []Element{
			{"ستون ها", getNumber(resultOfCalculates, "shaft"), "عدد", getPrice("", "", 0), ""},
			{"کمان ها", getNumber(resultOfCalculates, "bow"), "عدد", getPrice("", "", 0), ""},
			{"وتر ها", getNumber(resultOfCalculates, "chord"), "عدد", getPrice("", "", 0), ""},
			{"المان 267cm", getNumber(resultOfCalculates, "267cmElements"), "عدد", getPrice("", "", 0), ""},
			{"المان 176cm", getNumber(resultOfCalculates, "176cmElements"), "عدد", getPrice("", "", 0), ""},
			{"المان 150cm", getNumber(resultOfCalculates, "150cmElements"), "عدد", getPrice("", "", 0), ""},
			{"کانکتور مرکزی", getNumber(resultOfCalculates, "centralConnector"), "عدد", getPrice("", "", 0), ""},
			//{" المان شبکه ای", getNumber(resultOfCalculates,"shaft" ), "عدد"},
			{"خورشیدی", khorshidiNumber, "عدد", getPrice(khorshidiCnf, "khorshidi", khorshidiNumber), khorshidiCnf},
			{"بادبند های X", getNumber(resultOfCalculates, "windBreaker"), "عدد", getPrice("", "", 0), ""},
			{"سفت کن قبل پنجره", getNumber(resultOfCalculates, "hardenerBeforeWindow"), "عدد", getPrice("", "", 0), ""},
			{"ستون های فرعی", getNumber(resultOfCalculates, "secondaryShaft"), "عدد", getPrice("", "", 0), ""},
			{"سر ستون کناری", getNumber(resultOfCalculates, "sideHeadShaft"), "عدد", getPrice("", "", 0), ""},
			{"سرستون میانی", getNumber(resultOfCalculates, "centralHeadShaft"), "عدد", getPrice("", "", 0), ""},
			{"پروفیل لاگینگ", getNumber(resultOfCalculates, "locking"), "متر", getPrice("", "", 0), ""},
			{"فنر", getNumber(resultOfCalculates, "spring"), "متر", getPrice("", "", 0), ""},
			{"ناودان کناری", getNumber(resultOfCalculates, "sideGutter"), "تعداد", getPrice("", "", 0), ""},
			{"ناودان میانی", getNumber(resultOfCalculates, "centralGutter"), "تعداد", getPrice("", "", 0), ""},
		}},
		{"هوک ها", []Element{
			{"کمان دوم به ستون فرعی", getNumber(resultOfCalculates, "secondToShaft"), "عدد", getPrice("", "", 0), ""},
			{"هوک کمان مورب", getNumber(resultOfCalculates, "diagonal"), "عدد", getPrice("", "", 0), ""},
			{"هوک کمان اول به دوم", getNumber(resultOfCalculates, "firstBowToSecond"), "عدد", getPrice("", "", 0), ""},
		}},
		{"بست ها", []Element{
			{"بست گاتیک", getNumber(resultOfCalculates, "bindingGathic"), "عدد", getPrice("", "", 0), ""},
			{"بست 80x 80 یک طرفه 4cm", getNumber(resultOfCalculates, "oneWay80X804cm"), "عدد", getPrice("", "", 0), ""},
			{"بست 80x 80 یک طرفه 3cm+ بوشن", getNumber(resultOfCalculates, "OneWay80X803cmBushan"), "عدد", getPrice("", "", 0), ""},
			{"بست 80x 80 دوطرفه", getNumber(resultOfCalculates, "towWay80X80"), "عدد", getPrice("", "", 0), ""},
			{"بست لامپی 4", getNumber(resultOfCalculates, "lamp4"), "عدد", getPrice("", "", 0), ""},
			{"بست لامپی 6", getNumber(resultOfCalculates, "lamp6"), "عدد", getPrice("", "", 0), ""},
			{"بست سفت کن و بوشن", getNumber(resultOfCalculates, "hardenerBushen"), "عدد", getPrice("", "", 0), ""},
			{"بست سفت کن کنار", getNumber(resultOfCalculates, "sideHardener"), "عدد", getPrice("", "", 0), ""},
			{"بست سفت کن سر و ته", getNumber(resultOfCalculates, "headAndTailHardener"), "عدد", getPrice("", "", 0), ""},
			{"پیچ نعل اسبی", getNumber(resultOfCalculates, "horseShoe"), "عدد", getPrice("", "", 0), ""},
			{"پیچ شیروانی", getNumber(resultOfCalculates, "gableScrew"), "عدد", getPrice("", "", 0), ""},
			{"بست رابط درونی وبیرونی H", getNumber(resultOfCalculates, "H_InOutConnector"), "عدد", getPrice("", "", 0), ""},
		}},
		{"پنجره", []Element{
			{"رک", getNumber(resultOfCalculates, "rack"), "عدد", getPrice("", "", 0), ""},
			{"پینیون", getNumber(resultOfCalculates, "pinion"), "عدد", getPrice("", "", 0), ""},
			{"سفت کن زیرپنجره", getNumber(resultOfCalculates, "hardenerUnderTheWindow"), "عدد", getPrice("", "", 0), ""},
			{"لوله شفت", getNumber(resultOfCalculates, "shaftPipe"), "عدد", getPrice("", "", 0), ""},
			{"رابط لوله شفت", getNumber(resultOfCalculates, "shaftPipeConnector"), "عدد", getPrice("", "", 0), ""},
			{"رابط گلپیچ", getNumber(resultOfCalculates, "golpich"), "عدد", getPrice("", "", 0), ""},
			{" تعداد پروفیل H سرپنجره", getNumber(resultOfCalculates, "headOfWindowH"), "عدد", getPrice("", "", 0), ""},
			{"دستک پنجره", getNumber(resultOfCalculates, "windowPicket"), "عدد", getPrice("", "", 0), ""},
			{"بست LOF", getNumber(resultOfCalculates, "LOF"), "عدد", getPrice("", "", 0), ""},
			{"بست اکسل", getNumber(resultOfCalculates, "excel"), "عدد", getPrice("", "", 0), ""},
			{"بست پارویی", getNumber(resultOfCalculates, "rowing"), "عدد", getPrice("", "", 0), ""},
		}},
		{"پیچ و مهره", []Element{
			{"۲ سانتی", getNumber(resultOfCalculates, "bolt_2cm"), "عدد", getPrice("", "", 0), ""},
			{"۳ سانتی", getNumber(resultOfCalculates, "bolt_3cm"), "عدد", getPrice("", "", 0), ""},
			{"۴ سانتی", getNumber(resultOfCalculates, "bolt_4cm"), "عدد", getPrice("", "", 0), ""},
			{"۵ سانتی", getNumber(resultOfCalculates, "bolt_5cm"), "عدد", getPrice("", "", 0), ""},
			{"۶ سانتی", getNumber(resultOfCalculates, "bolt_6cm"), "عدد", getPrice("", "", 0), ""},
			{"۱۰ سانتی", getNumber(resultOfCalculates, "bolt_10cm"), "عدد", getPrice("", "", 0), ""},
			{"نیم رزوه با مهره کاسه نمدی", getNumber(resultOfCalculates, "bolt_halfThread"), "عدد", getPrice("", "", 0), ""},
		}},
	}
	return x
}
