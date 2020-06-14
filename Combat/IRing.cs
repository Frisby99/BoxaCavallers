using System.Collections.Generic;
using CombatCavallers.Lluitador;

namespace CombatCavallers.Combat
{
    /// <summary>
    /// Resultat del combat entre dos lluitadors
    /// </summary>
    public interface IRing
    {
        void EntradaLluitadors(ILluitador lluitador1, ILluitador lluitador2);
        IEnumerable<IResultat> Lluiteu();
    }
}