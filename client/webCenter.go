package client

import (
	"github.com/jlaffaye/ftp"
	"log"
	"time"
)

type WebCenter struct {
	Client *ftp.ServerConn
}

func ff() {

	// Do something with the FTP conn

	//if err := c.Quit(); err != nil {
	//	log.Fatal(err)
	//}
}
func NewWebCenter() *WebCenter {
	c, err := ftp.Dial("210.30.62.70:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("lyj", "123456")
	if err != nil {
		log.Fatal(err)
	}

	return &WebCenter{Client: c}
}

func (w *WebCenter) name() {

}
