package model

// 定义一个用户的结构体
type User struct {
	// 序列化和反序列化 -> tag !!!
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}