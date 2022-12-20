## `Bilibili` 直播间信息流处理内核

`golang`的b站信息流处理

### 介绍

`go_bilichat` 是一个用于处理 `Bilibili` 直播间信息流的内核，可以用于开发自己的 `Bilibili` 直播间信息流处理程序。

b站直播间信息流以`Websocket`传输并加密, 含有几十个不同的命令, 本项目对其进行了解析, 并提供了一些简单的处理方法, 以便于开发者快速开发自己的程序。

### 快速使用

一个简单的使用示例:

```go
package main

import (
	"fmt"
	"github.com/FishZe/go_bilichat_core"
	"github.com/FishZe/go_bilichat_core/Handler"
	"time"
)

func main() {
	// 创建一个新的消息处理器
	h := go_bilichat_core.GetNewHandler()
	// 将直播间号为15152878弹幕消息交由PrintDanmu处理
	h.New(go_bilichat_core.CmdDanmuMsg, 15152878, PrintDanmuMsg)
	// 将直播间号为15152878礼物消息交由PrintGift处理
	h.New(go_bilichat_core.CmdDanmuMsg, 48743, PrintDanmuMsg)
	// 将直播间绑定到消息处理器
	h.Binding(go_bilichat_core.LiveRoom{RoomId: 15152878})
	h.Binding(go_bilichat_core.LiveRoom{RoomId: 48743})
	// 开始处理消息
	h.Run()
	for {
		time.Sleep(30 * time.Second)
	}
}

func PrintDanmuMsg(event Handler.MsgEvent) {
	fmt.Printf("[%v] %v: %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Content)
}
```
**关于为什么在处理绑定函数时, 多一个直播间号的参数, 因为考虑到可能会有根据不同的直播间分发处理消息的需求**


#### 所有的消息类型:
这些常量请填入`go_bilichat_core.GetNewHandler().New()`的第一个参数
```go
常量名                            原始命令
CmdDanmuMsg                     "DANMU_MSG"
CmdSuperChatMessage             "SUPER_CHAT_MESSAGE"
CmdSuperChatMessageJpn          "SUPER_CHAT_MESSAGE_JPN"
CmdWatchedChange                "WATCHED_CHANGE"
CmdSendGift                     "SEND_GIFT"
CmdOnlineRankCount              "ONLINE_RANK_COUNT"
CmdOnlineRankV2                 "ONLINE_RANK_V2"
CmdOnlineRankTop3               "ONLINE_RANK_TOP3"
CmdLikeInfoV3Click              "LIKE_INFO_V3_CLICK"
CmdInteractWord                 "INTERACT_WORD"
CmdStopLiveRoomList             "STOP_LIVE_ROOM_LIST"
CmdLikeInfoV3Update             "LIKE_INFO_V3_UPDATE"
CmdHotRankChange                "HOT_RANK_CHANGED"
CmdNoticeMsg                    "NOTICE_MSG"
CmdRoomRealTimeMessageUpdate    "ROOM_REAL_TIME_MESSAGE_UPDATE"
CmdWidgetBanner                 "WIDGET_BANNER"
CmdHotRankChangedV2             "HOT_RANK_CHANGED_V2"
CmdGuardHonorThousand           "GUARD_HONOR_THOUSAND"
CmdLive                         "LIVE"
CmdRoomChange                   "ROOM_CHANGE"
CmdRoomBlockMsg                 "ROOM_BLOCK_MSG"
CmdFullScreenSpecialEffect      "FULL_SCREEN_SPECIAL_EFFECT"
CmdCommonNoticeDanmaku          "COMMON_NOTICE_DANMAKU"
CmdTradingScore                 "TRADING_SCORE"
CmdPreparing                    "PREPARING"
CmdGuardBuy                     "GUARD_BUY"
CmdGiftStarProcess              "GIFT_STAR_PROCESS"
CmdRoomSkinMsg                  "ROOM_SKIN_MSG"
CmdEnterEffect                  "ENTER_EFFECT"

```

由于我也不是很明白b站的命令, 所以这里只是列出了我知道的命令, 如果有人知道更多的命令, 请在issue中提出, 我会及时更新。

相信可以通过直译看懂这些命令...


### 消息处理函数

消息处理函数的原型为:

```go
func someFunc(event MsgEvent) {}
```

其中`MsgEvent`的定义为:

```go
type MsgEvent struct {
	//原始命令
	Cmd    string
	//直播间号
    RoomId int
    // 以下为不同的消息类型
    DanMuMsg DanMuMsg
	SuperChatMessage SuperChatMessage
    // 下同, 可参考上方的消息类型, 取消Cmd即为结构体名称...
}
```

### 消息结构体
这是弹幕消息的例子
```go
type DanMuMsg struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Sender struct {
			Uid    int64
			Name   string
			RoomId int64
		}
		Medal                    FansMedal
		Content                  string
		SendTimeStamp            int
		SendMillionTimeStamp     int64
		SenderEnterRoomTimeStamp int
	}
}
```
结构体内容, 请参考`go_bilichat_core/Handler/template.go`中的定义
因为30个实在是太多了, 所以我就不一一列出来了...

## 在弹幕消息中, 我还没有完全明白每个参数的含义, 强烈欢迎各位PR完善

### 抽象画

```text
+============================+   +=================+
|        LiveRoom 1          |   |                 |
| goroutine heart beat       |   |                 → → goroutine your function
| goroutine receive msg → → ↓|   |                 |
| goroutine msg handler ← ← ←|   |                 → → goroutine your function
+================|↓|=========+   _                 |
                  → → → → → → → →    goroutine     → → goroutine your function
+============================+   _                 |
|        LiveRoom 1          |   |  msg distribute → → goroutine your function
| goroutine heart beat       |   |                 |
| goroutine receive msg → → ↓|   |                 → → goroutine your function
| goroutine msg handler ← ← ←|   |                 |
+================|↓|=========+   _                 → → goroutine your function
                  → → → → → → → →                  |
                                 _                 → → goroutine your function
                                 +======|↑|========+
                                         ↑
                                         ↑
                                +=======|↑|========+
                                |                  | 
                                |  main goroutine  |
                                |                  |
                                +==================+
```
