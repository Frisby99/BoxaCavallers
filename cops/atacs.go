package cops

/*
* En Go no hi han enumerats de manera que cal falsificar-los
* amb constants
 */

// LlocOnPicar és un lloc en el que es pot picar
type Atac string

// Llista de llocs on es pot picar
const (
	Normal Atac = "atac normal"
)

// ILlocOnPicar és una interficie per definir on es pot picar
type IAtac interface {
	GetLlocsOnPicar() []LlocOnPicar
}
