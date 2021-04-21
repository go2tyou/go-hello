package process

import (
	"fmt"
	"os"
	"net"
	"encoding/json"
	"gocode/helloworld/unit18/chatroom/client/utils"
	"gocode/helloworld/unit18/chatroom/common/message"
)

// 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("``````````````````````恭喜XXX登录成功````````````````````````")
	fmt.Println("\t\t\t 1、显示在线用户列表")
	fmt.Println("\t\t\t 2、发送消息")
	fmt.Println("\t\t\t 3、信息列表")
	fmt.Println("\t\t\t 4、退出系统")
	fmt.Println("\t\t\t 请选择(1~4)：")
	var key int
	var content string
	fmt.Scanf("%d\n", &key)
	
	smsProcess := &SmsProcess{}

	switch key {
	case 1:
		// fmt.Println("显示在线用户列表......")
		outputOnlineUsers()
	case 2:
		// fmt.Println("发送消息......")
		fmt.Println("您想对大家说的什么：")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMsg(content)
	case 3:
		fmt.Println("信息列表......")
	case 4:
		fmt.Println("您选择退出了系统......")
		os.Exit(0)
	default:
		fmt.Println("您输入的选项不对，请重新输入......")
	}
}

// 和服务器端保持通讯
func ServerProcessMsg(conn net.Conn) {
	// 创建一个 Transfer 实例，不停的读取服务器发送的消息
	tf := &utils.Transfer {
		Conn : conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息......")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err......", err)
		}
		// fmt.Println("msg=", msg)
		switch msg.Type {
		case message.NotifyUserStatusMsgType:
			// 1、取出 NotifyUserStatusMsg
			var notifyMsg message.NotifyUserStatusMsg
			json.Unmarshal([]byte(msg.Data), &notifyMsg)
			// 2、把这个用户的状态保存到客户端的 map 中 -> onlineUsers
			updateUserStatus(&notifyMsg)
		case message.SmsMsgType: // 有群发消息
			outputGroupMsg(&msg)
		default:
			fmt.Println("服务器端返回了未知消息类型")	
		}
	}
}