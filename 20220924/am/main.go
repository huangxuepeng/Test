package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	res := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	// 连接redis 信息
	// _, err := rgb.Ping(ctx).Result()
	// if err != nil {
	// 	log.Println("连接redis失败: ", err)
	// 	return
	// }

	// fmt.Println("连接成功")

	// 获取所有的keys
	// keys, err := res.Keys(ctx, "*").Result()
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// fmt.Println(keys)

	// err := res.Set(ctx, "key", "value1", time.Second*2).Err()
	// fmt.Println(err)
	cc, _ := res.Get(ctx, "18088630924").Result()
	fmt.Println(cc)

}
