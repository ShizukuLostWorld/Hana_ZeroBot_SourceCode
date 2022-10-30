package funwork

import (
	"math/rand"
	"time"

	"github.com/FloatTech/zbputils/ctxext"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

// debug
var fail = "获取精华消息失败喵~可能是这条信息在数据库中无法查询~"

func init() {
	/* engine.OnRegex(`^查找信息(\d+)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		if !limitinfo.Load(ctx.Event.UserID).Acquire() {
			return
		}
		reachinfo, _ := strconv.ParseInt(ctx.State["regex_matched"].([]string)[1], 10, 64)
		getUserInfo := ctx.GetGroupMemberInfo(ctx.Event.GroupID, reachinfo, false)
		tempUserInfo := getUserInfo.String()
		userName := ctx.CardOrNickName(reachinfo)
		userSexInfo := gjson.Get(tempUserInfo, "sex")
		userJoinTimeUnix := gjson.Get(tempUserInfo, "join_time").Int()
		userJoinTime := time.Unix(userJoinTimeUnix, 0).Format("2006-01-02 03:04:05 PM")
		lastSendTimeUnix := gjson.Get(tempUserInfo, "last_sent_time").Int()
		userLastSendTIme := time.Unix(lastSendTimeUnix, 0).Format("2006-01-02 03:04:05 PM")
		userUnfriendly := gjson.Get(tempUserInfo, "unfriendly")
		userHonorTitle := gjson.Get(tempUserInfo, "title")
		// userArea := gjson.Get(tempUserInfo, "area")
		//	userAge := ctx.Event.Sender.Age
		// userLevel := gjson.Get(tempUserInfo, "level")
		ctx.SendChain(message.Text("你查询的人为: ", userName, "\n性别:", userSexInfo, "\n最后一次发送信息时间 :", userLastSendTIme, "\n加入时间:", userJoinTime, "\n是否有不友好记录:", userUnfriendly, "\n头衔:", userHonorTitle))
	})
	*/
	engine.OnFullMatch("今日龙王", zero.OnlyGroup).Limit(ctxext.LimitByGroup).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		list := ctx.GetGroupHonorInfo(ctx.Event.GroupID, "talkative")
		temp := list.String()
		id := gjson.Get(temp, "current_talkative.user_id")
		name := ctx.CardOrNickName(id.Int())
		ctx.SendChain(message.Text("今日的龙王是~", name, "哦"))
	})
	// https://github.com/Kittengarten/KittenCore 抄的

	engine.OnFullMatch("随机本群精华消息").SetBlock(true).Limit(ctxext.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		essenceList := ctx.GetThisGroupEssenceMessageList()
		essenceCount := len(essenceList.Array())
		if essenceCount == 0 {
			ctx.Send(fail)
		} else {
			IDx := rand.Intn(essenceCount)
			essenceMessage := essenceList.Array()[IDx]
			var (
				nickname = gjson.Get(essenceMessage.Raw, "sender_nick")
				msID     = gjson.Get(essenceMessage.Raw, "message_id")
			)
			ctx.GetGroupMessageHistory(ctx.Event.GroupID, msID.Int())
			ms := ctx.GetMessage(message.NewMessageIDFromInteger(msID.Int()))
			reportText := message.Text("Lucy抓到了这一条消息~\nUsername: ", nickname)
			report := make(message.Message, len(ms.Elements))
			report = append(report, reportText)
			report = append(report, ms.Elements...)
			deleteme := ctx.Send(report)
			time.Sleep(time.Second * 40)
			ctx.DeleteMessage(deleteme)
		}
	})
	engine.OnRegex(`!catch\s?.*?(.*)$`).Limit(ctxext.LimitByGroup).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		url := ctx.State["regex_matched"].([]string)[1]
		ctx.Send(message.Image("https://ovooa.com/API/Website_snapshot/?url=" + url))
	})
}
