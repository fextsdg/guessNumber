package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 用于判断玩家猜的数字是否正确
func IsTure(answer int, getNum int) bool {
	if getNum < answer {
		fmt.Println("你猜的数字太小了")
	} else if getNum > answer {
		fmt.Println("你猜的数字太大了")
	} else {
		fmt.Println("恭喜你，猜对了！")
		return true
	}
	return false
}
func main() {
	fmt.Println("亲爱的玩家，欢迎来到猜数字游戏的世界！")
	fmt.Println("游戏规则：")
	fmt.Println("1.电脑会随机生成一个1-100的数字。")
	fmt.Println("2.你输入一个数字，电脑会告诉你这个数字是大于还是小于。")
	fmt.Println("3.直到你猜对为止。")
	fmt.Println("开始游戏！")
	fmt.Println("______________________________________")

	//生成随机数字作为答案
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(100)

	result := false
	for !result {
		fmt.Print("请输入一个数字：")
		var getNum int
		fmt.Scanf("%d", &getNum)
		result = IsTure(answer, getNum)
	}
}
