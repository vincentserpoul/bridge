package main

import (
	"fmt"

	"github.com/jeanjacquesserpoul/bridge-distribution/bridge/distribution"
)

func main() {
	var listeDesCartes distribution.Distribution
	var ia distribution.SDistribution
	var somme int
	var err error
	var r string

	err = distribution.Melange(&listeDesCartes, ia, 4)
	r = distribution.MontreMain(&listeDesCartes, 0) + "\n"
	r += distribution.MontreMain(&listeDesCartes, 1) + "\n"
	r += distribution.MontreMain(&listeDesCartes, 2) + "\n"
	r += distribution.MontreMain(&listeDesCartes, 3)
	fmt.Println(r)
	somme = distribution.TotalPoints(&listeDesCartes, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Total des points:%d", somme)
	}

}
