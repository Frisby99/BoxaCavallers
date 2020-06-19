package combat

import "github.com/utrescu/CombatCavallers.Go/cops"

// AtacsPermesos defineix els atacs permesos en la batalla
type AtacsPermesos struct{}

// GetTotsElsAtacs diu tots els atacs
func (c AtacsPermesos) GetTotsElsAtacs() []cops.Atac {
	return []cops.Atac{
		cops.Normal,
	}
}

// GetAtacsPermesos diu quins atacs pots fer
func (c AtacsPermesos) GetAtacsPermesos() []cops.Atac {
	return []cops.Atac{
		cops.Normal,
	}
}
