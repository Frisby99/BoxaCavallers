package net.xaviersala.combatcavallers.combat;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;
import java.util.logging.Level;
import java.util.logging.Logger;

import net.xaviersala.combatcavallers.LlocOnPicar;
import net.xaviersala.combatcavallers.lluitador.AtacResult;
import net.xaviersala.combatcavallers.lluitador.ILluitador;

/**
 * Lloc en el que es desenvolupa la lluita
 *
 * Per defecte tots els lluitadors tenen 20 de vida
 */
public class Ring implements IRing {

    private static final int PALLISSA = 5;
    private static final int VIDAINICIAL = 20;

    Random aleatori = new Random();
    private final static Logger LOGGER = Logger.getLogger("Ring");

    List<ICombatents> _Lluitadors;
    List<Integer> copsIlegals;

    public Ring() {
        _Lluitadors = new ArrayList<ICombatents>(2);
        copsIlegals = new ArrayList<Integer>(2);
    }

    /**
     * Col·loca dos lluitadors al ring
     *
     * @param name lluitador1 Primer lluitador
     * @param name lluitador2 Segon lluitador
     */
    public void EntradaLluitadors(ILluitador lluitador1, ILluitador lluitador2) {

        _Lluitadors.clear();
        _Lluitadors.add(new Resultat(lluitador1, VIDAINICIAL));
        copsIlegals.add(0);
        _Lluitadors.add(new Resultat(lluitador2, VIDAINICIAL));
        copsIlegals.add(0);
    }

    /**
     * Fa que els dos lluitadors lluitin fins que s'acabi el combat
     *
     * @return Una llista IResultat amb el resultat del combat
     */
    public List<IResultat> Lluiteu() {
        if (_Lluitadors.size() != 2) {
            LOGGER.log(Level.SEVERE, "No es pot començar el combat. Falten Lluitadors");
            return null;
        }

        int elQuePica = aleatori.nextInt(2);
        LOGGER.log(Level.INFO, "Sorteig de qui comença: ....  " + _Lluitadors.get(elQuePica).getNom());

        while (!_Lluitadors.get(0).EsKo() && !_Lluitadors.get(0).EstaEliminat() && !_Lluitadors.get(1).EsKo()
                && !_Lluitadors.get(1).EstaEliminat()) {

            int elQueRep = (elQuePica + 1) % 2;
            List<LlocOnPicar> proteccio = _Lluitadors.get(elQueRep).getLluitador().Protegeix();
            AtacResult atac = _Lluitadors.get(elQuePica).getLluitador().Pica();
            LlocOnPicar pica = atac.getLlocOnPicar();

            int efecteSobreDefensor = 0;
            int efecteSobreAtacant = 0;
            String missatge = "";

            switch (pica) {
                case Cap:
                case CostatDret:
                case CostatEsquerra:
                case Panxa:
                    efecteSobreDefensor = 1;
                    efecteSobreAtacant = 0;
                    missatge = _Lluitadors.get(elQueRep).getNom() + "(" + _Lluitadors.get(elQueRep).getVida()
                            + ") rep un cop al " + pica + " de " + _Lluitadors.get(elQuePica).getNom() + "("
                            + _Lluitadors.get(elQueRep).getVida() + ")";
                    break;
                case CopIlegal:
                    efecteSobreDefensor = 1;
                    efecteSobreAtacant = 0;
                    missatge = _Lluitadors.get(elQueRep).getNom() + "(" + _Lluitadors.get(elQueRep).getVida()
                            + ") rep un cop ILEGAL de " + _Lluitadors.get(elQuePica).getNom() + "("
                            + _Lluitadors.get(elQueRep).getVida() + ")";
                    ;
                    break;
                default:
                    missatge = "No ha passat res";
                    break;
            }

            boolean haRebut = proteccio.contains(pica) || pica == LlocOnPicar.CopIlegal;
            if (haRebut) {
                _Lluitadors.get(elQueRep).TreuVida(efecteSobreDefensor);
                _Lluitadors.get(elQuePica).TreuVida(efecteSobreAtacant);
                LOGGER.log(Level.INFO, missatge);

            } else {
                LOGGER.log(Level.INFO, _Lluitadors.get(elQueRep).getNom() + " atura el cop al " + pica + " de "
                        + _Lluitadors.get(elQuePica).getNom());
            }

            if (pica == LlocOnPicar.CopIlegal) {
                copsIlegals.set(elQuePica, copsIlegals.get(elQuePica) + 1);
                if (copsIlegals.get(elQuePica) == 3) {
                    _Lluitadors.get(elQuePica).Elimina();
                }
            }

            elQuePica = elQueRep;
        }

        int guanya = _Lluitadors.get(0).EsKo() || _Lluitadors.get(0).EstaEliminat() ? 1 : 0;
        int perd = (guanya + 1) % 2;

        ILluitador guanyador = _Lluitadors.get(guanya).getLluitador();
        ILluitador perdedor = _Lluitadors.get(perd).getLluitador();

        String comentariLocutor = "";

        if (_Lluitadors.get(perd).EstaEliminat()) {
            comentariLocutor = perdedor.getNom() + " està ELIMINAT per cops il·legals";
        } else {
            LOGGER.log(Level.INFO, perdedor.getNom() + " cau a terra!");
            if (_Lluitadors.get(guanya).getVida() - _Lluitadors.get(perd).getVida() > PALLISSA) {
                comentariLocutor = "Quina Pallissa!!";
            }
        }

        LOGGER.log(Level.INFO, "VICTÒRIA DE " + guanyador.getNom() + "!!! " + comentariLocutor);

        return toIResultat();

    }

    private List<IResultat> toIResultat() {
        List<IResultat> resultat = new ArrayList<IResultat>();
        resultat.add(_Lluitadors.get(0));
        resultat.add(_Lluitadors.get(1));
        return resultat;
    }
}
