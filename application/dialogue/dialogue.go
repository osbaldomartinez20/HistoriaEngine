package dialogue

import (
	"encoding/json"
	"io/ioutil"
)

//Dialogue dictates what dialogues store in database neeed to have
type Dialogue struct {
	DialogueID int           `json:"DialogueID"`
	NextDialogueID int       `json:"NextDialogueID"`
	Character string         `json:"Character"`
	Dialogue string          `json:"Dialogue"`
	Choice bool              `json:"Choice"`
}

//Write stores dialogue in a JSON file
func Write(dialogue Dialogue) error {
	filename := "./dialogue/dialogue.json"
	var dialogues []Dialogue

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	} 

	err = json.Unmarshal(file, &dialogues)
	if err != nil {
		return err
	}

	dialogues = append(dialogues, dialogue)

	currDialogue, err := json.Marshal(dialogues)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, currDialogue, 0644)
	if err != nil {
		return err
	}

	return nil
}

//GetAll returns all the dialogues as a Dialogue array
func GetAll() []Dialogue {
	filename := "./dialogue/dialogue.json"
	var dialogues []Dialogue

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	} 

	err = json.Unmarshal(file, &dialogues)
	if err != nil {
		return nil
	}

	return dialogues
}