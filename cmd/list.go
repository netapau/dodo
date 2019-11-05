package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fatih/color"
	"github.com/netapau/dodo/tasks"
	"github.com/spf13/cobra"
)

const (
	titleTodoTasks = `
	--------------------------------------
	-     Liste des tâches a faire       -
	--------------------------------------
	`

	titleAllTasks = `
	--------------------------------------
	-     Liste de toutes les tâches     -
	--------------------------------------
	`

	titleEndTasks = `
	--------------------------------------
	-     Liste des tâches finies        -
	--------------------------------------
	`
)

// Pour l'affichage des tâches finies.
// "[ok]" : finie
// "--->" : en cours
func intToStr(finished int) string {
	if finished == 1 {
		return "[ok]"
	}
	return "--->"
}

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste les tâches dans la todo list.",
	Long: `
	Commande:
	$dodo list
	Permet de lister toutes les tâches.

	$dodo list -f
	Permet de lister toutes les tâches finies.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		endFlag, _ := cmd.Flags().GetBool("fin")

		listTasks(endFlag)

		developer, _ := rootCmd.Flags().GetString("developer")
		if developer != "" {
			fmt.Println("[Developer]:", developer)
			fmt.Println("")
		}
	},
}

//listTasks liste les tâches TODO : simplifier !!!
func listTasks(endFlag bool) {

	if endFlag == false {
		c := color.New(color.FgCyan) // color
		c.Println(titleAllTasks)
		c.DisableColor()
	} else {
		fmt.Println(titleEndTasks)
	}

	db, err := tasks.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	task := tasks.NewTask(db)

	todos := task.Get(endFlag)

	lc := color.New(color.FgCyan)
	for _, t := range todos {
		//colors
		if t.Finished == 1 {
			lc = color.New(color.FgCyan)
			lc.Println(" " + strconv.Itoa(t.ID) + "\t" + intToStr(t.Finished) + "\t" + t.TaskItem)

		} else {
			fmt.Println(" " + strconv.Itoa(t.ID) + "\t" + intToStr(t.Finished) + "\t" + t.TaskItem)
		}
		lc.DisableColor()
	}
	fmt.Println(" ")

}

func init() {
	rootCmd.AddCommand(listCmd)
	// $dodo list (liste toutes les tâches)
	// $dodo list -f (liste les tâches finies)
	listCmd.Flags().BoolP("fin", "f", false, "Tâches accomplies !!!")
}
