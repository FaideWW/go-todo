package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type todo struct {
	done bool
	task string
}

func (t *todo) print() {
	var statusString string
	if t.done {
		statusString = "[x]"
	} else {
		statusString = "[ ]"
	}

	fmt.Printf("%s %s\n", statusString, t.task)
}

func main() {

	fmt.Printf("Reading file...\n")

	start := time.Now()
	lines, err := readFile("todo.txt")
	if err != nil {
		panic(err)
	}

	var todos []todo
	for _, line := range lines {
		lineSlice := strings.Split(line, ",")

		doneInt, err := strconv.Atoi(lineSlice[0])
		if err != nil {
			panic(err)
		}

		var isDone bool
		if int(doneInt) == 1 {
			isDone = true
		} else {
			isDone = false
		}

		todos = append(todos, todo{done: isDone, task: lineSlice[1]})
	}

	for _, todo := range todos {
		todo.print()
	}

	elapsed := time.Since(start)

	fmt.Printf("Done: read %d lines in %s\n", len(lines), elapsed)

}
