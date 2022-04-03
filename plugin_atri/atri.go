package atri

import (
	"math/rand"
	"time"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	control "github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/process"

	"github.com/FloatTech/ZeroBot-Plugin/order"
)

const (
	// 服务名
	servicename = "atri"
	// ATRI 表情的 codechina 镜像
	res = "https://gitcode.net/u011570312/zbpdata/-/raw/main/Atri/"

)

func init() { // 插件主体
	engine := control.Register(servicename, order.PrioAtri, &control.Options{
		DisableOnDefault: false,
		Help: "- Hana醒醒\n- Hana睡吧\n- 萝卜子\n- 喜欢 | 爱你 | 爱 | suki | daisuki | すき | 好き | 贴贴 | 老婆 | 亲一个 | mua\n" +
			"- 草你妈 | 操你妈 | 脑瘫 | 废柴 | fw | 废物 | 战斗 | 爪巴 | sb | SB | 傻B\n- 早安 | 早哇 | 早上好 | ohayo | 哦哈哟 | お早う | 早好 | 早 | 早早早\n" +
			"- 中午好 | 午安 | 午好\n- 晚安 | oyasuminasai | おやすみなさい | 晚好 | 晚上好\n- 高性能 | 太棒了 | すごい | sugoi | 斯国一 | よかった\n" +
			"- 没事 | 没关系 | 大丈夫 | 还好 | 不要紧 | 没出大问题 | 没伤到哪\n- 好吗 | 是吗 | 行不行 | 能不能 | 可不可以\n- 啊这\n- 我好了\n- ？ | ? | ¿\n" +
			"- 离谱\n- 答应我",
	})

	zero.OnFullMatch("Hana醒醒", zero.AdminPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			c, ok := control.Lookup(servicename)
			if ok && !c.IsEnabledIn(ctx.Event.GroupID) {
				c.Enable(ctx.Event.GroupID)
				process.SleepAbout1sTo2s()
				ctx.SendChain(message.Text("嗯呜呜……Zzz……？"))
			}
		})
	engine.OnFullMatch("Hana睡吧", zero.AdminPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			c, ok := control.Lookup(servicename)
			if ok && c.IsEnabledIn(ctx.Event.GroupID) {
				c.Disable(ctx.Event.GroupID)
				process.SleepAbout1sTo2s()
				ctx.SendChain(message.Text("Zzz……Zzz……"))
			}
		})

	engine.OnFullMatchGroup([]string{"喜欢", "爱你", "爱", "suki", "daisuki", "すき", "好き", "贴贴", "老婆", "亲一个", "mua"}, atriSleep, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randImage("SUKI.jpg"))
		})

	engine.OnKeywordGroup([]string{"草你妈", "操你妈", "脑瘫", "废柴", "fw", "five", "废物", "战斗", "爪巴", "sb", "SB", "傻B"}, atriSleep, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randImage("FN.jpg", "WQ.jpg", "WQ1.jpg"))
		})

	engine.OnFullMatchGroup([]string{"早安", "早哇", "早上好", "ohayo", "哦哈哟", "お早う", "早好", "早", "早早早"},zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			process.SleepAbout1sTo2s()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...好涩哦..zzz....",
					"别...不要..zzz..那..zzz..",
					"嘻嘻..zzz..呐~..zzzz..",
					"...zzz....哧溜哧溜....",
				))
			case now >= 6 && now < 9:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"啊......早上好...(哈欠)",
					"唔......吧唧...早上...哈啊啊~~~\n早上好......",
					"早上好......",
					"早上好呜......呼啊啊~~~~",
					"吧唧吧唧......怎么了...已经早上了么...",
					"早上好！",
					"......看起来像是傍晚，其实已经早上了吗？",
					"早上好......欸~~~脸好近呢",
					"起的很早(＾Ｕ＾)ノ~ＹＯ~ 一定是个勤奋的孩子w",
					"早安吖~新的一天要继续加油哦w",
					"早安~祝你的坏心情被Hana带走~~~(*/ω＼*)",
				))
			case now >= 9 && now < 18:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"哼！这个点还早啥，昨晚干啥去了！？",
					"熬夜了对吧熬夜了对吧熬夜了对吧？？？！",
					"是不是熬夜是不是熬夜是不是熬夜？！",
					"Hana酱提醒你~不要熬夜哟w~ 否则就和架子一样被打死（）",
					"哼！ 这个点起床 一定是熬夜了！(｀･∪･´)",
				))
			case now >= 18 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"早个啥？哼唧！我都准备洗洗睡了！",
					"不是...你看看几点了，哼！",
					"晚上好哇",
					"是时区问题嘛~ 还是上夜班呢~ 总之记得对自己好一点哦w",
				))
			}
		})

	engine.OnFullMatchGroup([]string{"中午好", "午安", "午好"},zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			if now > 11 && now < 15 { // 中午
				process.SleepAbout1sTo2s()
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，Hana会陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"睡你午觉去！哼唧！！",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"晚安", "oyasuminasai", "おやすみなさい", "晚好", "晚上好"},zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			process.SleepAbout1sTo2s()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...好涩哦..zzz....",
					"别...不要..zzz..那..zzz..",
					"嘻嘻..zzz..呐~..zzzz..",
					"...zzz....哧溜哧溜....",
				))
			case now >= 6 && now < 11:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"你可猝死算了吧！",
					"？啊这",
					"亲，这边建议赶快去睡觉呢~~~",
					"不可忍不可忍不可忍！！为何这还不猝死！！",
					"? 你知道现在几点了嘛 \n还不快去睡觉~",
				))
			case now >= 11 && now < 15:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，Hana会陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"睡你午觉去！哼唧！！",
				))
			case now >= 15 && now < 19:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"难不成？？晚上不想睡觉？？现在休息",
					"就......挺离谱的...现在睡觉",
					"现在还是白天哦，睡觉还太早了",
				))
			case now >= 19 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"......(打瞌睡)",
					"呼...呼...已经睡着了哦~...呼......",
					"......我、我会在这守着你的，请务必好好睡着",
				))
			}
		})
	engine.OnKeywordGroup([]string{"高性能", "太棒了", "すごい", "sugoi", "斯国一", "よかった"}, atriSleep, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randText(
				"当然，我是高性能的嘛~！",
				"小事一桩，我是高性能的嘛",
				"怎么样？还是我比较高性能吧？",
				"哼哼！我果然是高性能的呢！",
				"因为我是高性能的嘛！嗯哼！",
				"因为我是高性能的呢！",
				"哎呀~，我可真是太高性能了",
				"正是，因为我是高性能的",
				"是的。我是高性能的嘛♪",
				"毕竟我可是高性能的！",
				"嘿嘿，我的高性能发挥出来啦♪",
				"我果然是很高性能的机器人吧！",
				"是吧！谁叫我这么高性能呢！哼哼！",
				"交给我吧，有高性能的我陪着呢",
				"呣......我的高性能，毫无遗憾地施展出来了......",
			))
		})

	engine.OnKeywordGroup([]string{"没事", "没关系", "大丈夫", "还好", "不要紧", "没出大问题", "没伤到哪"}, atriSleep, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randText(
				"当然，我是高性能的嘛~！",
				"没事没事，因为我是高性能的嘛！嗯哼！",
				"没事的，因为我是高性能的呢！",
				"正是，因为我是高性能的",
				"是的。我是高性能的嘛♪",
				"毕竟我可是高性能的！",
				"那种程度的事不算什么的。\n别看我这样，我可是高性能的",
				"没问题的，我可是高性能的",
			))
		})

	engine.OnKeywordGroup([]string{"我好了"}, atriSleep).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText("不许好！", "憋回去！"))
		})


}
func randText(text ...string) message.MessageSegment {
	return message.Text(text[rand.Intn(len(text))])
}

func randImage(file ...string) message.MessageSegment {
	return message.Image(res + file[rand.Intn(len(file))])
}

func randRecord(file ...string) message.MessageSegment {
	return message.Record(res + file[rand.Intn(len(file))])
}



// atriSleep 凌晨3点到6点，ATRI 在睡觉，不回应任何请求
func atriSleep(ctx *zero.Ctx) bool {
	if now := time.Now().Hour(); now >= 4 && now < 6 {
		return false
	}
	return true
}
