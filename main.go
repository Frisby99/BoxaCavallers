package main

import (
	"math/rand"
	"time"

	"github.com/utrescu/CombatCavallers.Go/combat"
	"github.com/utrescu/CombatCavallers.Go/lluitador"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// Obtenir la llista de llocs en que es pot picar
	copsPermesos := combat.CopsPermesos{}

	// Crear lluitadors informant-los d'on es pot picar
	lluitador1 := lluitador.CreateLluitadorAleatori("AixafaMandonguilles", copsPermesos)
	lluitador2 := lluitador.CreateLluitadorAleatori("MataCerilles", copsPermesos)

	// Afegir lluitadors al ring i fer-los lluitar
	ring := combat.Ring{}
	ring.EntrenJugadors(lluitador1, lluitador2)
	ring.Lluiteu()

}
