package cmd

import (
	"fmt"
	"log"

	"github.com/netapau/dodo/tasks"
	"github.com/spf13/cobra"
)

// endCmd represents the end task command.
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "Cocher la tâche terminée",
	Long: `
	$dodo end <n° de tâche> permet de cocher la tâche avec le n° donné comme étant finie.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		endFlag, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Println("Erreur en recuperant l'id de tâche !")
		}

		db, err := tasks.InitDB()
		if err != nil {
			log.Fatal(err)
		}
		task := tasks.NewTask(db)

		result := make(chan string)
		go task.End(endFlag, result)
		r := <-result
		fmt.Println(r)
	},
}

func init() {
	rootCmd.AddCommand(endCmd)
	endCmd.Flags().IntP("id", "i", 0, "Id de la tâche accomplie.")
}
