# HarmonicaBot

HarmonicaBot，基于Go-cqhttp的频道（Guild）机器人，使用正向协议

#### 项目初期较原始，实现基本功能

### 目前功能

- 被动消息
- 主动消息

### 快速上手

`go get github.com/MeowNazo233/HarmonicaBot/server` 导入包

```go
package main
import (
	HMC_Bot "github.com/MeowNazo233/HarmonicaBot/server"
	//...
)
func action_guild(eventinfo HMC_Bot.MessageGuild) {
	if eventinfo.Message == "hello" {
		HMC_Bot.SendGuildMsg("Hello World", eventinfo.GuildID, eventinfo.ChannelID)
	}
}
func main() {
	//绑定频道消息处理函数
	HMC_Bot.Listeners.OnGuildMsg = append(HMC_Bot.Listeners.OnGuildMsg, action_guild)
	
	Bot := HMC_Bot.NewBot()
	Bot.Config = HMC_Bot.Config{
		Loglvl:   HMC_Bot.LOGGER_LEVEL_INFO,
		Host:     "0.0.0.0:6700",
		MasterQQ: 1234567890,
		Path:     "/",
	}
	Bot.Run()
}
```

### 参考项目

- [go-Pichubot](https://github.com/0ojixueseno0/go-Pichubot)
- [ZeroBot](https://github.com/wdvxdr1123/ZeroBot)
