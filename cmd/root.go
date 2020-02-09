package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	titleTodo = `
	----------------------------------------------------------------
	  Avec dodo vous pouvez faire la gestion des tâches a faire...
	----------------------------------------------------------------

	`
)

var cfgFile string
var heure string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dodo",
	Short: "dodo permet de gérer une liste de tâches a éffectuer.",
	Long: `
` + titleTodo + `

  Par example:
  $dodo add -t "Aujourd'hui : Faire quelque chose de bon pour l'humanité !"

  Une fois vôtre tâche efectuée vous pouvez la cocher a l'aide de son n° de tâche:
  $dodo end -i 2

  Vous pouvez également supprimer des tâches de vôtre liste:
  $dodo del -i 2

	Pour lister les tâches :
	$dodo list
	$dodo list -f (tâches finies)
  `,
	// Uncomment the following line if your bare application has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" ")
		fmt.Println("dodo, gestion de tâches !")
		fmt.Println("Voir l'aide: dodo -h ")
		fmt.Println(" ")
		time, _ := cmd.Flags().GetString("time")
		if time != "" {
			fmt.Println("[time]", time)
			fmt.Println("")
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

	rootCmd.PersistentFlags().StringVar(&heure, "time", time.Now().Format("15:04:05"), "Time.")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./dodo.yaml", "config file (default is $HOME/.github.com/netapau/dodo.yaml)")
	//rootCmd.PersistentFlags().StringVar(&task, "task", "Rien a faire", "Tâche a réaliser.")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".github.com\netapau\dodo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("dodo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
