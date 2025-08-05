package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// reflect logic without map, all tasks will be write to file instantly if loops is valid.
// can be used slice of structs:
//
//		type Task struct {
//		ID   int    `json:"id"`
//		Text string `json:"text"`
//		}
//	 and add info to file with openfile and "encode" / "decode"
func main() {

	stock := map[int][]string{}
	cmd := []string{}
	count := 0

	// File service
	filename := "tasks.json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("File not found, it will be created automatically")
	}
	// Open the file for reading and writing, create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Ensure the file is closed when done
	defer file.Close()

	fmt.Println("Welcome to task-cli. Run any command. Try 'Help' to view help.")
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok {
		cmd = strings.Fields(scanner.Text())
	} else {
		fmt.Println("Input Error:", scanner.Err())
	}
	if strings.ToLower(cmd[0]) == "help" {
		fmt.Println("Avaliable commands:", "'Help' to view this help", "'add' to add a task", "'update' to update existing task", "'delete' to delete a existing task", "'exit' to close program")
		// add command
	} else if strings.ToLower(cmd[0]) == "add" {
		count++
		cmd[0] = "to do"
		stock[count] = cmd
		fmt.Println("Task added successfully, details:", stock[count])
		// update command
	} else if strings.ToLower(cmd[0]) == "update" {
		id, err := strconv.Atoi(cmd[1])
		if err != nil {
			fmt.Println("Invalid task ID:", cmd[1])
		}
		stock[id] = cmd[2:]
		fmt.Println("Task updated,", "now it is:", stock[id])
		// delete command
	} else if strings.ToLower(cmd[0]) == "delete" {
		id, err := strconv.Atoi(cmd[1])
		if err != nil {
			fmt.Println("Invalid task ID:", cmd[1])
		}
		delete(stock, id)
		fmt.Println("Task deleted:", id)
		// exit command
	} else if strings.ToLower(cmd[0]) == "exit" {
		return
	}
}
