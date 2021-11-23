package main

import (
	HMC_Bot "github.com/MeowNazo233/HarmonicaBot/server"
)

func action(eventinfo HMC_Bot.MessageGuild) {
	if eventinfo.Message == "hello" {
		HMC_Bot.SendGuildMsg("Hello", 1234567890987654321, 1234567)
	}
}

func main() {
	//创建事件函数
	HMC_Bot.Listeners.OnGuildMsg = append(HMC_Bot.Listeners.OnGuildMsg, action)
	Bot := HMC_Bot.NewBot()
	Bot.Config = HMC_Bot.Config{
		Loglvl:   HMC_Bot.LOGGER_LEVEL_INFO,
		Host:     "0.0.0.0:6700",
		MasterQQ: 123456789,
		Path:     "/",
	}
	Bot.Run()
}
