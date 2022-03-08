package main

import (
	"fmt"
	"strconv"
	"time"

	"git.code.oa.com/trpc-go/trpc-database/bigcache"
)

func main() {
	cache := bigcache.New(
		//设置分区数，默认256
		bigcache.WithShards(1024),
		//设置生命周期，默认7天
		bigcache.WithLifeWindow(time.Hour),
		//设置清理周期，默认1分钟
		bigcache.WithCleanWindow(-1*time.Second),
		//设置最大数据条数，默认200000，小值可降低初始化使用内存，仅初始化使用
		bigcache.WithMaxEntriesInWindow(50000),
		//设置最大数据长度（单位字节），默认30，小值可降低初始化使用内存，仅初始化使用
		bigcache.WithMaxEntrySize(20),
	)
	var err error
	var v string
	for j := 0; j < 100; j++ {
		key := fmt.Sprintf("k%010d", j)
		v = "v" + strconv.Itoa(j)
		// 写入数据
		err = cache.Set(key, v)
		if err != nil {
			fmt.Println(err)
		}
		// 查找数据
		val, err := cache.GetBytes(key)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("got value: %+v \n", val)
		if j%100 == 0 {
			fmt.Println(v)
		}
	}
}