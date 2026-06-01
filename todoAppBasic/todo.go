package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type TaskStatus string

const (
	StatusInProgress TaskStatus = "In Progress"
	StatusCompleted  TaskStatus = "Completed"
)

type Task struct {
	Id                 string
	Title, Description string
	Status             TaskStatus
}

func printMenu() {

	fmt.Println("\t\tTodo App")
	fmt.Println("\tBy: Osazuwa Matthew Ogbebor")
	fmt.Println()

	fmt.Println("\t1. Add a Task")
	fmt.Println("\t2. View All Tasks")
	fmt.Println("\t3. Edit Task Details")
	fmt.Println("\t4. View a Task")
	fmt.Println("\t5. Mark as Done")
	fmt.Println("\t6. Delete a Task")
	fmt.Println("\t7. Exit")
}

func callScanner(scanner *bufio.Scanner, message string) (string, bool) {
	fmt.Println()
	fmt.Print(message)

	if !scanner.Scan() {
		return "", false
	}

	text := strings.TrimSpace(scanner.Text())
	if text == "" {
		return "", false
	}

	return text, true
}

func exit() {
	time.Sleep(time.Second)
	fmt.Print("Exiting.")

	time.Sleep(time.Second)
	fmt.Print(".")

	time.Sleep(time.Second)
	fmt.Print(".")
	time.Sleep(time.Second)
	fmt.Println()

}

func startTodoApp() {

	var todoList []Task
	scanner := bufio.NewScanner(os.Stdin)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
Loop:
	for {
		if ctx.Err() != nil {
			exit()
			break
		}
		printMenu()
		userChoice, ok := callScanner(scanner, "Select your Option: ")
		if !ok {
			fmt.Println()
			continue
		}

		switch userChoice {
		case "1":
			title, ok := callScanner(scanner, "Enter Task Title: ")
			if !ok {
				fmt.Println("Invalid Input")
				fmt.Println()
				continue
			}

			desc, ok := callScanner(scanner, "Enter Task Description: ")
			if !ok {
				fmt.Println("Invalid Input")
				fmt.Println()
				continue
			}

			task := Task{
				Id:          strconv.FormatInt(int64(len(todoList)), 10),
				Title:       title,
				Description: desc,
				Status:      StatusInProgress,
			}

			todoList = append(todoList, task)

			fmt.Println("Task added successfully!")
			fmt.Println()
			time.Sleep(time.Second)
		case "2":
			fmt.Println(todoList)
			fmt.Println()
			time.Sleep(time.Second)
		case "3":
			id, ok := callScanner(scanner, "Enter the task ID: ")
			if !ok {
				fmt.Println("Invalid Id")
				fmt.Println()
				continue
			}

			found := false
			for i, task := range todoList {
				if task.Id == id {
					found = true
					title, ok := callScanner(scanner, "Enter New Title: ")
					if !ok {
						title = task.Title
					}

					desc, ok := callScanner(scanner, "Enter New Description: ")
					if !ok {
						desc = task.Description
					}

					todoList[i] = Task{
						Id:          task.Id,
						Title:       title,
						Description: desc,
						Status:      task.Status,
					}

					fmt.Println("Task Modified Successfully!")
				}
			}

			if !found {
				fmt.Println("Invalid Id")
			}
			fmt.Println()
			time.Sleep(time.Second)

		case "4":
			id, ok := callScanner(scanner, "Enter Task ID: ")
			if !ok {
				fmt.Println("Invalid Id")
				fmt.Println()
				continue
			}

			found := false
			for _, task := range todoList {
				if task.Id == id {
					found = true

					fmt.Printf("Title: %v\n", task.Title)
					time.Sleep(300 * time.Millisecond)

					fmt.Printf("Description: %v\n", task.Description)
					time.Sleep(300 * time.Millisecond)

					fmt.Printf("Status: %v\n", task.Status)
					time.Sleep(300 * time.Millisecond)

				}

				if !found {
					fmt.Println("Invalid Id")
				}

				fmt.Println()
				time.Sleep(time.Second)
			}
		case "5":
			id, ok := callScanner(scanner, "Enter Task ID: ")
			if !ok {
				fmt.Println("Invalid Id")
				fmt.Println()
				continue
			}

			found := false
			for i, task := range todoList {
				if task.Id == id {
					found = true

					todoList[i] = Task{
						Id:          task.Id,
						Title:       task.Title,
						Description: task.Description,
						Status:      StatusCompleted,
					}
					fmt.Println("Task Completed Successfully!")
				}
			}

			if !found {
				fmt.Println("Invalid Id")
			}
			fmt.Println()
			time.Sleep(time.Second)

		case "6":
			id, ok := callScanner(scanner, "Enter Task ID: ")
			if !ok {
				fmt.Println("Invalid Id")
				fmt.Println()
				continue
			}

			found := false
			for i, task := range todoList {
				if task.Id == id {
					found = true
					todoList = append(todoList[:i], todoList[i+1:]...)
				}
			}

			if !found {
				fmt.Println("Invalid Id")
			}

			fmt.Println()
			time.Sleep(time.Second)
		case "7":
			exit()
			break Loop
		default:
			fmt.Println("Invalid Option. Try Again")
			fmt.Println()
		}
	}
}
