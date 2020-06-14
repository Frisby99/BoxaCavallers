package lluitador

import (
	"math/rand"
	"strings"

	"github.com/utrescu/CombatCavallers.Go/cops"
)

// DefaultName Nom per defecte del lluitador si no se li dóna nom
const DefaultName = "LluitadorRandom"

// Aleatori es un lluitador que ataca aleatòriament
type Aleatori struct {
	llocs []cops.LlocOnPicar
	nom   string
}

// CreateLluitadorAleatori crea un lluitador que pica de forma aleatòria
func CreateLluitadorAleatori(nomLluitador string, onPicar cops.ILlocOnPicar) Aleatori {
	nomReal := nomLluitador
	if strings.TrimSpace(nomReal) == "" {
		nomReal = DefaultName
	}

	lluitador := Aleatori{nom: nomReal, llocs: onPicar.GetLlocsOnPicar()}
	return lluitador
}

// GetNom retorna el nom del lluitador
func (ll Aleatori) GetNom() string {
	return ll.nom
}

// Protegeix retorna els llocs en que es protegeix
func (ll Aleatori) Protegeix() []cops.LlocOnPicar {
	nous := make([]cops.LlocOnPicar, len(ll.llocs)-1, len(ll.llocs)-1)
	eliminar := ll.Pica()
	for _, element := range ll.llocs {
		if element != eliminar {
			nous = append(nous, element)
		}
	}
	return nous
}

// Pica retorna on pica
func (ll Aleatori) Pica() cops.LlocOnPicar {
	return ll.llocs[rand.Intn(len(ll.llocs))]
}
