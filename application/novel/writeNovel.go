package novel

import (
	"../dialogue"
	"../characters"
	"bufio"
	"strconv"
	"strings"
)

//NewChoice adds a new choice
func NewChoice(cid int, did int, nid int, choice string) {
	err := dialogue.WriteChoice(dialogue.Choice{ChoiceID: cid, DialogueID: did, NextDialogueID: nid, Choice: choice})
	if err != nil {
		println("Error writing choice.")
	} else {
		println("Successfu in writing choice.")
	}
}

//NewDialogue adds a new dialogue
func NewDialogue(did int, nid int, character string, dia string, choice bool) {
	err := dialogue.Write(dialogue.Dialogue{DialogueID: did, NextDialogueID: nid, Character: character, Dialogue: dia, Choice: choice})
	if err != nil {
		println("Error writing dialogue.")
	} else {
		println("Successfu in writing dialogue.")
	}
}

//NewCharacter creates a new character
func NewCharacter(chid int, name string, color string) {
	err := characters.CreateCharacter(characters.Character{CharacterID: chid, Name: name, DialogueColor: color})
	if err != nil {
		println("Error creating character.")
	} else {
		println("Successfu in creating character.")
	}
}

//AskForNewCharacter is
func AskForNewCharacter(reader *bufio.Reader) {
	print("Character ID(int): ")
	id, _ := reader.ReadString('\n')
	id = strings.Replace(id, "\n", "", -1)
	tid, _ := strconv.ParseInt(id, 10, 64)
	cid := int(tid)

	print("Character name: ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)

	print("Dialogue color: ")
	color, _ := reader.ReadString('\n')
	color = strings.Replace(color, "\n", "", -1)

	NewCharacter(cid, name, color)
}

//AskForNewDialogue is
func AskForNewDialogue(reader *bufio.Reader) {
	print("Dialogue ID(int): ")
	diaid, _ := reader.ReadString('\n')
	diaid = strings.Replace(diaid, "\n", "", -1)
	tid, _ := strconv.ParseInt(diaid, 10, 64)
	did := int(tid)

	print("Next Dialogue ID(int): ")
	nexid, _ := reader.ReadString('\n')
	nexid = strings.Replace(nexid, "\n", "", -1)
	tid, _ = strconv.ParseInt(nexid, 10, 64)
	nid := int(tid)

	print("Character: ")
	character, _ := reader.ReadString('\n')
	character = strings.Replace(character, "\n", "", -1)

	print("dialogue: ")
	dlg, _ := reader.ReadString('\n')
	dlg = strings.Replace(dlg, "\n", "", -1)

	print("Choice(Y/N): ")
	choi, _ := reader.ReadString('\n')
	choi = strings.Replace(choi, "\n", "", -1)
	choice := strings.Compare("Y", choi) == 0

	NewDialogue(did, nid, character, dlg, choice)

}

//AskForNewChoice is
func AskForNewChoice(reader *bufio.Reader) {
	print("Dialogue ID(int): ")
	choid, _ := reader.ReadString('\n')
	choid = strings.Replace(choid, "\n", "", -1)
	tid, _ := strconv.ParseInt(choid, 10, 64)
	cid := int(tid)


	print("Dialogue ID(int): ")
	diaid, _ := reader.ReadString('\n')
	diaid = strings.Replace(diaid, "\n", "", -1)
	tid, _ = strconv.ParseInt(diaid, 10, 64)
	did := int(tid)

	print("Next Dialogue ID(int): ")
	nexid, _ := reader.ReadString('\n')
	nexid = strings.Replace(nexid, "\n", "", -1)
	tid, _ = strconv.ParseInt(nexid, 10, 64)
	nid := int(tid)

	print("Choice Text: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.Replace(choice, "\n", "", -1)

	NewChoice(cid, did, nid, choice)
}
