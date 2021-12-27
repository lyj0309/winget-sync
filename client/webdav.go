package client

type Client interface {
	Upload(name string)
	List() (files []*File)
	Delete(name string)
}
