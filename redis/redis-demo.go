package main

import (
	"encoding/json"
	"fmt"
	"study/common"

	"github.com/garyburd/redigo/redis"
)

/*
* 1. 以key为对象进行操作整合一起组成接口
 */
type ManagerRedisKey interface {
	Strset(value interface{})
	RedisGetStr() string
	Strgetset(value interface{}) string
	ExpireKey(extime int)
	CheckKey() bool
	DelKey() error
	StrsetJson(value interface{})
	StrgetJson() interface{}
}

type RKey struct {
	key string
}

/*
* set(key, value)：给数据库中名称为key的string赋予值value
 */
func (manrk *RKey) Strset(value interface{}) {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	_, err := c.Do("SET", manrk.key, value)
	if err != nil {
		fmt.Println("set error", err.Error())
	} else {
		fmt.Println("set ok.")
	}
}

/*
* set(key, value)：给数据库中名称为key的string赋予值value
 */
func (manrk *RKey) StrsetEX(value interface{}, extime int) {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	_, err := c.Do("SET", manrk.key, value, "EX", extime)
	if err != nil {
		fmt.Println("set error", err.Error())
	} else {
		fmt.Println("set ok.")
	}
}

/*
* get(key)：返回数据库中名称为key的string的value
 */
func (manrk *RKey) RedisGetStr() string {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	res, err := redis.String(c.Do("GET", manrk.key))
	if err != nil {
		fmt.Println("GET error", err.Error())
		return ""
	} else {
		return res
	}
}

/*
* getset(key, value)：给名称为key的string赋予上一次的value
 */
func (manrk *RKey) Strgetset(v interface{}) string {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	res, err := redis.String(c.Do("GETSET", manrk.key, v))
	if err != nil {
		fmt.Println("GETSET error", err.Error())
		return ""
	} else {
		fmt.Println("GETSET ok.")
		return res
	}
}

/*
* 给keys 设置时间
 */
func (manrk *RKey) ExpireKey(extime int) {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", manrk.key, extime)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

/*
* 判断是否存在
 */
func (manrk *RKey) CheckKey() bool {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", manrk.key))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return exist
	}
}

/*
* 删除key
 */
func (manrk *RKey) DelKey() error {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	_, err := c.Do("DEL", manrk.key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

/*
* 存入json数据
 */
func (manrk *RKey) StrsetJson(value interface{}) {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	datas, _ := json.Marshal(value)
	_, err := c.Do("SET", manrk.key, datas)
	if err != nil {
		fmt.Println("set error", err.Error())
	} else {
		fmt.Println("set ok.")
	}
}

/*
* 获取json数据
 */
func (manrk *RKey) StrgetJson() interface{} {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	var getdatas map[string]string
	valueget, err := redis.Bytes(c.Do("GET", manrk.key))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	errshal := json.Unmarshal(valueget, &getdatas)
	if errshal != nil {
		return 0
	}
	return getdatas
}

/*
* 2. 以批量操作数据整合一起组成接口
 */
type ManagerRedisBatch interface {
	Strmset()
	Strmget() []string
}
type BatchDatas struct {
	datas interface{}
}

/*
* mset(key N, value N)：批量设置多个string的值
 */
func (manrb *BatchDatas) Strmset() {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	_, err := c.Do("mset", redis.Args{}.AddFlat(manrb.datas)...)
	if err != nil {
		fmt.Println("mset error", err.Error())
	} else {
		fmt.Println("ok")
	}
}

/*
* mget(key1, key2,…, key N)：返回库中多个string的value
 */
func (manrb *BatchDatas) Strmget() []string {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()
	list := []string{}
	res, err := redis.Values(c.Do("mget", redis.Args{}.AddFlat(manrb.datas)...))
	if err != nil {
		fmt.Println("mset error", err.Error())
		return list
	} else {
		for _, v := range res {
			fmt.Println(string(v.([]byte)))
			list = append(list, string(v.([]byte)))
		}
		return list
	}
}

func test() {
	keys := "test1"
	var test1 ManagerRedisKey = &RKey{key: keys}
	test1.Strset("aaa1")
	fmt.Println(test1.RedisGetStr())
	fmt.Println(test1.Strgetset("aaa2"))
	test1.ExpireKey(100)
	fmt.Println(test1.CheckKey())
	test1.DelKey()

	datas1 := map[string]string{"username": "666", "phonenumber": "888"}
	var test2 ManagerRedisBatch = &BatchDatas{datas: datas1}
	test2.Strmset()

	var test3 ManagerRedisBatch = &BatchDatas{datas: []string{"username", "phonenumber"}}
	fmt.Println(test3.Strmget())

	var test4 ManagerRedisKey = &RKey{key: "test4"}
	test4.StrsetJson(map[string]string{"aa1": "111", "aa2": "222"})
	fmt.Println(test4.StrgetJson())
	fmt.Println(test4.RedisGetStr())
}

/*
* 管道 按照队列先进先出的原则进行send,receive操作
 */
func sendr(keyName, values string) {
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()

	c.Send("SET", keyName, values)
	c.Send("GET", keyName)
	c.Flush()
	c.Receive()                   // reply from SET
	valueGet, errr := c.Receive() // reply from GET
	fmt.Println(redis.String(valueGet, errr))

}

/*
* 发布/订阅
 */
func Subs() { //订阅者
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()

	psc := redis.PubSubConn{c}
	psc.Subscribe("channel1") //订阅channel1频道
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

func Push(message string) { //发布者
	c := common.NewRedisPool(common.RedisURL, 1).Get()
	defer c.Close()

	_, err1 := c.Do("PUBLISH", "channel1", message)
	if err1 != nil {
		fmt.Println("pub err: ", err1)
		return
	}
}

func main() {

	test()
	sendr("myname", "Baineng")
	//发布/订阅
	go Subs()
	go Push("mange *******...")
}
