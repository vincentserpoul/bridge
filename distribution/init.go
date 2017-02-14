package distribution

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type listeMain struct {
	valeur int
	tri    int
}
type liste []listeMain

func (slice liste) Len() int {
	return len(slice)
}

func (slice liste) Less(i, j int) bool {
	return slice[i].tri < slice[j].tri
}

func (slice liste) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type tabTri struct {
	position int
	valeur   int
}

type randomTab []int

type listeCartes []tabTri

//NCM Nombre de cartes par main
const NCM = 13

func (slice listeCartes) Len() int {
	return len(slice)
}

func (slice listeCartes) Less(i, j int) bool {
	return slice[i].valeur < slice[j].valeur
}

func (slice listeCartes) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//ArrayDistribution ...
type ArrayDistribution [4][NCM]int

//Distribution ...
type Distribution struct {
	cartes ArrayDistribution
}

//IDistribution gestion methode Alea
type IDistribution interface {
	Alea(nombreDeMain int) []int
}

//SDistribution ...
type SDistribution struct{}

//MontreMain ...
func MontreMain(d *Distribution, main int) string {
	var amain, r string
	r = ""
	amain = Orientation(main)
	r += amain + " (" + strconv.Itoa(ValeurMain(d.cartes[main])) + "):"
	for j := 0; j < NCM; j++ {
		r += GetFacialCarte(d.cartes[main][j])
		if j == NCM-1 {
			r += ""
		} else {
			r += ", "
		}

	}
	return r
}

//Alea ...
func (SDistribution) Alea(nombreDeMains int) []int {
	var alea []int
	alea = make([]int, nombreDeMains*NCM)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nombreDeMains*NCM; i++ {
		alea[i] = rand.Int()
	}
	return alea
}

//TotalPoints ...
func TotalPoints(d *Distribution, nombreDeMains int) int {
	var somme, carte int
	somme = 0
	for i := 0; i < nombreDeMains; i++ {
		for j := 0; j < NCM; j++ {
			carte = d.cartes[i][j] >> 2
			somme += GetValeur(carte)
		}
	}

	return somme
}

//Melange ...
func Melange(d *Distribution, ia IDistribution, nombreDeMains int) error {
	var listeMain liste
	var valeurCarte, couleurCarte, triCarte int
	var tab listeCartes
	var listeDesCartes Distribution
	var alea randomTab
	var k, m, v, w, z, somme int
	var x int64
	// Initialisation carte 4 mains triées une par couleur
	for i := 0; i < nombreDeMains; i++ {
		for j := 0; j < NCM; j++ {
			x, _ = strconv.ParseInt(strconv.FormatInt(int64(j), 2)+SetCouleurs(i), 2, 8)
			d.cartes[i][j] = int(x)
		}
	}
	// On mélange
	tab = make(listeCartes, nombreDeMains*NCM)
	// Génération aléatoire
	alea = ia.Alea(nombreDeMains * NCM)
	for i := 0; i < nombreDeMains*NCM; i++ {
		tab[i].valeur = alea[i]
		tab[i].position = i
	}
	//Répartition dans chaque main
	sort.Sort(tab)
	listeDesCartes = *d
	for i := 0; i < nombreDeMains*NCM; i++ {
		m = tab[i].position
		v = m % NCM
		w = m / NCM
		k = i % NCM
		z = i / NCM
		listeDesCartes.cartes[z][k] = d.cartes[w][v]
	}
	*d = listeDesCartes
	//Tri des mains par couleur puis par hauteur
	listeMain = make(liste, NCM)
	for i := 0; i < nombreDeMains; i++ {
		for j := 0; j < NCM; j++ {
			valeurCarte = (d.cartes[i][j]) >> 2
			couleurCarte = (d.cartes[i][j]) & 3
			triCarte = (couleurCarte << 4) + valeurCarte
			listeMain[j].valeur = d.cartes[i][j]
			listeMain[j].tri = triCarte
		}
		sort.Sort(listeMain)
		for j := 0; j < NCM; j++ {
			d.cartes[i][j] = listeMain[j].valeur
		}
	}
	somme = TotalPoints(d, nombreDeMains)
	if somme != 40 {
		return fmt.Errorf("Erreur: Le nombre de points de la donne doit valoir 40")
	}
	return nil

}
