using CombatCavallers.cops;
using CombatCavallers.Lluitador;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;

namespace CombatCavallers.Combat
{
    public class Ring : IRing
    {
        private const int VidaInicialDelsCombatents = 20;

        private readonly Random _rnd;
        private readonly ICombatents[] _Lluitadors;
        private int[] copsIlegals;
        private readonly ILogger<Ring> _logger;

        public Ring(ILogger<Ring> logger)
        {
            _rnd = new Random();
            _Lluitadors = new Resultat[2];
            copsIlegals = new int[2];
            _logger = logger;
        }

        /// <summary>
        /// Col·loca dos lluitadors al ring
        /// </summary>
        /// <param name="lluitador1">Primer lluitador</param>
        /// <param name="lluitador2">Segon lluitador</param>
        public void EntradaLluitadors(ILluitador lluitador1, ILluitador lluitador2)
        {
            _Lluitadors[0] = new Resultat(lluitador1, VidaInicialDelsCombatents);
            _Lluitadors[1] = new Resultat(lluitador2, VidaInicialDelsCombatents);
            copsIlegals = new int[2];
        }

        /// <summary>
        /// Fa que els dos lluitadors lluitin fins que s'acabi el combat
        /// </summary>
        /// <returns>Retorna una llista IResultat amb el resultat del combat</returns>
        public IEnumerable<IResultat> Lluiteu()
        {
            if (_Lluitadors.Any(l => l == null))
            {
                _logger.LogError("No es pot començar el combat. Falten Lluitadors");
                return null;
            }

            var elQuePica = _rnd.Next(2);
            _logger.LogInformation($"Sorteig de qui comença: .... {_Lluitadors[elQuePica].Nom}");

            while (_Lluitadors.All(l => l.EsKo() == false && l.EstaEliminat() == false))
            {
                var elQueRep = (elQuePica + 1) % 2;
                var proteccio = _Lluitadors[elQueRep].Lluitador.Protegeix();
                var (_, pica) = _Lluitadors[elQuePica].Lluitador.Pica();

                var efecteSobreDefensor = 0;
                var efecteSobreAtacant = 0;
                var missatge = "";

                switch (pica)
                {
                    case LlocOnPicar.Cap:
                    case LlocOnPicar.CostatDret:
                    case LlocOnPicar.CostatEsquerra:
                    case LlocOnPicar.Panxa:
                        efecteSobreDefensor = 1;
                        efecteSobreAtacant = 0;
                        missatge = $"{_Lluitadors[elQueRep].Nom} ({_Lluitadors[elQueRep].Vida}) rep un cop al {pica} de {_Lluitadors[elQuePica].Nom} ({_Lluitadors[elQuePica].Vida})";
                        break;
                    case LlocOnPicar.CopIlegal:
                        efecteSobreDefensor = 1;
                        efecteSobreAtacant = 0;
                        missatge = $"{_Lluitadors[elQueRep].Nom} ({_Lluitadors[elQueRep].Vida}) rep un cop ILEGAL de {_Lluitadors[elQuePica].Nom} ({_Lluitadors[elQuePica].Vida})";
                        break;
                    default:
                        missatge = "No ha passat res";
                        break;
                }

                var haRebut = proteccio.Any(l => l == pica) || pica == LlocOnPicar.CopIlegal;
                if (haRebut)
                {
                    _Lluitadors[elQueRep].TreuVida(efecteSobreDefensor);
                    _Lluitadors[elQuePica].TreuVida(efecteSobreAtacant);
                    _logger.LogInformation(missatge);
                }
                else
                {
                    _logger.LogInformation($"{_Lluitadors[elQueRep].Nom} atura el cop al {pica} de {_Lluitadors[elQuePica].Nom}");
                }

                if (pica is LlocOnPicar.CopIlegal)
                {
                    copsIlegals[elQuePica]++;
                    if (copsIlegals[elQuePica] == 3)
                    {
                        _logger.LogInformation($" {_Lluitadors[elQuePica].Nom} ELIMINAT PER COPS ILEGALS");
                        _Lluitadors[elQueRep].Elimina();
                    }
                }

                _logger.LogDebug($"{_Lluitadors[0].Nom}-({_Lluitadors[0].Vida}) vs {_Lluitadors[1].Nom}-({_Lluitadors[1].Vida})");
                elQuePica = elQueRep;
            }

            var guanyador = _Lluitadors.First(l => l.EsKo() == false && l.EstaEliminat() == false);
            var perdedor = _Lluitadors.First(l => l.EsKo() == true || l.EstaEliminat() == true);

            var comentariLocutor = "";

            if (perdedor.EstaEliminat())
            {
                comentariLocutor = $"{perdedor.Nom} està ELIMINAT!";
            }
            else
            {
                comentariLocutor = (guanyador.Vida - perdedor.Vida) > 5 ? "Quina Pallissa!!" : "";
            }

            _logger.LogInformation($"{perdedor.Nom} cau a terra!");
            _logger.LogInformation($"VICTÒRIA DE {guanyador.Nom}!!! {comentariLocutor}");

            return _Lluitadors;

        }
    }
}
