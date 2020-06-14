package combat

import (
	"github.com/utrescu/CombatCavallers.Go/lluitador"
)

// IRing és la interfície que defineix el que es pot fer en un Ring
type IRing interface {
	EntrenJugadors(jugador1 lluitador.ILluitador, jugador2 lluitador.ILluitador)
	Lluiteu() ([]Resultat, error)
}
