package client

import (
	"fmt"
	"testing"
	"winget-sync/constants"
)

func TestAli(t *testing.T) {
	ali := NewAli()
	//list := ali.List()
	ali.Sync()
}

func TestWb(t *testing.T) {
	wb := NewWebCenter()
	list, err := wb.Client.List(constants.WebDavBasePath)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range list {
		fmt.Println(entry.Name)
	}
}

//func TestDiff(t *testing.T) {
//
//}
