package test

import (
	"fmt"
	"study/common"
)

type Users struct {
	Username string
	Password string
	Age      int
}

func main() {
	user := Users{
		"Beijing",
		"123456-hakdf",
		35,
	}

	userBytes := common.Encode(user)
	// fmt.Println(userBytes)
	q := common.Decode(userBytes)
	fmt.Println(q)

	fmt.Println(q.Username, q.Password, q.Age)

	//struct to json
	u := Users{}
	u.Username = "wangfang"
	u.Password = "qaz123d"
	u.Age = 18
	common.Struct2Json(common.Users(u))

	// slice to Json
	arr := []interface{}{"Apple", "Orange", "Banana"}

	common.Slice2Json(arr)

	//将map转成json字符串
	m := map[string]interface{}{"浙江": "杭州", "湖南": "长沙"}
	common.Map2Json(m)

	//json转成struct
	jsonStr := `{"Username":"Tom","Password":"123456","FriendName":["Li","Fei"]}`
	common.Json2Struct(jsonStr)
	//json转成slice
	jsonFruit := `["Apple","Orange","Banana"]`
	common.Json2Slice(jsonFruit)
	//json转成map

	j := `{"浙江":"杭州","湖南":"长沙"}`
	// j := `{"浙江","湖南"}`
	common.Json2Map(j)

}
