using CombatCavallers.Combat;
using CombatCavallers.Lluitador;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Logging.Console;
using System;

namespace CombatCavallers
{
    class Program
    {
        static void Main(string[] args)
        {
            var serviceCollection = new ServiceCollection();
            ConfigureServices(serviceCollection);
            var serviceProvider = serviceCollection.BuildServiceProvider();

            var ring = serviceProvider.GetService<IRing>();

            // Comença la lluita

            var lluitador1 = new LluitadorRandom("Totxo");
            var lluitador2 = new LluitadorTrampos("Tramposete");

            ring.EntradaLluitadors(lluitador1, lluitador2);
            var resultats = ring.Lluiteu();


            // Tancar logger
            serviceProvider.Dispose();

            // Imprimir resultats
            Console.WriteLine();
            foreach (var resultat in resultats)
            {
                Console.WriteLine($"{resultat.Nom} : {resultat.Vida}");
            }
        }

        private static void ConfigureServices(ServiceCollection serviceCollection)
        {
            serviceCollection.AddLogging(configure => configure.AddConsole(options =>
            {
                options.Format = ConsoleLoggerFormat.Systemd;
            }))
                .Configure<LoggerFilterOptions>(options => options.MinLevel = LogLevel.Information)
                .AddTransient<IRing, Ring>();

        }
    }
}
