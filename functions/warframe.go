package functions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Warframe_main(msg string) string {
	if msg == "wf" {
		send := `Warframe查询指令\n当前支持子命令：活动、警报、突击、夜灵平原、金星平原、火卫二平原、裂缝、促销商品、入侵、奸商、达尔沃、小小黑、舰队、电波、仲裁、紫卡+名字`
		return send
	}
	if strings.Contains(msg, "活动") {
		send := wf_events()
		return send
	}
	if strings.Contains(msg, "警报") {
		send := wf_alerts()
		return send
	}
	if strings.Contains(msg, "突击") {
		send := wf_sortie()
		return send
	}
	if strings.Contains(msg, "夜灵平原") {
		send := wf_Ostrons()
		return send
	}
	if strings.Contains(msg, "金星平原") {
		send := wf_Solaris()
		return send
	}
	if strings.Contains(msg, "火卫二平原") {
		send := wf_EntratiSyndicate()
		return send
	}
	if strings.Contains(msg, "裂缝") {
		send := wf_fissures()
		return send
	}
	if strings.Contains(msg, "促销") {
		send := wf_flashSales()
		return send
	}
	if strings.Contains(msg, "入侵") {
		send := wf_invasions()
		return send
	}
	if strings.Contains(msg, "奸商") {
		send := wf_voidTrader()
		return send
	}
	if strings.Contains(msg, "达尔沃") {
		send := wf_dailyDeals()
		return send
	}
	if strings.Contains(msg, "小小黑") {
		send := wf_persistentEnemies()
		return send
	}
	if strings.Contains(msg, "舰队") {
		send := wf_constructionProgress()
		return send
	}
	if strings.Contains(msg, "电波") {
		send := wf_nightwave()
		return send
	}
	if strings.Contains(msg, "仲裁") {
		send := wf_arbitration()
		return send
	}

	if strings.Contains(msg, "紫卡") {
		item := strings.ReplaceAll(msg, "wf紫卡", "")
		send := wf_rm(item)
		return send
	}
	if strings.Contains(msg, "交易") {
		item := strings.ReplaceAll(msg, "wf交易", "")
		send := wf_wm(item)
		return send
	}
	return ""
}

//活动 http://nymph.rbq.life:3000/wf/robot/events
func wf_events() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/events")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//警报 http://nymph.rbq.life:3000/wf/robot/alerts
func wf_alerts() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/alerts")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//突击 http://nymph.rbq.life:3000/wf/robot/sortie
func wf_sortie() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/sortie")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//集团任务 (仅限dev、detail的json数据) http://nymph.rbq.life:3000/wf/robot/syndicateMissions
func wf_syndicateMissions() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/events")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
} //JSON数据待重做

//地球赏金 (仅限robot的string数据) http://nymph.rbq.life:3000/wf/robot/Ostrons
func wf_Ostrons() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/Ostrons")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//金星赏金 (仅限robot的string数据) http://nymph.rbq.life:3000/wf/robot/sortie
func wf_Solaris() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/sortie")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//火卫二赏金 (仅限robot的string数据) http://nymph.rbq.life:3000/wf/robot/EntratiSyndicate
func wf_EntratiSyndicate() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/EntratiSyndicate")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//裂缝 http://nymph.rbq.life:3000/wf/robot/fissures
func wf_fissures() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/fissures")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//促销 http://nymph.rbq.life:3000/wf/robot/flashSales
func wf_flashSales() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/flashSales")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//入侵 http://nymph.rbq.life:3000/wf/robot/invasions
func wf_invasions() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/invasions")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//奸商 http://nymph.rbq.life:3000/wf/robot/voidTrader
func wf_voidTrader() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/voidTrader")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//达尔沃 http://nymph.rbq.life:3000/wf/robot/dailyDeals
func wf_dailyDeals() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/dailyDeals")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//小小黑 http://nymph.rbq.life:3000/wf/robot/persistentEnemies
func wf_persistentEnemies() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/persistentEnemies")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//舰队 http://nymph.rbq.life:3000/wf/robot/constructionProgress
func wf_constructionProgress() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/constructionProgress")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//电波 http://nymph.rbq.life:3000/wf/robot/nightwave
func wf_nightwave() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/nightwave")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//仲裁 http://nymph.rbq.life:3000/wf/robot/arbitration
func wf_arbitration() string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wf/robot/arbitration")
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//紫卡价格 http://nymph.rbq.life:3000/rm/robot/+itemname
func wf_rm(item string) string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/rm/robot/" + item)
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}

//warframe market http://nymph.rbq.life:3000/wm/robot/+itemname
func wf_wm(item string) string {
	resp, resp_err := http.Get("http://nymph.rbq.life:3000/wm/robot/" + item)
	if resp_err != nil {
		// handle error
		fmt.Println(resp_err)
	}
	defer resp.Body.Close()

	body, body_err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if body_err != nil {
		// handle error
		fmt.Println(body_err)
	}
	return string(body)
}
