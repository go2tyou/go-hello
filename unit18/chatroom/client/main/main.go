package main

import (
    "fmt"
    "os"
    "gocode/helloworld/unit18/chatroom/client/process"
)

var (
    userId int
    userPwd string
    userName string
)

func main()  {
    // // 接收用户的选择
    var key int
    // // 判断是否还继续显示菜单
    // var loop = true
    for {
        fmt.Println("`````````````````````欢迎登录多人聊天系统``````````````````````")
        fmt.Println("\t\t\t 1 登录聊天室")
        fmt.Println("\t\t\t 2 注册用户")
        fmt.Println("\t\t\t 3 退出系统")
        fmt.Println("\t\t\t 请选择(1~3)：")

        fmt.Scanf("%d\n", &key)
        switch key {
        case 1:
            fmt.Println("登录聊天室......")
            fmt.Println("请输入的ID：")
            fmt.Scanf("%d\n", &userId)
            fmt.Println("请输入的密码：")
            fmt.Scanf("%s\n", &userPwd)
            up := &process.UserProcess{}
            up.Login(userId, userPwd)
            // loop = false
        case 2:
            fmt.Println("注册用户......")
            fmt.Println("请输入的ID：")
            fmt.Scanf("%d\n", &userId)
            fmt.Println("请输入的密码：")
            fmt.Scanf("%s\n", &userPwd)
            fmt.Println("请输入的名字：")
            fmt.Scanf("%s\n", &userName)
            up := &process.UserProcess{}
            up.Register(userId, userPwd, userName)
            // loop = false
        case 3:
            fmt.Println("退出系统......")
            // loop = false
            os.Exit(0)
        default:
            fmt.Println("您的输入有误，请重新输入：")	
        }
    }

    // // 根据用户的输入，显示新的提示信息
    // if key == 1 {
    //     // 用户登录
        
    // } else if key == 2 {
    //     fmt.Println("用户注册")
    // } else if key == 3 {
        
    // } else {

    // }
}