package distribution

func MockSeedDistribution(i int) int {
	var t = [52]int{6557656767793228214, 7001713455464880029, 5641557149578671256, 7340620975452528024,
		9018662454251328848, 7081691525610742019, 7300482329487060319, 800193923855012045,
		4178860451557774013, 370226610021623984, 7084967646861838833, 4734535648075306608,
		4371804806294316792, 7679108395215856930, 184060070512281621,
		1543857178756158216, 2962082534040367486, 2013862716327628120, 7382504487626277544,
		4125419142925215916, 1481607125531674309, 553873083935818906, 1933624817553623601,
		7844300803605428424, 5193646226957662729, 4978059754121862968, 4072967079906883225,
		9194287614505867302, 825706703135595394, 5381191702175920675,
		6011521299003686509, 7356535707923529802, 7416220184612542210,
		4462659083952081245, 7570491316003320002, 8503951705800838088, 4566708990123589583, 3191744104973416895,
		4985228858347885032, 3648700154125034516, 630200367234550479, 228680610405627893,
		5018877883163560525, 7911457462280715148, 2443576849415948692,
		5917890275451941034, 5530737869665678128, 5913659668741068281, 7484951419787565398,
		9080017362341776237, 4122383095093646999, 8384499133450967238}

	return t[i]
}
func MockTabResultDistribution() ArrayDistribution {
	var t ArrayDistribution
	var t0 = [13]int{28, 36, 5, 9, 13, 17, 29, 33, 37, 10, 7, 11, 23}
	var t1 = [13]int{32, 44, 48, 25, 49, 2, 30, 42, 46, 50, 3, 15, 47}
	var t2 = [13]int{0, 4, 8, 12, 20, 24, 40, 45, 14, 18, 27, 31, 35}
	var t3 = [13]int{16, 1, 21, 41, 6, 22, 26, 34, 38, 19, 39, 43, 51}
	t[0] = t0
	t[1] = t1
	t[2] = t2
	t[3] = t3
	return t
}

func MockResultDistribution(i int, j int) int {
	var t ArrayDistribution
	t = MockTabResultDistribution()
	return t[i][j]
}

func MockImprimeMain(orientation int) string {
	var result string
	if orientation == 0 {
		result = Orientation(0) + " (2):9 T, Valet T, 3 K, 4 K, 5 K, 6 K, 9 K, 10 K, Valet K, 4 C, 3 P, 4 P, 7 P"
	}
	if orientation == 1 {
		result = Orientation(1) + " (23):10 T, Roi T, As T, 8 K, As K, 2 C, 9 C, Dame C, Roi C, As C, 2 P, 5 P, Roi P"
	}
	if orientation == 2 {
		result = Orientation(2) + " (5):2 T, 3 T, 4 T, 5 T, 7 T, 8 T, Dame T, Roi K, 5 C, 6 C, 8 P, 9 P, 10 P"
	}
	if orientation == 3 {
		result = Orientation(3) + " (10):6 T, 2 K, 7 K, Dame K, 3 C, 7 C, 8 C, 10 C, Valet C, 6 P, Valet P, Dame P, As P"
	}
	return result
}