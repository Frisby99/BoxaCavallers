package net.xaviersala.combatcavallers.combat;

import net.xaviersala.combatcavallers.lluitador.ILluitador;

/**
 * Estructura per anar radiant com va el combat entre dos lluitadors
 */
public interface ICombatents extends IResultat {
    ILluitador getLluitador();

    /**
     * Treu vida a un combatent
     * @param vida Vida que li treu
     * @return Vida que li queda al combatent
     */
    int TreuVida(int vida);
}