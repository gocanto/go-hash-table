package main

import (
	"hash_table/collection"
	"hash_table/support"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	clear()

	table := collection.Init()
	support.Menu(table)
}

func clear() {
	platform := runtime.GOOS

	if platform == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()

		if err != nil {
			return
		}
	}

	if platform == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()

		if err != nil {
			return
		}
	}
}
