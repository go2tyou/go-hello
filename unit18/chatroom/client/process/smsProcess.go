package process

import (
	"fmt"
	"encoding/json"
	"gocode/helloworld/unit18/chatroom/common/message"
	"gocode/helloworld/unit18/chatroom/client/utils"
)

type SmsProcess struct {

}

// 发送群聊的消息
func (this *SmsProcess) SendGroupMsg(content string) (err error) {
	// 1、创建一个 msg
	var msg message.Message
	msg.Type = message.SmsMsgType
	// 2、创建一个 SmsMsg
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = curUser.UserId
	smsMsg.UserStatus = curUser.UserStatus
	// 3、序列化 smsMsg
	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("json.Marshal smsMsg err", err)
		return
	}
	msg.Data = string(data)
	// 4、序列化 msg
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal msg err", err)
		return
	}
	// 5、将 msg 发送给服务器
	tf := &utils.Transfer {
		Conn : curUser.Conn,
	}
	// 发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg send msg err", err)
		return
	}
	return
}