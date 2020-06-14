package net.xaviersala.combatcavallers.combat;

/**
 * Resultat del combat entre dos lluitadors
 */
public interface IResultat {
    /**
     * Nom del lluitador
     *
     * @return nom del lluitador
     */
    String getNom();

    /**
     * Vida del lluitador
     *
     * @return vida que té el Lluitador
     */
    int getVida();

    /**
     * Determina si el Lluitador està derrotat
     *
     * @return Si està derrotat
     */
    boolean EsKo();
}