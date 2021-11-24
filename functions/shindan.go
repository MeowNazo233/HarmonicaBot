package functions

import (
	"fmt"

	"github.com/FloatTech/AnimeAPI/shindanmaker"
)

//今天是什么少女 code:今天是什么少女
func GirlToday(nickname string) string {
	text, err := shindanmaker.Shindanmaker(162207, nickname)
	if err != nil {
		fmt.Println(err)
	}

	return text
}

//异世界转生
func Ysjzs(nickname string) string {
	text, err := shindanmaker.Shindanmaker(587874, nickname)
	if err != nil {
		fmt.Println(err)
	}

	return text
}

//卖萌
func Maimeng(nickname string) string {
	text, err := shindanmaker.Shindanmaker(360578, nickname)
	if err != nil {
		fmt.Println(err)
	}

	return text
}

//抽老婆
func GetWaifu(nickname string) string {
	text, err := shindanmaker.Shindanmaker(1075116, nickname)
	if err != nil {
		fmt.Println(err)
	}

	return text
}

// engine.OnPrefix("卖萌", number(360578)).SetBlock(true).FirstPriority().Handle(handle)
// engine.OnPrefix("抽老婆", number(1075116)).SetBlock(true).FirstPriority().Handle(handle)
