package common

import (
	"bytes"
	"encoding/gob"
	"io"
)

type P struct {
	Username string
	Password string
	Age      int
}

func Encode(data interface{}) *bytes.Buffer {
	//Buffer类型实现了io.Writer接口
	var buf bytes.Buffer
	//得到编码器
	enc := gob.NewEncoder(&buf)
	//调用编码器的Encode方法来编码数据data
	enc.Encode(data)
	//编码后的结果放在buf中
	return &buf
}

func Decode(data interface{}) *P {
	d := data.(io.Reader)
	//获取一个解码器，参数需要实现io.Reader接口
	dec := gob.NewDecoder(d)
	var q P
	//调用解码器的Decode方法将数据解码，用Q类型的q来接收
	dec.Decode(&q)
	return &q
}
