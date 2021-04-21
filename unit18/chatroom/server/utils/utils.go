package utils

import (
	"fmt"
	"net"
	"gocode/helloworld/unit18/chatroom/common/message"
	"encoding/json"
	"encoding/binary"
	"errors"
)

// 将这些方法关联到结构体中
type Transfer struct {
	// 分析应该有哪些字段
	Conn net.Conn
	Buf [8096]byte // 传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (msg message.Message, err error) {
	// buf := make([]byte, 8096)
	// conn.Read 只有在 conn 没有被关闭的情况下，才会阻塞
	// 如果客户端关闭了 conn，则不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}
	// fmt.Println("读到的长度为", buf[:4])
	// 把 buf[:4] -> uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	// 根据 pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen]) // 从 conn 读 pkgLen 字节放入 buf
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	// 把 pkgLen 反序列化为 message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &msg) // &msg !!!
	if err != nil {
		fmt.Println("json.Unmarshal msg err", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) len(data) err", err)
		return
	}
	// 发送 data 本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err", err)
		return
	}
	return
}