package functions

import (
	HMC_Bot "github.com/MeowNazo233/HarmonicaBot/server"
)

func Fun_Input(eventinfo HMC_Bot.MessageGuild) string {
	if eventinfo.Message == "Hello" {
		Send_Messag := `Heloo World`
		return Send_Messag
	}
	return ""
}
