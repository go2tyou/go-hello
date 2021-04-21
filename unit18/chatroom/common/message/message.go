package message

const (
    LoginMsgType = "LoginMsg"
    LoginResMsgType = "LoginResMsg"
    RegisterMsgType = "RegisterMsg"
    RegisterResMsgType = "RegisterResMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	SmsMsgType = "SmsMsg"
)

// 用户在线的状态
const (
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

// 定义消息
type LoginMsg struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMsg struct {
	Code int `json:"code"` // 状态码 500 -> 未注册; 200 -> 登录成功
	Error string `json:"error"` // 错误信息
	UserIds []int `json:"userIds"` // 增加一个保存用户 ID 的切片
}

type RegisterMsg struct {
	User User `json:"user"`
}

type RegisterResMsg struct {
	Code int `json:"code"` // 400 -> 该用户已被占用; 200 -> 注册成功
	Error string `json:"error"`
}

// 为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMsg struct {
	UserId int `json:"userId"`
	UserStatus int `json:"userStatus"`
}

// 发送消息
type SmsMsg struct {
	Content string `json:"content"`
	User // 匿名结构体 -> 继承
}