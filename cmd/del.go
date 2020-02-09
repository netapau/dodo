package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/netapau/dodo/tasks"
	"github.com/netapau/dodo/tools"
	"github.com/spf13/cobra"
)

// delCmd represents the delelete task command.
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Effacer une tâche de la liste.",
	Long: `
	Pour effacer une tâche de la liste vous devez entrer le n° de identification de tâche.
	Exemple:
	$dodo del -i 2
	effacera la tâche numero 2 de la liste.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetInt("id")

		if id >= 1 {
			fmt.Println("Effacer la tâche avec l'id " + strconv.Itoa(id) + " ? oui/non [o/n]")
			if tools.Valid() == true {

				db, err := tasks.InitDB()
				if err != nil {
					log.Fatal(err)
				}
				task := tasks.NewTask(db)

				result := make(chan string)
				go task.Del(id, result) // Delete task.
				r := <-result
				fmt.Println(r)
			}

		} else {
			fmt.Println("Vous devez passer le paramettre -i (identifiant) pour effacer une tâche\nExemple :\n$todo del -i 12 ")
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	delCmd.Flags().IntP("id", "i", 0, "i est le n° de la tâche a effacer.")
}
