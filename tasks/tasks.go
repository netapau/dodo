package tasks

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

//Task est un type qui contient la database.
type Task struct {
	DB     *sql.DB
	dbName string
}

// //Exist return true if task exists.
// func (task *Task) Exists(idTask int) bool {

// 	// , result chan bool
// 	var verif bool = false
// 	fmt.Println("Entrée verif = :", verif)

// 	stmt, _ := task.DB.Prepare(`
// 		SELECT * FROM tasks WHERE id=(?);
// 	`)

// 	err := stmt.QueryRow(stmt, idTask).Scan(&idTask)

// 	if err != nil {
// 		if err != sql.ErrNoRows {
// 			//TODO://log.Print(err)
// 			//verif = false
// 		}
// 		//result <- false
// 		//verif = false
// 	} else {
// 		verif = true
// 	}
// 	//task.DB.Close()

// 	//result <- true

// 	fmt.Println("Sortie verif = :", verif)
// 	return (verif)
// }

//Exist check if task exists.
func (task *Task) Exists(idTask int, result chan bool) {

	stmt, _ := task.DB.Prepare(`
		SELECT id FROM tasks WHERE id=(?);
	`)

	err := stmt.QueryRow(idTask).Scan(&idTask)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
		}
		fmt.Println("On ne peut effacer la tache avec l'id", idTask, "car elle n'existe pas.")
		result <- false
	}
	result <- true
}

//End marque une tâche terminée avec une ID passée en paramèttre.
func (task *Task) End(idTask int, result chan string) {

	stmt, _ := task.DB.Prepare(`
		UPDATE tasks SET finished='1' WHERE id=(?);
	`)
	_, err := stmt.Exec(idTask)
	if err != nil {
		result <- "Erreur ! on a pas pu marquer cette tâche comme executée !!!"
	}
	task.DB.Close()
	result <- "La tâche avec l'id n° " + strconv.Itoa(idTask) + " a été marquée comme étant terminée"
}

//Del efface une tâche avec une ID passée en paramèttre.
func (task *Task) Del(idTask int, result chan string) {

	stmt, _ := task.DB.Prepare(`
		DELETE FROM tasks WHERE id = (?);
		`)
	stmt.Exec(idTask)
	task.DB.Close()
	result <- "La tâche avec l'id n° " + strconv.Itoa(idTask) + " a été effacée"

}

//Add ajoute une nouvelle tâche.
func (task *Task) Add(item Item, result chan string) {

	statement, _ := task.DB.Prepare(`
		INSERT INTO tasks (task,finished) VALUES (?,?)
		`)
	_, err := statement.Exec(item.TaskItem, item.Finished)
	if err != nil {
		// TODO: faire un gestion des erreurs plus fine.
		result <- "Erreur, pas d'ajout de la tâche : [ " + item.TaskItem + " ] !!!"
	}
	task.DB.Close()

	result <- "Tâche : [ " + item.TaskItem + " ] ajouteé."

}

// Get recupère un slice de taches.
func (task *Task) Get(end bool) []Item {
	// READ !
	var rows *sql.Rows

	if end == true {
		rows, _ = task.DB.Query("SELECT id, task, finished FROM tasks WHERE finished = 1")
	} else {
		rows, _ = task.DB.Query("SELECT id, task, finished FROM tasks")
	}
	var items = []Item{}
	var id int
	var t string
	var finished int

	for rows.Next() {
		rows.Scan(&id, &t, &finished)

		items = append(items, Item{id, t, finished})
	}
	task.DB.Close()
	return items
}

// GetTask recupère une tâche par son id. TODO /: A VERIFIER !
func (task *Task) GetTask(idTask int) Item {

	var rows *sql.Rows
	rows, _ = task.DB.Query("SELECT id, task, finished FROM tasks WHERE id = " + strconv.Itoa(idTask) + ";")

	var item = Item{}
	var id int
	var t string
	var finished int

	for rows.Next() {
		rows.Scan(&id, &t, &finished)
		item = Item{id, t, finished}
	}
	task.DB.Close()
	return item
}

//NewTask crée la table et initialise la database.
func NewTask(db *sql.DB) *Task {
	statement, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "tasks" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"task" TEXT,
		"finished" INTEGER);
		`)
	statement.Exec()

	return &Task{
		DB: db,
	}
}
