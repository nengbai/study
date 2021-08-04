package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

//json转成map
func Json2Map(j string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(j), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	return m, nil
}

//json转成slice
func Json2Slice(j string) ([]interface{}, error) {
	s := Users{}
	var sli = make([]interface{}, 0)

	err := json.Unmarshal([]byte(j), &s)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	sli = append(sli, s)
	return sli, nil
}

//json转成struct
func Json2Struct(j string) (interface{}, error) {
	var userJSON Users
	err := json.Unmarshal([]byte(j), &userJSON)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
	}
	return userJSON, nil
}

//将map转成json字符串
func Map2Json(m map[string]interface{}) (string, error) {
	s := []map[string]interface{}{}
	s = append(s, m)
	mJSON, err := json.Marshal(s)
	//
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
	}
	return (string(mJSON)), nil
}

// slice to Json
func Slice2Json(arr []interface{}) (string, error) {
	s := []interface{}{}
	s = append(s, arr...)
	mJSON, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
	}
	return (string(mJSON)), nil
}

// struct to json
func Struct2Json(arr Users) (string, error) {
	mJSON, err := json.Marshal(arr)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
	}
	//print(string(mJSON))
	return (string(mJSON)), nil
}

// String to Upper
func UpperCase(str string) string {

	return strings.ToUpper(str)
}

func ReplaceDotToUnderline(str string) string {
	return strings.ReplaceAll(str, ".", "_")
}

func ReplaceUnderlineToDot(str string) string {
	return strings.ReplaceAll(str, "_", ".")
}

func GetFolderList(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	var dirs []string

	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}

	return dirs
}
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
