package storage

import (
	"database/sql"

	"github.com/arshakroshandev/go-todo-api/models"
)

var DB *sql.DB

func InitPostgres(connStr string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil
}

func AddTask(task models.Task) error {
	query := `INSERT INTO tasks (name, done) VALUES($1, $2)`
	_, err := DB.Exec(query, task.Name, task.Done)
	return err
}

func GetAllTasks() ([]models.Task, error) {
	rows, err := DB.Query(`SELECT id, name, done FROM tasks ORDER BY id`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Done); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func MarkTaskDone(id int) (models.Task, error) {
	query := `UPDATE tasks SET done = true WHERE id = $1 RETURNING id, name, done`

	var task models.Task
	err := DB.QueryRow(query, id).Scan(&task.ID, &task.Name, &task.Done)
	return task, err
}

func DeleteTask(id int) error {
	_, err := DB.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	return err
}
