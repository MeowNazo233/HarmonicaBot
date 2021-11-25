package functions

import "strings"

func RuleInto(selif_id uint64, groupid int64, guild_id uint64, channel_id uint64, user_id uint64, nickname string, msg string, msgtype string) string {
	if user_id == selif_id {
		return ""
	}
	if strings.Contains(msg, "占卜") {
		send := Lucky(user_id)
		return send
	}
	if strings.Contains(msg, "今天是什么少女") {
		send := GirlToday(nickname)
		return send
	}
	if strings.Contains(msg, "异世界转生") {
		send := Ysjzs(nickname)
		return send
	}
	if strings.Contains(msg, "卖萌") {
		send := Maimeng(nickname)
		return send
	}
	if strings.Contains(msg, "抽老婆") {
		send := GetWaifu(nickname)
		return send
	}
	if strings.Contains(msg, "wf") {
		send := Warframe_main(msg)
		return send
	}

	return ""
}
