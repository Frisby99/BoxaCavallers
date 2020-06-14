namespace CombatCavallers.Combat
{
    /// <summary>
    /// Resultat del combat entre dos lluitadors
    /// </summary>
    public interface IResultat
    {
        /// <summary>
        /// Nom del lluitador
        /// </summary>
        string Nom { get; }

        /// <summary>
        /// Vida del lluitador
        /// </summary>
        int Vida { get; }

        /// <summary>
        /// Determina si el Lluitador està derrotat
        /// </summary>
        /// <returns>Retorna si el lluitador està KO o no</returns>
        bool EsKo();
    }
}
