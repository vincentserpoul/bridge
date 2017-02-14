package distribution

import (
	"strconv"
	"testing"
)

type fakeAlea struct{}

func (fakeAlea) Alea(nombreDeCartes int) []int {
	var alea []int
	alea = make([]int, nombreDeCartes)
	for i := 0; i < nombreDeCartes; i++ {
		alea[i] = MockSeedDistribution(i)
	}

	return alea
}

func convertArrayToString(d ArrayDistribution) string {
	var chaine string
	chaine = ""
	for i := 0; i <= 3; i++ {
		for j := 0; j < 13; j++ {
			chaine += strconv.Itoa(d[i][j]) + "|"
		}
	}
	return chaine

}

func TestMelange(t *testing.T) {
	var listeDesCartes Distribution
	var ta fakeAlea
	_ = Melange(&listeDesCartes, ta, 4)
	a := convertArrayToString(listeDesCartes.cartes)
	b := convertArrayToString(MockTabResultDistribution())
	if a != b {
		t.Errorf("Mélange: %s obtenu au lieu de %s", a, b)
	}

}

func TestMontreMain(t *testing.T) {
	var listeDesCartes Distribution
	var ta fakeAlea
	for i := 0; i <= 3; i++ {
		_ = Melange(&listeDesCartes, ta, 4)
		a := MontreMain(&listeDesCartes, i)
		b := MockImprimeMain(i)
		if a != b {
			t.Errorf("Montre Main %s: %s obtenu au lieu de %s", Orientation(i), a, b)
		}
	}

}
func BenchmarkMelange(b *testing.B) {
	var listeDesCartes Distribution
	var ta fakeAlea
	for n := 0; n < b.N; n++ {
		_ = Melange(&listeDesCartes, ta, 4)
	}
}
