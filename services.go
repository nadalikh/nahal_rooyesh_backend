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
	centralRows := (land.Length / 3) + 1
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
