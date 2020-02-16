package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	titleTodo = `	----------------------------------------------------------------
	  Avec dodo vous pouvez faire la gestion des tâches a faire...
	----------------------------------------------------------------
	`
)

var cfgFile string
var heure string
var author string
var license string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dodo",
	Short: "dodo permet de gérer une liste de tâches a éffectuer.",
	Long: `
` + titleTodo + `
Examples:
   Ajouter une tache :
   $dodo add -t "Donner a manger au chat !"

   Vôtre tâche effectuée, vous pouvez la marquer a l'aide de son n° de tâche :
   $dodo end -i 2

   Supprimer la tâche avec l'identifiant '2' de vôtre liste :
   $dodo del -i 2

   Lister les tâches :
   $dodo list    (toutes les tâches)
   $dodo list -f (tâches finies)
  `,
	// Uncomment the following line if your bare application has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" ")
		fmt.Println("dodo - Application de gestion des tâches")
		fmt.Println("Voir l'aide: dodo -h ")
		fmt.Println(" ")

		time, _ := cmd.Flags().GetString("time")
		if time != "" {
			fmt.Println("[time]   ", time)
			fmt.Println("[author] ", author)
			fmt.Println("[license]", license)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.Flags().StringVar(&heure, "time", time.Now().Format("15:04:05"), "Time.")
	rootCmd.Flags().StringVar(&cfgFile,
		"config",
		"./dodo.yaml",
		"config file (default is $HOME/.github.com/netapau/dodo.yaml)")

	// Cobra also supports local flags, ex:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Définir le chemin d'accès, y compris le nom et l'extension du fichier de configuration.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Défini (le ou les chemins) pour les fichiers de configuration: ".github.com\netapau\dodo" (sans extension).
		viper.AddConfigPath(home)
		//Nom du fichier de configuration (sans extension)
		viper.SetConfigName("dodo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	author = isSet("author")
	license = isSet("license")
}

func isSet(varName string) string {
	if !viper.IsSet(varName) {
		log.Fatal("missing" + varName)
	}
	return viper.GetString(varName)
}
