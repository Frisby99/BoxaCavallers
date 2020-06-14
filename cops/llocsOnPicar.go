package cops

/*
* En Go no hi han enumerats de manera que cal falsificar-los
* amb constants
 */

// LlocOnPicar és un lloc en el que es pot picar
type LlocOnPicar string

// Llista de llocs on es pot picar
const (
	Cap            LlocOnPicar = "al cap"
	CostatEsquerra LlocOnPicar = "al costat esquerra"
	CostatDret     LlocOnPicar = "al costat dret"
	Panxa          LlocOnPicar = "a la panxa"
	Collons        LlocOnPicar = "als collons"
)

// ILlocOnPicar és una interficie per definir on es pot picar
type ILlocOnPicar interface {
	GetLlocsOnPicar() []LlocOnPicar
}
