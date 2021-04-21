package process

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
	"encoding/binary"
	"gocode/helloworld/unit18/chatroom/common/message"
	"gocode/helloworld/unit18/chatroom/client/utils"
)

type UserProcess struct {

}

// 登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// 开始定协议
	// fmt.Println("您输入的 ID 和用户名分别是", userId, userPwd)
	// return nil

	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	// 延时关闭
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err......", err)
		return
	}

	// 准备通过 conn 发送消息给服务器
	var msg message.Message
	msg.Type = message.LoginMsgType
	// 创建 LoginMsg 结构体
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd
	// 将 loginMsg 序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal loginMsg err", err)
		return
	}
	// 把 data 赋值给 msg.Data
	msg.Data = string(data)
	// 将 msg 进行序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	// 此时 data 就是要发送的信息
	// 先把 data 的 len 发送给服务器 len(data) -> []byte
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) len(data) err", err)
		return
	}
	fmt.Printf("客户端发送消息的长度=%d 内容=%s\n", len(data), string(data))
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err", err)
		return
	}
	// 此处还需处理服务器返回的消息
	// 创建一个 Transfer 实例
	tf := &utils.Transfer {
		Conn : conn,
	}
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err", err)
		return
	}
	// 将 msg 的 Data 反序列化为 LoginResMsg
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if loginResMsg.Code == 200 {
		// fmt.Println("登录成功")

		// 对 curUser 初始化
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline

		// 显示当前在线用户列表
		fmt.Println("当前在线用户列表：")
		for _, v := range loginResMsg.UserIds {
			if v == userId {
				continue
			}
			fmt.Printf("用户ID=%d\t\n", v)
			// 完成客户端 onlineUsers 初始化
			user := &message.User {
				UserId : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v] = user
		}

		// 这里还需要在客户端启动一个协程
		// 该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		// 接收并显示在客户端的终端
		go ServerProcessMsg(conn)

		// 1、显示登录成功的菜单(循环显示)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMsg.Error)
	}
	return
}

// 注册
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	// 延时关闭
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err......", err)
		return
	}

	// 准备通过 conn 发送消息给服务器
	var msg message.Message
	msg.Type = message.RegisterMsgType
	// 创建 RegisterMsg 结构体
	var registerMsg message.RegisterMsg
	registerMsg.User.UserId = userId
	registerMsg.User.UserPwd = userPwd
	registerMsg.User.UserName = userName
	// 将 RegisterMsg 序列化
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("json.Marshal registerMsg err", err)
		return
	}
	// 把 data 赋值给 msg.Data
	msg.Data = string(data)
	// 将 msg 进行序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	// 创建一个 Transfer 实例
	tf := &utils.Transfer{
		Conn : conn,
	}
	// 发送 data 给服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Register send msg err", err)		
	}
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Register readPkg err", err)
		return
	}
	// 将 msg 的 Data 反序列化为 RegisterResMsg
	var registerResMsg message.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)
	if registerResMsg.Code == 200 {
		fmt.Println("注册成功，重新登录.......")
	} else {
		fmt.Println(registerResMsg.Error)
	}
	os.Exit(0)
	return
}