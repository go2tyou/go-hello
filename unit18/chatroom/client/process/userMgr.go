package process

import (
	"fmt"
	"gocode/helloworld/unit18/chatroom/common/message"
	"gocode/helloworld/unit18/chatroom/client/model"
)

// 客户端维护 map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var curUser model.CurUser // 用户登录成功后，完成对 CurUser 初始化

// 在客户端显示当前在线用户
func outputOnlineUsers() {
	fmt.Println("当前在线用户列表：")
	for id, user := range onlineUsers {
		fmt.Printf("用户ID=%v\t用户状态=%v\t\n", id, user.UserStatus)
	}
}

// 处理返回的 NotifyUserStatusMsg
func updateUserStatus(notifyMsg *message.NotifyUserStatusMsg) {
	user, ok := onlineUsers[notifyMsg.UserId]
	if !ok {
		user = &message.User {
			UserId : notifyMsg.UserId,
		}
	}
	user.UserStatus = notifyMsg.UserStatus
	onlineUsers[notifyMsg.UserId] = user

	outputOnlineUsers()
}