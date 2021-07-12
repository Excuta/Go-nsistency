package service

import (
	"encoding/json"
	"excuta/go-nsistency/repo"
	"fmt"
)

func GetCounter() ([]byte, error) {
	counter, error := repo.GetCounter()
	if error != nil {
		return []byte{}, error
	}
	resp, err := json.Marshal(&Counter{counter})
	fmt.Printf("Response json: %v\n", string(resp))
	return resp, err
}

func Increment() error {
	return repo.Increment()
}

type Counter struct {
	Value int64
}
