package main

import (
	HMC_Bot "github.com/MeowNazo233/HarmonicaBot/server"
)

func action_MessageGuild(eventinfo HMC_Bot.MessageGuild) {
	if eventinfo.Message == "hello" {
		HMC_Bot.SendGuildMsg("Hello [频道消息]", eventinfo.GuildID, eventinfo.ChannelID)
	}
}

func action_MessagePrivate(eventinfo HMC_Bot.MessagePrivate) {
	if eventinfo.Message == "hello" {
		HMC_Bot.SendPrivateMsg("Hello [私聊消息]", eventinfo.UserID)
	}
}

func action_MessageGroup(eventinfo HMC_Bot.MessageGroup) {
	if eventinfo.Message == "hello" {
		HMC_Bot.SendGroupMsg("Hello [群聊消息]", eventinfo.GroupID)
	}
}
func main() {
	//创建事件函数
	HMC_Bot.Listeners.OnGuildMsg = append(HMC_Bot.Listeners.OnGuildMsg, action_MessageGuild)
	HMC_Bot.Listeners.OnPrivateMsg = append(HMC_Bot.Listeners.OnPrivateMsg, action_MessagePrivate)
	HMC_Bot.Listeners.OnGroupMsg = append(HMC_Bot.Listeners.OnGroupMsg, action_MessageGroup)

	Bot := HMC_Bot.NewBot()
	Bot.Config = HMC_Bot.Config{
		Loglvl:   HMC_Bot.LOGGER_LEVEL_INFO,
		Host:     "0.0.0.0:6700",
		MasterQQ: 779019185,
		Path:     "/",
	}
	Bot.Run()
}
