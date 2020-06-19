package net.xaviersala.combatcavallers.lluitador;

import net.xaviersala.combatcavallers.Atac;
import net.xaviersala.combatcavallers.LlocOnPicar;

public class AtacResult {

    private Atac ComAtaca;
    private LlocOnPicar OnAtaca;

    public AtacResult(Atac atac, LlocOnPicar lloc) {
        ComAtaca = atac;
        OnAtaca = lloc;
    }
    
    public Atac getAtac() {
        return ComAtaca;
    }

    public LlocOnPicar getLlocOnPicar() {
        return OnAtaca;
    }
}