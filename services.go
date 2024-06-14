package main

func calculateShafts(land Land) uint16 {
	shaft := ((land.Length / 3) + 1) * ((land.Width * 10 / 96) + 1)
	return shaft
}
func calculateArcAndChord(land Land) (uint16, uint16) {
	bow := ((land.Length / 3) + 1) * (land.Width * 10 / 96)
	chord := bow - ((land.Width * 10 / 96) * 2)
	return bow, chord
}
func calculateElementsOfChord(chord uint16) uint16 {
	const elementCategory = 3 //SMALL, BIG, BOW TO SHAFT
	const elementCategoryPerChord = 2
	return chord * elementCategory * elementCategoryPerChord
}
func calculateKhorshidi(chord uint16) uint16 {
	const numberOfKhorshidiPerChord = 2
	return chord * numberOfKhorshidiPerChord
}
func calculateCentralConnector(land Land) uint16 {
	return land.Length / 6 * (land.Width * 10 / 96)
}
func calculateHardenerBeforeWindow(land Land) uint16 {
	return land.Length / 6 * (land.Width * 10 / 96)
}
func calculateWindBreaker(land Land) uint16 {
	return ((land.Width * 10 / 96) + 1) * 2
}
func calculateSecondaryShaft(land Land) uint16 {
	const secondaryShaftPerAperture = 4
	return (land.Width * 10 / 96) * secondaryShaftPerAperture
}
func calculateTheSideGutter(land Land) uint16 {
	return land.Length / 6 * 2
}
func calculateCentralGutter(land Land, sideGutter uint16) uint16 {
	return ((land.Length / 6) * ((land.Width * 10 / 96) + 1)) - sideGutter
}
func calculateTheSideHeadShaft(land Land) uint16 {
	return 2 * land.Length / 3
}
func calculateTheCentralHeadShaft(land Land, sideHeadShaft uint16) uint16 {
	return ((land.Length / 3) * ((land.Width * 10 / 96) + 1)) - sideHeadShaft
}
func calculateBindingGothic(bow uint16) uint16 {
	return bow
}
func calculateHooks(land Land) (uint16, uint16, uint16) {
	firstBowToSecond := (land.Width * 10 / 96) * 4
	diagonal := (land.Width * 10 / 96) * 4
	secondToShaft := (land.Width * 10 / 96) * 4
	return firstBowToSecond, diagonal, secondToShaft
}
func calculateWindowPicket(bow uint16) uint16 {
	return bow
}
func calculateRack(bow uint16) uint16 {
	return bow
}
func calculate6lamp(firstBowToSecond uint16, secondToShaft uint16, chord uint16) uint16 {
	return (firstBowToSecond * 2) + secondToShaft + (chord * 4)
}
func calculate4lamp(khorshidi uint16, chord uint16) uint16 {
	return (chord * 2) + chord + khorshidi
}
func calculate80X80OneWay(land Land, windbreakerShafts uint16, secondToShaft uint16) uint16 {
	sideShaft := 2 * ((land.Length / 3) + 1)
	centralRows := (land.Width * 10 / 96)
	return 1 + (windbreakerShafts * 2) +
		secondToShaft +
		(sideShaft * 2) -
		(centralRows)
}
func calculate80X80TowWay(land Land, shaft uint16) uint16 {
	sideShaft := 2 * ((land.Length / 3) + 1)
	centralRows := (land.Length / 3) + 1
	return (shaft - sideShaft) - (centralRows * 2)
}
func calculateRowing(windowPicket uint16) uint16 {
	return windowPicket
}
func calculateLOF(windowPicket uint16) uint16 {
	return windowPicket
}
func calculateExcel(windowPicket uint16) uint16 {
	return windowPicket
}
func calculatePinion(windowPicket uint16) uint16 {
	return windowPicket
}
func calculateHeadOfWindowH(land Land) uint16 {
	return land.Length / 6
}
func calculateHardenerUnderTheWindow(land Land) uint16 {
	return land.Length / 6
}
func calculateShaftPipe(land Land) uint16 {
	return land.Length / 6
}
func calculateShaftPipeConnector(land Land) uint16 {
	return (land.Length / 6) - 1
}
func calculateHardenerBushen(windowPicket uint16) uint16 {
	return windowPicket
}
func calculate_H_InOutConnector(headOfWindowH uint16) uint16 {
	return headOfWindowH - 1
}
func calculateKapage(rowing uint16) uint16 {
	return rowing * 2
}
func calculateLocking(land Land, sideGutter uint16, centralGutter uint16) uint16 {
	return (6 * 2 * (sideGutter + centralGutter)) +
		((land.Width * 10 / 96) * 11 * 2) + 20
}
func calculateSpring(land Land, headOfWindowH uint16, locking uint16) uint16 {
	return ((land.Width * 10 / 96) * 2) +
		(headOfWindowH * 6) + locking
}
func getNumber(resultOfCalculates map[string]uint16, key string) uint16 {
	if val, ok := resultOfCalculates[key]; ok {
		return val
	} else {
		return 0
	}

}

