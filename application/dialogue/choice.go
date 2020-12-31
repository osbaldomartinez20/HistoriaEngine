package dialogue

import (
	"encoding/json"
	"io/ioutil"
)

//Choice cotains a choice correspondong to a dialogue id
type Choice struct {
	ChoiceID int         `json:"ChoiceID"`
	DialogueID int       `json:"DialogueID"`
	Choice string        `json:"Choice"`
}

//ChoiceWrite stores the choice in a JSON
func ChoiceWrite(choice Choice) error {
	filename := "./dialogue/choice.json"
	var choices []Choice

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	} 

	err = json.Unmarshal(file, &choices)
	if err != nil {
		return err
	}

	choices = append(choices, choice)

	currChoices, err := json.Marshal(choices)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, currChoices, 0644)
	if err != nil {
		return err
	}

	return nil
}


//GetChoices returns an array of Choice with all the choices in the novel
func GetChoices() []Choice {
	filename := "./dialogue/choice.json"
	var choices []Choice

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	} 

	err = json.Unmarshal(file, &choices)
	if err != nil {
		return nil
	}

	return choices
}