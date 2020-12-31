package characters

import (
	"encoding/json"
	"io/ioutil"
)

//Character helps organize the data for the characters
type Character struct {
	CharacterID int          `json:"CharacterID"`
	Name string              `json:"Name"`
	DialogueColor string     `json:"DialogueColor"`
}

//CreateCharacter creates a new character and adds it to the JSON
func CreateCharacter(character Character) error {
	filename := "./characters/characters.json"
	var characters []Character

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	} 

	err = json.Unmarshal(file, &characters)
	if err != nil {
		return err
	}

	characters = append(characters, character)

	currCharacters, err := json.Marshal(characters)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, currCharacters, 0644)
	if err != nil {
		return err
	}

	return nil
}

//GetCharacters returns an array of Character with all of the chracters in the novel
func GetCharacters() []Character {
	filename := "./characters/characters.json"
	var characters []Character

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	} 

	err = json.Unmarshal(file, &characters)
	if err != nil {
		return nil
	}

	return characters
}