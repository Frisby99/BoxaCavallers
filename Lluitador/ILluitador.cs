using System.Collections.Generic;
using CombatCavallers.cops;

namespace CombatCavallers.Lluitador
{
    public interface ILluitador
    {
        string Nom { get; }
        /// <summary>
        /// Retorna una llista amb els llocs que el combatent protegeix
        /// </summary>
        /// <returns></returns>
        IEnumerable<LlocOnPicar> Protegeix();

        /// <summary>
        /// Retorna el lloc en el que el combatent pica
        /// </summary>
        /// <returns></returns>
        LlocOnPicar Pica();
    }
}