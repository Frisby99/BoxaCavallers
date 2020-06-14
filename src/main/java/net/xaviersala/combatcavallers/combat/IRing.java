package net.xaviersala.combatcavallers.combat;

import java.util.List;

import net.xaviersala.combatcavallers.lluitador.ILluitador;

public interface IRing {
	void EntradaLluitadors(ILluitador lluitador1, ILluitador lluitador2);
	List<IResultat> Lluiteu();
}
