package process2

import (
	"fmt"
	"net"
	"encoding/json"
	"gocode/helloworld/unit18/chatroom/common/message"
	"gocode/helloworld/unit18/chatroom/server/utils"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMsg(msg *message.Message) {
	// 遍历服务器端 onlineUsers map[int]*UserProcess
	// 将消息转发取出
	// 取出 msg 的内容 SmsMsg
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal smsMsg err", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Unmarshal msg err", err)
		return	
	}

	for id, up := range userMgr.onlineUsers {
		// 不要发给自己
		if id == smsMsg.UserId {
			continue
		}
		this.SendMsgToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMsgToEachOnlineUser(data []byte, conn net.Conn) {
	// 创建一个 Transfer 实例，发送 data
	tf := &utils.Transfer {
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败......", err)
	}
}