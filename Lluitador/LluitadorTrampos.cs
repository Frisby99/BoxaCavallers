using System;
using System.Collections.Generic;
using System.Linq;
using CombatCavallers.cops;

namespace CombatCavallers.Lluitador
{
    /// <summary>
    /// Crea un lluitador que fa servir el cop il·legal
    /// Com a mínim sempre es protegeix el cap
    /// </summary>
    class LluitadorTrampos : ILluitador
    {
        private readonly Random rnd;
        private readonly List<LlocOnPicar> copsPossibles;

        public LluitadorTrampos(string nom)
        {
            rnd = new Random(Guid.NewGuid().GetHashCode());
            copsPossibles = Enum.GetValues(typeof(LlocOnPicar)).Cast<LlocOnPicar>().ToList();
            copsPossibles = new List<LlocOnPicar> {
                LlocOnPicar.Cap,
                LlocOnPicar.CostatDret,
                LlocOnPicar.CostatEsquerra,
                LlocOnPicar.Panxa,
                LlocOnPicar.CopIlegal
             };

            Nom = nom;
        }

        public string Nom { get; }

        public (Atac, LlocOnPicar) Pica()
        {
            int index = rnd.Next(copsPossibles.Count);
            return (Atac.Normal, copsPossibles[index]);
        }

        public IEnumerable<LlocOnPicar> Protegeix()
        {
            var punts = new List<LlocOnPicar>(copsPossibles);
            LlocOnPicar noProtegir;
            do {
                noProtegir = (LlocOnPicar)rnd.Next(copsPossibles.Count);
            }while(noProtegir==LlocOnPicar.Cap);
            
            return punts.Where(val => val != noProtegir).ToArray(); ;
        }
    }
}
