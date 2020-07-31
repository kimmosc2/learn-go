package main

import (
	"fmt"
)

/*
位操作运算符 ^ 作为二元运算符时是按位异或（ XOR） ， 当用作一元运算符时表示按位取
反； 也就是说， 它返回一个每个bit位都取反的数。 位操作运算符 &^ 用于按位置零（ AND
NOT） ： 如果对应y中bit位为1的话, 表达式 z = x &^ y 结果z的对应的bit位为0， 否则z对应的
bit位等于x相应的bit位的值
*/
func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	var z uint8 = 254
	fmt.Printf("%08b\n", x) // 00100010
	fmt.Printf("%08b\n", y) // 00000110
	fmt.Printf("%08b\n", z) // 11111110

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5} 个人理解:用y去取x的反，只要y上有这一位，x上的这一位就要取反
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5} 用y清楚x的位
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

}
