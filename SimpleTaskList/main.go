package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define the Task struct
type Task struct {
	Name      string
	Completed bool
}

func addTask(tasks []Task, taskName string, taskStatus bool) []Task {
	newTask := Task{
		Name:      taskName,
		Completed: taskStatus,
	}
	return append(tasks, newTask)
}

func removeTask(tasks []Task, index int) []Task {
	return append(tasks[:index], tasks[index+1:]...)
}

func completeTask(tasks []Task, index int) {
	tasks[index].Completed = true
}

func viewAllTasks(tasks []Task) {
	for _, task := range tasks {
		if task.Completed {
			fmt.Printf("[x] %s\n", task.Name)
		} else {
			fmt.Printf("[ ] %s\n", task.Name)
		}
	}
}

func main() {
	// Read input
	var numTasksStr string
	var taskDataStr string
	var toRemoveTaskIndexStr string

	fmt.Scan(&numTasksStr)
	numTasks, _ := strconv.Atoi(numTasksStr)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskDataStr = scanner.Text()

	scanner.Scan()
	toRemoveTaskIndexStr = scanner.Text()
	toRemoveTaskIndex, _ := strconv.Atoi(toRemoveTaskIndexStr)

	// Filling task list
	tasks := make([]Task, 0, numTasks)
	taskDataList := strings.Split(taskDataStr, ",")
	for _, taskData := range taskDataList {
		taskName := strings.Split(taskData, ":")[0]
		taskStatusStr := strings.Split(taskData, ":")[1]
		taskStatus, _ := strconv.ParseBool(taskStatusStr)
		tasks = addTask(tasks, taskName, taskStatus)
	}

	// Remove a task
	removedTask := tasks[toRemoveTaskIndex]
	tasks = removeTask(tasks, toRemoveTaskIndex)
	var remainingTasks int
	var completedTasks int
	for _, task := range tasks {
		if task.Completed {
			completedTasks++
		} else {
			remainingTasks++
		}
	}

	// Summary
	viewAllTasks(tasks)
	fmt.Printf("Task '%s' removed successfully!\n", removedTask.Name)
	fmt.Printf("Total: %d tasks (%d completed, %d remaining)", len(tasks), completedTasks, remainingTasks)
}