func elementFactory(resultOfCalculates map[string]uint16) []Category {
	var x = []Category{
		{"لوله ها و المان ها", []Element{
			{"ستون ها", getNumber(resultOfCalculates, "shaft"), "عدد"},
			{"کمان ها", getNumber(resultOfCalculates, "bow"), "عدد"},
			{"وتر ها", getNumber(resultOfCalculates, "chord"), "عدد"},
			{"المان ها", getNumber(resultOfCalculates, "elementsOfchord"), "عدد"},
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
			{" ناودان کناری", getNumber(resultOfCalculates, "sideGutter"), "متر"},
			{" ناودان میانی", getNumber(resultOfCalculates, "centralGutter"), "متر"},
		}},
		{"هوک ها", []Element{
			{"کمان دوم به ستون فرعی", getNumber(resultOfCalculates, "secondToShaft"), "عدد"},
			{"هوک کمان مورب", getNumber(resultOfCalculates, "diagonal"), "عدد"},
			{"هوک کمان اول به دوم", getNumber(resultOfCalculates, "firstBowToSecond"), "عدد"},
		}},
		{"بست ها", []Element{
			{"بست گاتیک", getNumber(resultOfCalculates, "bindingGathic"), "عدد"},
			{"بست 80x 80 یک طرفه", getNumber(resultOfCalculates, "oneWay80X80"), "عدد"},
			{"بست 80x 80 دوطرفه", getNumber(resultOfCalculates, "towWay80X80"), "عدد"},
			{"بست پارویی", getNumber(resultOfCalculates, "rowing"), "عدد"},
			{"بست لامپی 4", getNumber(resultOfCalculates, "lamp4"), "عدد"},
			{"بست لامپی 6", getNumber(resultOfCalculates, "lamp6"), "عدد"},
			{"بست LOF", getNumber(resultOfCalculates, "LOF"), "عدد"},
			{"بست اکسل", getNumber(resultOfCalculates, "excel"), "عدد"},
			{"بست سفت کن و بوشن", getNumber(resultOfCalculates, "hardenerBushen"), "عدد"},
			{"بسط رابط درونی وبیرونی H", getNumber(resultOfCalculates, "H_InOutConnector"), "عدد"},
		}},
		{"پنجره", []Element{
			{"رک", getNumber(resultOfCalculates, "rack"), "عدد"},
			{"پینیون", getNumber(resultOfCalculates, "pinion"), "عدد"},
			{"سفت کن زیرپنجره", getNumber(resultOfCalculates, "hardenerUnderTheWindow"), "عدد"},
			{"لوله شفت", getNumber(resultOfCalculates, "shaftPipe"), "عدد"},
			{"رابط لوله شفت", getNumber(resultOfCalculates, "shaftPipeConnector"), "عدد"},
			{"رابط کاپیج", getNumber(resultOfCalculates, "kapage"), "عدد"},
			{" تعداد پروفیل H سرپنجره", getNumber(resultOfCalculates, "headOfWindowH"), "عدد"},
			{"دستک پنجره", getNumber(resultOfCalculates, "windowPicket"), "عدد"},
		}},
	}
	return x
}
