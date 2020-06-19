package combat

import (
	"errors"
	"fmt"
	"log"
	"math/rand"

	"github.com/utrescu/CombatCavallers.Go/cops"
	"github.com/utrescu/CombatCavallers.Go/lluitador"
)

// Ring és on es desenvolupen els combats
type Ring struct {
	resultat    []Resultat
	copsIlegals []int
}

// EntrenJugadors entren els jugadors al ring
func (ring *Ring) EntrenJugadors(jugador1 lluitador.ILluitador, jugador2 lluitador.ILluitador) {
	ring.resultat = make([]Resultat, 2)
	ring.resultat[0] = CreateResultat(jugador1, 20)
	ring.resultat[1] = CreateResultat(jugador2, 20)
	ring.copsIlegals = make([]int, 2)
	ring.copsIlegals[0] = 0
	ring.copsIlegals[1] = 0
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
		_, pica := ring.resultat[elQuePica].GetLluitador().Pica()

		efecteSobreDefensor := 0
		efecteSobreAtacant := 0
		missatge := ""

		switch pica {
		case cops.Cap, cops.CostatDret, cops.CostatEsquerra, cops.Panxa:
			efecteSobreDefensor = 1
			efecteSobreAtacant = 0
			missatge = fmt.Sprintf("%s rep un cop %s de %s", ring.resultat[elQueRep].GetNom(), pica, ring.resultat[elQuePica].GetNom())
			break
		case cops.Collons:
			efecteSobreDefensor = 1
			efecteSobreAtacant = 0
			missatge = fmt.Sprintf("%s rep un cop ILEGAL de %s", ring.resultat[elQueRep].GetNom(), ring.resultat[elQuePica].GetNom())
			break
		default:
			missatge = "No ha passat res"
			break
		}

		haRebut := contains(proteccio, pica) || pica == cops.Collons

		if haRebut {
			ring.resultat[elQueRep].TreuVida(efecteSobreDefensor)
			ring.resultat[elQuePica].TreuVida(efecteSobreAtacant)
			log.Printf(missatge)
		} else {
			log.Printf("%s atura el cop %s de %s", ring.resultat[elQueRep].GetNom(), pica, ring.resultat[elQuePica].GetNom())
		}

		if pica == cops.Collons {
			ring.copsIlegals[elQuePica]++
			if ring.copsIlegals[elQuePica] == 3 {
				ring.resultat[elQuePica].Elimina()
				break
			}
		}

		log.Printf("%s-(%d) vs %s-(%d)", ring.resultat[0].GetNom(), ring.resultat[0].GetVida(),
			ring.resultat[1].GetNom(), ring.resultat[1].GetVida())
		elQuePica = elQueRep
	}

	guanyador := ring.resultat[1]
	perdedor := ring.resultat[0]
	if ring.resultat[1].EsKo() || ring.resultat[1].EstaEliminat() {
		guanyador = ring.resultat[0]
		perdedor = ring.resultat[1]
	}

	comentariLocutor := ""

	if perdedor.EstaEliminat() {
		comentariLocutor = perdedor.GetNom() + " està ELIMINAT per cops il·legals"
	} else {
		log.Printf("%s cau a terra!", perdedor.GetNom())
		if (guanyador.GetVida() - perdedor.GetVida()) > 5 {
			comentariLocutor = "Quina Pallissa!!"
		}
	}

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
