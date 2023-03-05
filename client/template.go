package client

import "errors"

const (
	AuthProto      = 1
	HeartBeatProto = 1
	CmdZlibProto   = 2
	CmdBrotliProto = 3
)

const (
	OpError          = 1
	OpHeartBeat      = 2
	OpHeartBeatReply = 3
	OpCmd            = 5
	OpAuth           = 7
	OpAuthReply      = 8
)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
var ErrespCodeNot = errors.New("resp code not 0")
=======
var INF = 0x3f3f3f
=======
>>>>>>> a7a3af0 (Revert "更换 ws 库")
=======
var INF = 0x3f3f3f
>>>>>>> 61a0301 (添加根据延迟排序的功能)
var RespCodeNotError = errors.New("resp code not 0")
>>>>>>> 61a0301 (添加根据延迟排序的功能)
=======
var ErrespCodeNot = errors.New("resp code not 0")
>>>>>>> 92e7fae (更换 ws 库)

type ApiLiveAuth struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Group            string  `json:"group"`
		BusinessID       int     `json:"business_id"`
		RefreshRowFactor float64 `json:"refresh_row_factor"`
		RefreshRate      int     `json:"refresh_rate"`
		MaxDelay         int     `json:"max_delay"`
		Token            string  `json:"token"`
		HostList         []struct {
			Host    string `json:"host"`
			Port    int    `json:"port"`
			WssPort int    `json:"wss_port"`
			WsPort  int    `json:"ws_port"`
		} `json:"host_list"`
	} `json:"data"`
}

type ApiLiveRoomId struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		RoomID          int           `json:"room_id"`
		ShortID         int           `json:"short_id"`
		UID             int           `json:"uid"`
		NeedP2P         int           `json:"need_p2p"`
		IsHidden        bool          `json:"is_hidden"`
		IsLocked        bool          `json:"is_locked"`
		IsPortrait      bool          `json:"is_portrait"`
		LiveStatus      int           `json:"live_status"`
		HiddenTill      int           `json:"hidden_till"`
		LockTill        int           `json:"lock_till"`
		Encrypted       bool          `json:"encrypted"`
		PwdVerified     bool          `json:"pwd_verified"`
		LiveTime        int           `json:"live_time"`
		RoomShield      int           `json:"room_shield"`
		IsSp            int           `json:"is_sp"`
		SpecialType     int           `json:"special_type"`
		PlayURL         interface{}   `json:"play_url"`
		AllSpecialTypes []interface{} `json:"all_special_types"`
	} `json:"data"`
}
type WsHeader struct {
	PackageLen uint32
	HeaderLen  uint16
	ProtoVer   uint16
	OpCode     uint32
	Sequence   uint32
}

type WsAuthBody struct {
	UID      int    `json:"uid"`
	Roomid   int    `json:"roomid"`
	Protover int    `json:"protover"`
	Platform string `json:"platform"`
	Type     int    `json:"type"`
	Key      string `json:"key"`
}

type WsAuthMessage struct {
	WsHeader WsHeader
	Body     WsAuthBody
}

type WsAuthReplyBody struct {
	Code int `json:"code"`
}

type WsAuthReplyMessage struct {
	WsHeader WsHeader
	Body     WsAuthReplyBody
}

type WsHeartBeatMessage struct {
	WsHeader WsHeader
	Body     []byte
}

type WsHeartBeatReply struct {
	WsHeader WsHeader
	Hot      uint32
	Msg      []byte
}

type WsCmdMessage struct {
	WsHeader WsHeader
	Cmd      string
	Body     []byte
}

type WsCmdBrotliMessage struct {
	WsHeader WsHeader
	Body     []WsCmdMessage
}
