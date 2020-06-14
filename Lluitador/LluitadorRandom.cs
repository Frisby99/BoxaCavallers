using System;
using System.Collections.Generic;
using System.Linq;
using CombatCavallers.cops;

namespace CombatCavallers.Lluitador
{
    /// <summary>
    /// Crea un lluitador que pica i es protegeix de forma
    /// aleatòria
    /// </summary>
    class LluitadorRandom : ILluitador
    {
        private readonly Random rnd;

        public LluitadorRandom(string nom)
        {
            rnd = new Random(Guid.NewGuid().GetHashCode());
            Nom = nom;
        }

        public string Nom { get; }

        public LlocOnPicar Pica()
        {
            return (LlocOnPicar)rnd.Next(Enum.GetNames(typeof(LlocOnPicar)).Length);
        }

        public IEnumerable<LlocOnPicar> Protegeix()
        {
            var punts = Enum.GetValues(typeof(LlocOnPicar)).Cast<LlocOnPicar>();
            var noProtegir = (LlocOnPicar)rnd.Next(Enum.GetNames(typeof(LlocOnPicar)).Length);
            return punts.Where(val => val != noProtegir).ToArray(); ;
        }
    }
}
