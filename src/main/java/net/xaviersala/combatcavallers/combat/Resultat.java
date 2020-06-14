package net.xaviersala.combatcavallers.combat;

import net.xaviersala.combatcavallers.lluitador.ILluitador;

public class Resultat implements ICombatents {

	int Vida;
	ILluitador Lluitador;

	public Resultat(ILluitador lluitador, int vida) {
		Lluitador = lluitador;
		Vida = vida;
	}

	public String getNom() {

		return Lluitador.getNom();
	}

	public int getVida() {

		return Vida;
	}

	public boolean EsKo() {
		return Vida == 0;
	}

	public ILluitador getLluitador() {
		return Lluitador;
	}

	public int TreuVida() {
		Vida--;
		return Vida;
	}

	public String toString() {
		return Lluitador.getNom() + "(" + Vida + ")";
	}

}
