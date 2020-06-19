package lluitador

import "github.com/utrescu/CombatCavallers.Go/cops"

// ILluitador és una interfície per les operacions que poden fer els lluitadors
type ILluitador interface {
	GetNom() string
	Protegeix() []cops.LlocOnPicar
	Pica() cops.LlocOnPicar
}
