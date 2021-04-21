package main

import (
	"fmt"
	"time"
	"strconv"
)

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	// 1、获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now) // now=2020-11-25 21:38:47.7559268 +0800 CST m=+0.005024301 now type=time.Time

	// 2、获取其他相关信息
	// year=2020 month=11 day=25 hour=21 minute=45 second=50
	fmt.Printf("year=%v month=%v day=%v hour=%v minute=%v second=%v\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	// 3、格式化日期和时间
	// 方式一 同 2
	dateStr := fmt.Sprintf("%d-%d-%d %d:%d:%d\n", now.Year(), int(now.Month()), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr=", dateStr) // dateStr= 2020-11-25 21:52:36
	// 方式二 数字固定，必须这样
	fmt.Printf(now.Format("2006/01/02 15:04:05") + "\n") // 2020/11/25 21:54:08

	// 4、时间常量应用
	i := 0
	for {
		i++
		fmt.Println(i)
		// 休眠
		// time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 10)
		if i == 10 {
			break
		}
	}

	// 5、time.Time Unix() UnixNano()
	fmt.Printf("unix=%v unixnano=%v\n", now.Unix(), now.UnixNano()) // unix=1606313786 unixnano=1606313786411224800

	// 6、测试代码执行时间
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("process test() time=%vs\n", end - start)
}
