package save


import (
	"encoding/json"
	"io/ioutil"
)

//Data keeps track of what is saved by the novel
type Data struct {
	DialogueID int                 `json:"DialogueID"`
	FinishedNovelOnce bool          `json:"FinishedNovelOnce"`
}

//Write writes the current state of the novel into a JSON
func Write(data Data) error {
	filename := "./save/savefile.json"
	save, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, save, 0644)
	if err != nil {
		return err
	}

	return nil
}

//Read reads the data found in savefile JSON
func Read() Data {
	filename := "./save/savefile.json"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Data{DialogueID: -1, FinishedNovelOnce: false}
	}

	var save Data
	err = json.Unmarshal(file, &save)
	if err != nil {
		return Data{DialogueID: -1, FinishedNovelOnce: false}
	}

	return save
}