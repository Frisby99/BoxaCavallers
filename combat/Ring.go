package combat

import (
	"errors"
	"log"
	"math/rand"

	"github.com/utrescu/CombatCavallers.Go/lluitador"
	"github.com/utrescu/CombatCavallers.Go/cops"
)

// Ring és on es desenvolupen els combats
type Ring struct {
	resultat []Resultat
}

// EntrenJugadors entren els jugadors al ring
func (ring *Ring) EntrenJugadors(jugador1 lluitador.ILluitador, jugador2 lluitador.ILluitador) {
	ring.resultat = make([]Resultat, 2)
	ring.resultat[0] = CreateResultat(jugador1, 20)
	ring.resultat[1] = CreateResultat(jugador2, 20)
}

// Lluiteu és la ordre de començar el combat
func (ring *Ring) Lluiteu() ([]Resultat, error) {
	if ring.resultat == nil {
		return nil, errors.New("ERROR Falten lluitadors")
	}

	elQuePica := rand.Intn(2)
	log.Printf("Sorteig de qui comença: .... %s", ring.resultat[elQuePica].GetNom())

	for ring.resultat[0].EsKo() == false && ring.resultat[1].EsKo() == false {
		elQueRep := (elQuePica + 1) % 2

		proteccio := ring.resultat[elQueRep].GetLluitador().Protegeix()
		pica := ring.resultat[elQuePica].GetLluitador().Pica()

		haRebut := contains(proteccio, pica)

		if haRebut {
			ring.resultat[elQueRep].TreuVida()
			log.Printf("%s rep un cop %s de %s", ring.resultat[elQueRep].GetNom(), pica, ring.resultat[elQuePica].GetNom())
		} else {
			log.Printf("%s atura el cop %s de %s", ring.resultat[elQueRep].GetNom(), pica, ring.resultat[elQuePica].GetNom())
		}

		log.Printf("%s-(%d) vs %s-(%d)", ring.resultat[0].GetNom(), ring.resultat[0].GetVida(),
			ring.resultat[1].GetNom(), ring.resultat[1].GetVida())
		elQuePica = elQueRep
	}

	guanyador := ring.resultat[1]
	perdedor := ring.resultat[0]
	if ring.resultat[1].EsKo() {
		guanyador = ring.resultat[0]
		perdedor = ring.resultat[1]
	}

	comentariLocutor := ""
	if (guanyador.GetVida() - perdedor.GetVida()) > 5 {
		comentariLocutor = "Quina Pallissa!!"
	}

	log.Printf("%s cau a terra!", perdedor.GetNom())
	log.Printf("VICTÒRIA DE %s!!! %s", guanyador.GetNom(), comentariLocutor)

	return ring.resultat, nil
}

func contains(s []cops.LlocOnPicar, e cops.LlocOnPicar) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
