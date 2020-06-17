package combat

import "github.com/utrescu/CombatCavallers.Go/cops"

// CopsPermesos defineix els cops permesos en la batalla
type CopsPermesos struct{}

// GetTotsElsCops diu en quins llocs es pot picar ja sigui permès o no
func (c CopsPermesos) GetTotsElsCops() []cops.LlocOnPicar {
	return []cops.LlocOnPicar{
		cops.Cap,
		cops.CostatEsquerra,
		cops.CostatDret,
		cops.Panxa,
		cops.Collons,
	}
}

// GetLlocsOnPicar diu en quins llocs està permès picar
func (c CopsPermesos) GetLlocsOnPicar() []cops.LlocOnPicar {
	return []cops.LlocOnPicar{
		cops.Cap,
		cops.CostatEsquerra,
		cops.CostatDret,
		cops.Panxa,
	}
}
