package distribution

import "strconv"

//GetValeur ...
func GetValeur(v int) int {
	var result int
	result = 0
	switch {
	case v <= 8:
		result = 0
	case v == 9:
		result = 1
	case v == 10:
		result = 2
	case v == 11:
		result = 3
	case v == 12:
		result = 4
	}
	return result
}

//GetFacial ...
func GetFacial(valCarte int) string {
	var result string
	result = strconv.Itoa(valCarte)
	switch {
	case valCarte <= 8:
		result = strconv.Itoa(valCarte + 2)
	case valCarte == 9:
		result = "Valet"
	case valCarte == 10:
		result = "Dame"
	case valCarte == 11:
		result = "Roi"
	case valCarte == 12:
		result = "As"
	}
	return (result)
}

//GetCouleur ...
func GetCouleur(valCouleur int) string {
	var result string
	result = "T"
	switch valCouleur {
	case 0:
		result = "T"
	case 1:
		result = "K"
	case 2:
		result = "C"
	case 3:
		result = "P"
	}
	return (result)

}

//SetCouleurs ...
func SetCouleurs(c int) string {
	var result string
	result = "00"
	switch c {
	case 0:
		result = "00"
	case 1:
		result = "01"
	case 2:
		result = "10"
	case 3:
		result = "11"
	}
	return (result)
}

// Orientation ...
func Orientation(o int) string {
	var result string
	result = "Nord"
	switch o {
	case 0:
		result = "Nord"
	case 1:
		result = "Est"
	case 2:
		result = "Sud"
	case 3:
		result = "Ouest"
	}
	return (result)

}

// ValeurMain ...
func ValeurMain(m [13]int) int {
	var result int
	result = 0
	for i := 0; i <= 12; i++ {
		result += GetValeur(m[i] >> 2)
	}
	return result
}

// GetFacialCarte ...
func GetFacialCarte(v int) string {
	return GetFacial(v>>2) + " " + GetCouleur(v&3)
}
