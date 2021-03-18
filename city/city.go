package city

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Item struct {
	Name string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {

		return err
	}
	fmt.Println(string(b))

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	return items, err
}
