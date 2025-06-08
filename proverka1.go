package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type note struct {
	id   int
	body string
	date time.Time
}

var (
	notes   = []note{}               // Срез, где хранятся все заметки
	usedIDs = make(map[int]struct{}) // Множество используемых ID для гарантии уникальности
)

// локальный генератор. Эта хуйня хранит ссылку на math/rand
var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func generateID() int {
	var id int
	for {
		id = rnd.Intn(900000) + 100000
		if _, exists := usedIDs[id]; !exists {
			usedIDs[id] = struct{}{}
			break
		}
	}
	return id
}

func addNote() {
	reader := bufio.NewReader(os.Stdin) // эта хуйня нужна, что бы можно было юзать пробелы

	fmt.Print("Введите текст задачи: ")
	input, _ := reader.ReadString('\n') // считывает строку до нажатия Enter
	input = strings.TrimSpace(input)    // убирает лишние пробелы и символ новой строки

	// Запрос даты после ввода текста задачи
	fmt.Print("Введите дату задачи (в формате год-месяц-день): ")
	dateInput, _ := reader.ReadString('\n')
	dateInput = strings.TrimSpace(dateInput)

	date, err := time.Parse("2006-01-02", dateInput)
	if err != nil {
		fmt.Println("Ошибка формата даты. Попробуйте снова.")
		return
	}

	// создает новую заметку с уникальным ID, введенным текстом и датой
	newNote := note{
		id:   generateID(),
		body: input,
		date: date,
	}

	notes = append(notes, newNote)
	fmt.Println("Задача добавлена!")
}

func deleteNote() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID задачи для удаления: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Преобразуем строку в число
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Некорректный формат ID")
		return
	}

	// Вызываем функцию удаления и информируем пользователя
	if deleteNoteByID(id) {
		fmt.Println("Задача удалена")
	} else {
		fmt.Println("Задача с таким ID не найдена")
	}
}

// Вспомогательная функция, которая ищет заметку по ID и удаляет её из среза
func deleteNoteByID(id int) bool {
	for i, n := range notes {
		if n.id == id {
			// Удаляем элемент из среза — объединяем части среза "до" и "после" искомой заметки
			notes = append(notes[:i], notes[i+1:]...)
			// Удаляем ID из множества использованных, чтобы он мог быть сгенерирован повторно
			delete(usedIDs, id)
			return true // Успешно удалили
		}
	}
	return false // Не нашли заметку с таким ID
}

// Функция для вывода на экран всех заметок
func listNotes() {
	if len(notes) == 0 {
		fmt.Println("Нет задач")
		return
	}

	for _, n := range notes {
		fmt.Printf("ID: %d, Задача: %s, Дата: %s\n", n.id, n.body, n.date.Format("2006-01-02"))
	}
}

func main() {
	for {
		fmt.Println("1 - добавить запись \n2 - удалить записи \n3 - показать все записи \n0 - выйти")

		fmt.Print("Введи команду: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 0:
			return
		case 1:
			addNote()
		case 2:
			deleteNote()
		case 3:
			listNotes()
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

