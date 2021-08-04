package main

import "strconv"

type Services interface {
	GetService() (string, string)
	PostService(name, leixi string)
}

type ServicesType struct {
	Name        string
	ServiceType string
}

type Binary uint64

func (i Binary) String() string {
	return strconv.Itoa(int(i))
}

// func NewServicesManger(s *ServicesType) *Services {
// 	return &Services{
// 		s.GetService(),
// 		s.PostService(),
// 	}
// }

func (t ServicesType) GetService() (string, string) {
	return t.Name, t.ServiceType
}

func (t ServicesType) PostService(name, leixi string) {
	t.Name = name
	t.ServiceType = leixi
	println(t.Name, t.ServiceType)
}

func main() {
	var bn Services
	bn = ServicesType{}
	bn.PostService("Zhang san", "VM")

}
