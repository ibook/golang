package main

import "fmt"
import "time"

func FormatTime() {
    // get current timestamp
    currentTime := time.Now().Local()

    //print time
    fmt.Println(currentTime)

    //format Time, string type
    newFormat := currentTime.Format("2006-01-02 15:04:05.000")
    fmt.Println(newFormat)


    //build Time 2016-02-17 23:59:59.999, DateTime type
    myTime := time.Date(0162, time.February, 17, 23, 59, 59, 999, time.UTC)

    //output the myTime
    fmt.Println("MyTime:", myTime.Format("2006-01-02 15:04:05.000"))

    //current milliseconds
    fmt.Println("milliseconds:", time.Now().UnixNano()/int64(time.Millisecond))

    //TODO Changing time layout(form)
}

func main(){

	
	// 当前时间戳
	fmt.Println(time.Now().Unix())
	
	// str格式化时间,	当前格式化时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) 
	
	
	// 时间戳转str格式化时间
	str_time := time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")
	fmt.Println(str_time)
	
	
	// str格式化时间转时间戳 这个比较麻烦
	the_time := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local)
	unix_time := the_time.Unix()
	fmt.Println(unix_time)
	
	// 还有一种方法,使用time.Parse
	the_time, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
	if err == nil {
			unix_time := the_time.Unix()
		fmt.Println(unix_time)		
	}

	FormatTime()
}