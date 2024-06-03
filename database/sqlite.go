package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Task{})
	return err
}

func createTask(db *gorm.DB, task *Task) error {
	err := db.Create(task).Error
	return err
}

func getTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	err := db.Find(&tasks).Error
	return tasks, err
}

func updateTask(db *gorm.DB, task *Task) error {
	err := db.Save(task).Error
	return err
}

func deleteTask(db *gorm.DB, id uint) error {
	err := db.Delete(&Task{}, id).Error
	return err
}

func TaskManager() {
	db, err := connect()
	if err != nil {
		panic(err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Println(err)
		} else {
			err := sqlDB.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}()

	err = migrate(db)
	if err != nil {
		panic(err)
	}

	// Create tasks
	task1 := &Task{Title: "Buy groceries", Done: false}
	err = createTask(db, task1)
	if err != nil {
		panic(err)
	}

	task2 := &Task{Title: "Finish report", Done: true}
	err = createTask(db, task2)
	if err != nil {
		panic(err)
	}

	// Get all tasks
	tasks, err := getTasks(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("%d - %s (%t)\n", task.ID, task.Title, task.Done)
	}

	// Update a task
	task1.Done = true
	err = updateTask(db, task1)
	if err != nil {
		panic(err)
	}

	// Delete a task
	err = deleteTask(db, task2.ID)
	if err != nil {
		panic(err)
	}
}
