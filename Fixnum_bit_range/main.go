package main

import "fmt"

/*
	* 整型取值范围和作用
  * 有符号整数统一公式为:-2的n-1次幂到2的n-1次幂减一
  * 无符号整数统一公式为:0到2的n次幂减一
|      类型 | 取值范围                                     |
| ------: | :--------------------------------------- |
|    int8 | [-128 , 127]                             |
|   int16 | [-32768 , 32767]                         |
|   int32 | [-2147483648 , 2147483647] Go语言中没有字符类型,所有字符都使用int32存储 |
|   int64 | [-9223372036854775808 , 9223372036854775807] |
|     int | 受限于计算机系统,系统是多少位,int为多少位                  |
|   uint8 | [0 , 255]                                |
|  uint16 | [0 , 65535]                              |
|  uint32 | [0 , 4294967295]                         |
|  uint64 | [0 , 18446744073709551615]               |
|    uint | 受限于计算机系统,系统是多少位,uint为多少位                 |
|    rune | 与int32类似,常用在获取值的Unicode码                 |
|    byte | 与uint8类似.强调值为原始数据.一个字节占用8个二进制            |
| uintptr | 大小不确定,类型取决于底层编程                          |
*/
func main() {
	s := []byte{0,255}
	fmt.Println(s[0])
}
