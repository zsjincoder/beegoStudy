package cache

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	"time"
)

func RegisterCache()  {
	Bm , err := cache.NewCache("memory",`{"interval":60}`)
	if err != nil{
		fmt.Println(err)
	}
	Bm.Put("userName", "张三", 10*time.Second)
	fmt.Println(Bm.Get("userName"),10*time.Second)
}
