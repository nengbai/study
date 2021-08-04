package main

import (
	"fmt"
	"io/ioutil"
	"study/client"
	"study/common"
	"study/config"
	"study/prometheus"

	"gopkg.in/yaml.v2"
)

type Server struct {
	ServerName string
	ServerIP   string
}
type Serverslice struct {
	Servers []Server
}

func main() {
	h := client.Hello{
		Url:  "",
		Name: "",
	}
	h.Name, h.Url = h.SetSay("Bai Neng", "http://127.0.0.1:8000")
	h.HelloSay(h.Name, h.Url)
	h.Name, h.Url = h.SetSay("Jing Dong", "http://127.0.0.1:8000/post")
	h.PostOut(h.Name, h.Url)
	// j := `["Apple", "Orange", "Banana"]`
	// var a = common.Users{}
	// a.Username = "Bai Hong yi"
	// a.Password = "长沙"
	// a.Age = 10
	j := `{"username":"Bai Hong yi","password":"长沙","age":2}`
	p, err := common.Json2Slice(j)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
	}
	for k, v := range p {
		fmt.Printf("%v: %v\n", k, v)
	}
	print(p)
	q, erss := common.Json2Struct(j)
	if erss != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
	}
	print(q)
	// 	var s Serverslice
	// 	str := `{"servers":
	//    [{"serverName":"Guangzhou_Base","serverIP":"127.0.0.1"},
	//    {"serverName":"Beijing_Base","serverIP":"127.0.0.2"}]}`
	// 	errs := json.Unmarshal([]byte(str), &s)
	// 	if errs != nil {
	// 		fmt.Println(errs)
	// 	}
	// 	fmt.Println(s)
	// 	fmt.Println(s.Servers[0].ServerName)
	// m1 := map[string]interface{}{"name": "John", "age": 10}
	// m := common.Map2Json(m1)
	// print(m)
	// m1 := []interface{}{"John", "Baineng", "Zhangsan", 1}
	// m, _ := common.Slice2Json(m1)
	// print(m)
	//var v interface{}
	a := common.Users{Username: "Bianeng", Password: "AZWER13", Age: 400}
	// v = a
	b, _ := common.Struct2Json(a)
	print(b)

	var setting config.Config
	config, err := ioutil.ReadFile("./config/first.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &setting)
	fmt.Println(setting.Name)
	fmt.Println(setting.Addr)
	fmt.Println(setting.HTTPS)
	fmt.Println(setting.SiteNginx.Port)
	fmt.Println(setting.SiteNginx.LogPath)
	fmt.Println(setting.SiteNginx.Path)
	fmt.Println(setting.SiteDatabase.Hostname)
	fmt.Println(setting.SiteDatabase.Port)
	fmt.Println(setting.SiteDatabase.Username)
	fmt.Println(setting.SiteDatabase.Password)
	prometheus.Testflag()

}
