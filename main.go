package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	rp "github.com/yogavredizon/todolist/pkg/repository"
)

func main() {

	run := 1
	var todo rp.Todo
	scanner := bufio.NewScanner(os.Stdin)

	for run > 0 {
		fmt.Print("\nTodolist\n")
		todo.ShowTodo()

		fmt.Println("\nMenu Todolist")
		fmt.Print(" 1. Tambah Todo \n 2. Update Todo \n 3. Update Time \n 4. Hapus Todo \n 5. Keluar \n")

		// Get input from terminal dan convert into integer
		fmt.Print("Pilih Menu : ")
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())

		switch num {
		case 1:
			fmt.Println("Tambah Todo : ")

			fmt.Print("Todo : ")
			scanner.Scan()
			scanner.Text()
			task := scanner.Text()

			fmt.Print("Waktu : ")
			scanner.Scan()
			scanner.Text()
			time := scanner.Text()

			_, err := todo.AddTodo(task, time)
			if err != nil {
				fmt.Printf("\nError : %v \n", err)
			}
		case 2:
			fmt.Println("Update Todo dari Id")
			fmt.Print("Masukan ID : ")
			scanner.Scan()
			scanner.Text()
			id, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Masukan Todo : ")
			scanner.Scan()
			scanner.Text()
			task := scanner.Text()
			res, err := todo.UpdateTodo(id, task)

			if err != nil {
				fmt.Printf("\nError : %v \n", err)
			} else {
				fmt.Printf("\n %v Updated \n", res)
			}

		case 3:
			fmt.Println("Update Waktu dari Id")
			fmt.Print("Masukan ID : ")
			scanner.Scan()
			scanner.Text()
			id, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Masukan Waktu : ")
			scanner.Scan()
			scanner.Text()
			task := scanner.Text()
			res, err := todo.UpdateTime(id, task)

			if err != nil {
				fmt.Printf("\nError : %v \n", err)
			} else {
				fmt.Printf("\n Time in  ID %v updated \n", res)
			}
		case 4:
			scanner.Scan()
			scanner.Text()
			id, _ := strconv.Atoi(scanner.Text())

			res, err := todo.RemoveTodo(id)

			if err != nil {
				fmt.Printf("\nError %v : \n", err)
			} else {
				fmt.Printf("\n Removed ID %v \n", res)
			}
		case 5:
			run = 0
		default:
			fmt.Println("Pilihan tidak ada")
		}
	}

}
