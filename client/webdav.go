package client

import (
	"os"
)

type Client interface {
	Upload(name string)
	List() []os.FileInfo
	Delete(name string)
}
