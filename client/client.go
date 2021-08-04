package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Helloservice interface {
	HelloSay()
	PostSay()
	SetSay()
}

type Hello struct {
	Url  string
	Name string
}

func (h Hello) HelloSay(name, url string) {
	h.Url = url
	h.Name = name
	resp, err := http.Get(h.Url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	data := make([]byte, 1025)
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	print(string(data))
}

func (h Hello) PostOut(name, url string) {
	h.Name = name
	h.Url = url
	resp, err := http.Post(h.Url,
		"application/x-www-form-urlencoded",
		strings.NewReader(h.Name))
	if err != nil {
		print(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Println(string(body))
}

func (h Hello) SetSay(name, url string) (res_name, res_url string) {
	h.Name = name
	h.Url = url
	return h.Name, h.Url
}
