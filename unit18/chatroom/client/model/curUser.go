package model

import (
	"net"
	"gocode/helloworld/unit18/chatroom/common/message"
)

// 在客户端，很多地方用到

type CurUser struct {
	message.User
	Conn net.Conn `json:"conn"`
}