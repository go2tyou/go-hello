package model

import (
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"gocode/helloworld/unit18/chatroom/common/message"
)

// 在服务器启动后，就初始化一个 UserDao 实例
// 把它做成全局的，在需要和 redis 操作时，就可直接使用
var (
	MyUserDao *UserDao
)

// 定义一个 UserDao，完成对 User 的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式创建一个 UserDao 实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao {
		pool : pool,
	}
	return
}

// 根据用户 ID 返回一个 User 实例
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		// 在 users 哈希中，没有找到 id
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	// 把 res 反序列化成 User 实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal user err", err)
		return
	}
	return
}

// 登录校验
// 1、完成对用户的验证
// 2、如果通过，则返回一个 User 实例
// 3、如果有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	// 先从 UserDao 中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	// 密码验证
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	// 先从 UserDao 中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	// 此时，说明 ID 在 redis 还没有，则可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal Register user err", err)
		return
	}
	// 入库
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误......", err)
		return
	}
	return
}