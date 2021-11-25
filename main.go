package main

import (
	"github.com/MeowNazo233/HarmonicaBot/functions"
	HMC_Bot "github.com/MeowNazo233/HarmonicaBot/server"
)

func action_MessageGuild(eventinfo HMC_Bot.MessageGuild) {

	get_send := functions.RuleInto(eventinfo.SelfTinyID, 0, eventinfo.GuildID, eventinfo.ChannelID, eventinfo.UserID, eventinfo.Sender.Nickname, eventinfo.Message, "guild")
	if get_send != "" {
		HMC_Bot.SendGuildMsg(get_send, eventinfo.GuildID, eventinfo.ChannelID)
	}
}
func action_MessagePrivate(eventinfo HMC_Bot.MessagePrivate) {

	get_send := functions.RuleInto(uint64(eventinfo.SelfID), 0, 0, 0, uint64(eventinfo.UserID), eventinfo.Sender.Nickname, eventinfo.Message, "private")
	if get_send != "" {
		HMC_Bot.SendPrivateMsg(get_send, eventinfo.UserID)
	}
}
func action_MessageGroup(eventinfo HMC_Bot.MessageGroup) {

	get_send := functions.RuleInto(uint64(eventinfo.SelfID), eventinfo.GroupID, 0, 0, uint64(eventinfo.UserID), eventinfo.Sender.Nickname, eventinfo.Message, "group")
	if get_send != "" {
		HMC_Bot.SendGroupMsg(get_send, eventinfo.GroupID)
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
func SendGroupMsg_active(group_id uint64, msg string) {
	HMC_Bot.SendGroupMsg(msg, int64(group_id))
}
func SendPrivateMsg_active(user_id uint64, msg string) {
	HMC_Bot.SendPrivateMsg(msg, int64(user_id))
}
func SendGuildMsg_active(guild_id uint64, channel_id uint64, msg string) {
	HMC_Bot.SendGuildMsg(msg, guild_id, channel_id)
}
