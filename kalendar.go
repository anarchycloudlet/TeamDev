package main

import (
	"bufio"
	"fmt"
	"os"

	"time"
)

type Task struct {
	Description string
	Date        time.Time
}

var tasks []Task
 {

	fmt.Print("Введите дату задачи (в формате YYYY-MM-DD): ")
	scanner.Scan()
	dateInput := scanner.Text()

	date, err := time.Parse("2006-01-02", dateInput)
	if err != nil {
		fmt.Println("Ошибка формата даты. Попробуйте снова.")
		return
	}

	task := Task{
		Description: description,
		Date:        date,
	}

	tasks = append(tasks, task)
	fmt.Println("Задача добавлена!")
}
