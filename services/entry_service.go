package services

import "fmt"

type EntryService struct {
}

func NewEntryService() EntryService {
	return EntryService{}
}

func (s EntryService) GetApi() {
	fmt.Println("Get Api called!")
}
