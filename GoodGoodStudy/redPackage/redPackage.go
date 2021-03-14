package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type redPackage struct {
	RemainSize  int32
	RemainMoney int32
}

func getRandomMoney(redPackage *redPackage) int32 {
	// remainSize 剩余的红包数量
	// redPackage.RemainMoney 剩余的钱
	if redPackage.RemainSize == 1 {
		redPackage.RemainSize--
		return redPackage.RemainMoney
	}
	//rand.Seed(time.Now().UnixNano())
	min := int32(10)
	max := (redPackage.RemainMoney / redPackage.RemainSize) * 2
	tmp, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	money := int32(tmp.Int64())
	if money <= min {
		money = min
	}
	redPackage.RemainSize--
	redPackage.RemainMoney -= money
	return money
}

func main() {
	redPackage := &redPackage{
		RemainSize:  5000,
		RemainMoney: 3000000,
	}
	for i := 0; i < 5000; i++ {
		fmt.Println(getRandomMoney(redPackage))
	}
}
