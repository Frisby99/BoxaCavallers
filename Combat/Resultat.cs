using CombatCavallers.Lluitador;

namespace CombatCavallers.Combat
{
    /*
     * Aquesta classe la fa servir el ring per anar portant els comptes
     * d'en quin estat es troba el combat.
     */
    class Resultat : ICombatents
    {
        public ILluitador Lluitador { get; }
        public string Nom => Lluitador.Nom;
        public int Vida { get; private set; }

        public bool Eliminat { get; set; }

        public Resultat(ILluitador lluitador, int vida)
        {
            Lluitador = lluitador;
            Vida = vida;
            Eliminat = false;
        }

        public int TreuVida(int forca)
        {
            Vida -= forca;
            return Vida;
        }

        public bool EsKo()
        {
            return Vida == 0;
        }

        public bool EstaEliminat() => Eliminat;

        public void Elimina()
        {
            Eliminat = true;
        }
    }
}
