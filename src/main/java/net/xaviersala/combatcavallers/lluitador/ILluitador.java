package net.xaviersala.combatcavallers.lluitador;

import java.util.List;

import net.xaviersala.combatcavallers.LlocOnPicar;

public interface ILluitador {

    /**
     * @return Retorna el nom del Lluitador
     */
    String getNom();

    /**
     * @return Retorna els llocs en que el lluitador es protegeix
     */
    List<LlocOnPicar> Protegeix();

    /**
     * /// @return Retorna el lloc en el que el lluitador pica
     */
    LlocOnPicar Pica();
    
    /**
     * @return Retorna la força del cop
     */
    int ForçaDelCop();

}
