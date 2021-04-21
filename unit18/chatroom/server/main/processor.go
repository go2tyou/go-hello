package main

import (
	"fmt"
	"net"
	"gocode/helloworld/unit18/chatroom/common/message"
	"gocode/helloworld/unit18/chatroom/server/utils"
	"gocode/helloworld/unit18/chatroom/server/process"
	"io"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端发送的消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMsg(msg *message.Message) (err error) {
	// fmt.Println("客户端发送了群聊消息......")
	switch msg.Type {
	case message.LoginMsgType:
		up := &process2.UserProcess {
			Conn : this.Conn,
		}
		err = up.ServerProcessLogin(msg)
	case message.RegisterMsgType:
		up := &process2.UserProcess {
			Conn : this.Conn,
		}
		err = up.ServerProcessRegister(msg)
	case message.SmsMsgType:
		smsP := &process2.SmsProcess{}
		smsP.SendGroupMsg(msg)
	default:
		fmt.Println("消息类型不存在，无法处理......")	
	}
	return
}

func (this *Processor) process2() (err error) {
	for {
		// 封装成一个函数
		// 先创建一个 Transfer，完成读包任务
		tf := &utils.Transfer {
			Conn : this.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭连接，服务器端也退出......")
				return err
			}
			fmt.Println("readPkg(conn) err", err)
			return err	 
		}
		// fmt.Println("msg=", msg)

		err = this.serverProcessMsg(&msg)
		if err != nil {
			return err
		}
	}
}