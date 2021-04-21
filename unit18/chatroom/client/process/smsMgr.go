package process

import (
	"fmt"
	"encoding/json"
	"gocode/helloworld/unit18/chatroom/common/message"
)

func outputGroupMsg(msg *message.Message) {
	// 1、反序列化 msg.Data
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal res msg err", err)
		return
	}
	// 显示信息
	info := fmt.Sprintf("用户ID= %d 对大家说：\t%s\n", smsMsg.UserId, smsMsg.Content)
	fmt.Println(info)
}