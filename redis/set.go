package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Set() {
	c := P.Get()    //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池
	_, err := c.Do("Set", "abc", 200)
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
	P.Close() //关闭连接池
}
