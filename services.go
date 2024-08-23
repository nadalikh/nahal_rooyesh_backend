package main

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

func elementFactory(resultOfCalculates map[string]float32) []Category {
	var x = []Category{
		{"لوله ها و المان ها", []Element{
			{"ستون ها", getNumber(resultOfCalculates, "shaft"), "عدد"},
			{"کمان ها", getNumber(resultOfCalculates, "bow"), "عدد"},
			{"وتر ها", getNumber(resultOfCalculates, "chord"), "عدد"},
			{"المان 267cm", getNumber(resultOfCalculates, "267cmElements"), "عدد"},
			{"المان 176cm", getNumber(resultOfCalculates, "176cmElements"), "عدد"},
			{"المان 150cm", getNumber(resultOfCalculates, "150cmElements"), "عدد"},
			{"کانکتور مرکزی", getNumber(resultOfCalculates, "centralConnector"), "عدد"},
			//{" المان شبکه ای", getNumber(resultOfCalculates,"shaft" ), "عدد"},
			{" خورشیدی", getNumber(resultOfCalculates, "khorshidi"), "عدد"},
			{" بادبند های X", getNumber(resultOfCalculates, "windBreaker"), "عدد"},
			{" سفت کن قبل پنجره", getNumber(resultOfCalculates, "hardenerBeforeWindow"), "عدد"},
			{" ستون های فرعی", getNumber(resultOfCalculates, "secondaryShaft"), "عدد"},
			{"سر ستون کناری", getNumber(resultOfCalculates, "sideHeadShaft"), "عدد"},
			{"سرستون میانی", getNumber(resultOfCalculates, "centralHeadShaft"), "عدد"},
			{" پروفیل لاگینگ", getNumber(resultOfCalculates, "locking"), "متر"},
			{" فنر", getNumber(resultOfCalculates, "spring"), "متر"},
			{" ناودان کناری", getNumber(resultOfCalculates, "sideGutter"), "تعداد"},
			{" ناودان میانی", getNumber(resultOfCalculates, "centralGutter"), "تعداد"},
		}},
		{"هوک ها", []Element{
			{"کمان دوم به ستون فرعی", getNumber(resultOfCalculates, "secondToShaft"), "عدد"},
			{"هوک کمان مورب", getNumber(resultOfCalculates, "diagonal"), "عدد"},
			{"هوک کمان اول به دوم", getNumber(resultOfCalculates, "firstBowToSecond"), "عدد"},
		}},
		{"بست ها", []Element{
			{"بست گاتیک", getNumber(resultOfCalculates, "bindingGathic"), "عدد"},
			{"بست 80x 80 یک طرفه 4cm", getNumber(resultOfCalculates, "oneWay80X804cm"), "عدد"},
			{"بست 80x 80 یک طرفه 3cm+ بوشن", getNumber(resultOfCalculates, "OneWay80X803cmBushan"), "عدد"},
			{"بست 80x 80 دوطرفه", getNumber(resultOfCalculates, "towWay80X80"), "عدد"},
			{"بست لامپی 4", getNumber(resultOfCalculates, "lamp4"), "عدد"},
			{"بست لامپی 6", getNumber(resultOfCalculates, "lamp6"), "عدد"},
			{"بست سفت کن و بوشن", getNumber(resultOfCalculates, "hardenerBushen"), "عدد"},
			{"بست سفت کن کنار", getNumber(resultOfCalculates, "sideHardener"), "عدد"},
			{"بست سفت کن سر و ته", getNumber(resultOfCalculates, "headAndTailHardener"), "عدد"},
			{"پیچ نعل اسبی", getNumber(resultOfCalculates, "horseShoe"), "عدد"},
			{"پیچ شیروانی", getNumber(resultOfCalculates, "gableScrew"), "عدد"},
			{"بست رابط درونی وبیرونی H", getNumber(resultOfCalculates, "H_InOutConnector"), "عدد"},
		}},
		{"پنجره", []Element{
			{"رک", getNumber(resultOfCalculates, "rack"), "عدد"},
			{"پینیون", getNumber(resultOfCalculates, "pinion"), "عدد"},
			{"سفت کن زیرپنجره", getNumber(resultOfCalculates, "hardenerUnderTheWindow"), "عدد"},
			{"لوله شفت", getNumber(resultOfCalculates, "shaftPipe"), "عدد"},
			{"رابط لوله شفت", getNumber(resultOfCalculates, "shaftPipeConnector"), "عدد"},
			{"رابط گلپیچ", getNumber(resultOfCalculates, "golpich"), "عدد"},
			{" تعداد پروفیل H سرپنجره", getNumber(resultOfCalculates, "headOfWindowH"), "عدد"},
			{"دستک پنجره", getNumber(resultOfCalculates, "windowPicket"), "عدد"},
			{"بست LOF", getNumber(resultOfCalculates, "LOF"), "عدد"},
			{"بست اکسل", getNumber(resultOfCalculates, "excel"), "عدد"},
			{"بست پارویی", getNumber(resultOfCalculates, "rowing"), "عدد"},
		}},
		{"پیچ و مهره", []Element{
			{"۲ سانتی", getNumber(resultOfCalculates, "bolt_2cm"), "عدد"},
			{"۳ سانتی", getNumber(resultOfCalculates, "bolt_3cm"), "عدد"},
			{"۴ سانتی", getNumber(resultOfCalculates, "bolt_4cm"), "عدد"},
			{"۵ سانتی", getNumber(resultOfCalculates, "bolt_5cm"), "عدد"},
			{"۶ سانتی", getNumber(resultOfCalculates, "bolt_6cm"), "عدد"},
			{"۱۰ سانتی", getNumber(resultOfCalculates, "bolt_10cm"), "عدد"},
			{"نیم رزوه با مهره کاسه نمدی", getNumber(resultOfCalculates, "bolt_halfThread"), "عدد"},
		}},
	}
	return x
}
