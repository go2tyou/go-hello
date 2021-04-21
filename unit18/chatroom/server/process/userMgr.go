package process2

import (
	"fmt"
)

// 因为 UserMgr 实例在服务器端有且仅有一个
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对 UserMgr 初始化工作
func init() {
	userMgr = &UserMgr {
		onlineUsers : make(map[int]*UserProcess, 1024),
	}
}

// 完成对 UserMgr 添加
func (this *UserMgr) AddOnlineUsers(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 删除
func (this *UserMgr) DelOnlineUsers(userId int) {
	delete(this.onlineUsers, userId)
}

// 返回所有在线用户
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据 ID 返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户 %d 不存在", userId)
		return
	}
	return
}