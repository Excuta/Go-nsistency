package service

import (
	"encoding/json"
	"excuta/go-nsistency/repo"
)

func GetCounter() ([]byte, error) {
	counter, error := repo.GetCounter()
	if error != nil {
		return []byte{}, error
	}
	return json.Marshal(Counter{value: counter})
}

func Increment() error {
	return repo.Increment()
}

type Counter struct {
	value int
}
