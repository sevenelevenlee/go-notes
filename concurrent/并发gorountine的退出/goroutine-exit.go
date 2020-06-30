package main

import (
	"fmt"
)

/*
	使用for-range退出
		使用场景，当协程只从1个channel读取数据的时候，range处理channel, 当channel关闭时， 协程自动退出
*/

func main() {
	var ch = make(chan int)
	/*
		使用for-range退出
			使用场景，当协程只从1个channel读取数据的时候，range处理channel, 当channel关闭时， 协程自动退出
	*/
	go func(in <-chan int) {
		for v := range in {
			//当in关闭时候，输出0
			fmt.Println(v)
		}
	}(ch)

	/*
		使用ok退出：
			针对for-select无法感知channel关闭的情景，继续在关闭的channel上读，会读到0值或nil
			继续在关闭的通道上写，会panic
	*/
	go func() {
		
	}()
}