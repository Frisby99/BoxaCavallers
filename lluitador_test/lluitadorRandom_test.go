package lluitador_test

import (
	"testing"

	"github.com/utrescu/CombatCavallers.Go/cops"
	"github.com/utrescu/CombatCavallers.Go/lluitador"
)

type MockCops struct{}

func (c MockCops) GetLlocsOnPicar() []cops.LlocOnPicar {
	return []cops.LlocOnPicar{
		cops.Cap,
	}
}

func TestGetNomRetornaElQueEsPosaAlCreate(t *testing.T) {
	expected := "X11"
	guerrer := lluitador.CreateLluitadorAleatori(expected, MockCops{})

	result := guerrer.GetNom()

	if result != expected {
		t.Errorf("El nom no coincideix, rebut: %s esperava: %s", result, expected)
	}
}

var nomsBuitsTest = []struct {
	nom string
}{
	{""},
	{" "},
}

func TestGetNomSiEsBuitElNomEsDefaultName(t *testing.T) {
	expected := lluitador.DefaultName

	for _, data := range nomsBuitsTest {
		guerrer := lluitador.CreateLluitadorAleatori(data.nom, MockCops{})
		result := guerrer.GetNom()
		if result != expected {
			t.Errorf("El nom no coincideix, rebut: '%s' esperava: '%s'", result, expected)
		}
	}
}

func TestPicaRetornaElValorAletoriDeUn(t *testing.T) {
	expected := cops.Cap
	guerrer := lluitador.CreateLluitadorAleatori("Nom", MockCops{})

	result := guerrer.Pica()

	if result != expected {
		t.Errorf("El cop ha de ser a %s i és a %s", expected, result)
	}
}

func TestProtegeixRetornaLlistaBuidaSiNomesTeUnValor(t *testing.T) {
	expected := []cops.LlocOnPicar{}

	guerrer := lluitador.CreateLluitadorAleatori("Nom", MockCops{})

	result := guerrer.Protegeix()

	if len(result) != len(expected) {
		t.Errorf("El cop ha de ser a %s i és a %s", expected, result)
	}
}
