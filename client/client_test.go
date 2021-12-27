package client

import (
	"fmt"
	"testing"
)

func TestAli(t *testing.T) {
	//list := ali.List()
	Alic.Upload("123.txt")
	//ali.Sync()
}

func TestWb(t *testing.T) {
	list := WBc.List()
	for _, entry := range list {
		fmt.Println(entry.Name)
	}
	WBc.Delete("123.txt")
	//WBc.Upload("123.txt")

}

//func TestDiff(t *testing.T) {
//
//}
