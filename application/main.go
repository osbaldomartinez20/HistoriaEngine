package main

import (
	"./save"
)

var saveData save.Data

func loadSave() bool {
	saveFile := save.Read()
	if saveFile.DialogueID == -1 {
		println(string("\033[31m"), "There was an error loading the save data.", string("\033[0m"))
		return false
	}

	saveData = saveFile
	return true
}


func main() {
	println(string("\033[31m"),"hello", string("\033[0m"))
}