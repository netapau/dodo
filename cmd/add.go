package cmd

import (
	"fmt"
	"log"

	"github.com/netapau/dodo/tasks"
	"github.com/spf13/cobra"
)

// addCmd represents the add command (Cobra)
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajoute une tâche a la liste.",
	Long: `
	Exemple:
	$todo add -t "Aller chercher le pain."
	`,
	Run: func(cmd *cobra.Command, args []string) {

		addFlag, _ := cmd.Flags().GetString("task")
		if addFlag != "" {
			fmt.Println("Ajout de tâche en cours...")
			addTask(addFlag)
		} else {
			fmt.Println("Vous devez utiliser le paramettre -t !!!\nExample :\n$todo add -t 'Aller chercher le pain.'")
		}
	},
}

func addTask(mytask string) {
	db, err := tasks.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	task := tasks.NewTask(db)
	t := tasks.Item{ID: 0, TaskItem: mytask, Finished: 0}
	result := make(chan string)

	go task.Add(t, result)

	output := <-result
	fmt.Println(output)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("task", "t", "", "Ajouter une tâche")
}
