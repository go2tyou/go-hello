package process2

import (
	"fmt"
	"net"
	"encoding/json"
	"gocode/helloworld/unit18/chatroom/common/message"
	"gocode/helloworld/unit18/chatroom/server/utils"
	"gocode/helloworld/unit18/chatroom/server/model"
)

type UserProcess struct {
	Conn net.Conn
	UserId int // 判断是哪个用户
}

// userId 通知其他所有在线用户
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	// 遍历 userMgr.onlineUsers 然后逐个发送 NotifyUserStatusMsg
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		// 开始通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	// 组装 NotifyUserStatusMsg
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType

	var notifyMsg message.NotifyUserStatusMsg
	notifyMsg.UserId = userId
	notifyMsg.UserStatus = message.UserOnline
	// 将 NotifyUserStatusMsg 序列化
	data, err := json.Marshal(notifyMsg)
	if err != nil {
		fmt.Println("json.Marshal notifyMsg err", err)
		return
	}
	msg.Data = string(data)
	// 对 msg 序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal msg err", err)
		return
	}
	// 发送
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyMsg send err", err)
		return
	}
}

// 处理登录
func (this *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	// 1、先从 msg 中取出 msg.Data，并反序列成 LoginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal loginMsg err", err)
		return
	}

	// 2、先声明一个 Message，并完成赋值
	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType
	// 再声明一个 LoginResMsg
	var loginResMsg message.LoginResMsg
	
	// 校验用户名、密码
	// if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
	// 	loginResMsg.Code = 200
	// } else {
	// 	loginResMsg.Code = 500
	// 	loginResMsg.Error = "该用户不存在，请注册再使用......"
	// }
	// 现在需要到 redis 数据库完成验证
	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)	
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		} else if err == model.ERROR_USER_EXISTS {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 505
			loginResMsg.Error = "服务器内部错误......"
		}
	} else {
		loginResMsg.Code = 200
		// 此处，因为用户登录成功，应该把登录成功的用户放入到 userMgr 中
		// 登录成功的用户 ID 赋值给 this
		this.UserId = loginMsg.UserId
		userMgr.AddOnlineUsers(this)
		// 通知其他在线的用户我上线了
		this.NotifyOthersOnlineUser(loginMsg.UserId)
		// 将当前在线的用户 ID 放入 loginResMsg
		// 遍历 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMsg.UserIds = append(loginResMsg.UserIds, id)
		}
		fmt.Println(user, "登录成功......")
	}

	// 3、将 LoginResMsg 序列化
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json.Marshal loginResMsg err", err)
		return
	}
	// 4、将 data 赋值给 resMsg
	resMsg.Data = string(data)
	// 5、对 resMsg 序列化
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal resMsg err", err)
		return
	}
	// 6、发送 data，将其封装到 writePkg
	// 因为使用了分层模式(MVC)，先创建一个 Transfer，然后读取
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 处理注册
func (this *UserProcess) ServerProcessRegister(msg *message.Message) (err error) {
	// 1、先从 msg 中取出 msg.Data，并反序列成 RegisterMsg
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("json.Unmarshal registerMsg err", err)
		return
	}
	// 2、先声明一个 Message，并完成赋值
	var resMsg message.Message
	resMsg.Type = message.RegisterResMsgType
	// 再声明一个 LoginResMsg
	var registerResMsg message.RegisterResMsg
	// 现在需要到 redis 数据库完成注册
	err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMsg.Code = 400
			registerResMsg.Error = err.Error()
		} else {
			registerResMsg.Code = 506
			registerResMsg.Error = "注册发生未知错误......"
		}
	} else {
		registerResMsg.Code = 200
	}
	// 3、将 RegisterResMsg 序列化
	data, err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("json.Marshal registerResMsg err", err)
		return
	}
	// 4、将 data 赋值给 resMsg
	resMsg.Data = string(data)
	// 5、对 resMsg 序列化
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal resMsg err", err)
		return
	}
	// 6、发送 data，将其封装到 writePkg
	// 因为使用了分层模式(MVC)，先创建一个 Transfer，然后读取
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}