package net.xaviersala.combatcavallers.combat;

import net.xaviersala.combatcavallers.lluitador.ILluitador;

public class Resultat implements ICombatents {

	int Vida;
	ILluitador Lluitador;
	boolean Eliminat;

	public Resultat(ILluitador lluitador, int vida) {
		Lluitador = lluitador;
		Vida = vida;
		Eliminat = false;
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

	public int TreuVida(int vida) {
		Vida-=vida;
		return Vida;
	}
	
	@Override
	public boolean EstaEliminat() {
		return Eliminat;
	}

	@Override
	public void Elimina() {
		Eliminat = true;		
	}

	public String toString() {
		return Lluitador.getNom() + "(" + Vida + ")";
	}



}
