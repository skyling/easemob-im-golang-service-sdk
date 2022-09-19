package types

// doc: https://docs-im.easemob.com/push/apppush/pushkv

type PushSingleReq struct {
	Targets     []string    `json:"targets"`               // 推送目标用户 ID。最大可传 100 个。
	Async       bool        `json:"async,omitempty"`       // 是否异步推送
	Strategy    int         `json:"strategy,omitempty"`    // 推送策略
	PushMessage interface{} `json:"pushMessage,omitempty"` // 推送消息
}

type PushDataBaseMessage struct {
	Title    string           `json:"title"`              // 推送显示标题，默认为“您有一条新消息”。该字段长度不能超过 32 字符（一个汉字等于两个字符）。
	SubTitle string           `json:"subTitle,omitempty"` // 推送显示子标题。该字段长度不能超过 10 字符。
	Content  string           `json:"content"`            // 推送显示内容，默认为“请及时查看”。该字段长度不能超过 100 字符（一个汉字等于两个字符）。
	Ext      interface{}      `json:"ext,omitempty"`      // 推送自定义扩展键值。  object
	Config   *PushDataConfig  `json:"config,omitempty"`   // 推送自定义特性
	EaseMob  *PushDataEaseMob `json:"easeMob,omitempty"`  // easemob 推送特性
	APNs     *PushDataAPNs    `json:"apns,omitempty"`     // APNs 推送特性
	FCM      *PushDataFCM     `json:"fcm,omitempty"`      // FCM 推送特性
	XiaoMi   *PushDataXiaoMi  `json:"xiaomi,omitempty"`   // 小米推送特性。
	Vivo     *PushDataVivo    `json:"vivo,omitempty"`     // Vivo 推送特性
	Oppo     *PushDataOppo    `json:"oppo,omitempty"`     // Oppo 推送特性
	Meizu    *PushDataMeizu   `json:"meizu,omitempty"`    // 魅族推送特性
	Huawei   *PushDataHuawei  `json:"huawei,omitempty"`   // 华为推送特性
}

type PushDataBudge struct {
	AddNum   int    `json:"addNum,omitempty"` // 推送角标，自增。
	SetNum   int    `json:"setNum,omitempty"` // 推送角标，覆盖。
	Activity string `json:"activity,omitempty"`
}

type PushDataClickAction struct {
	Url      string `json:"url,omitempty"`      // 推送点击，打开 URL。
	Action   string `json:"action,omitempty"`   // 推送点击，打开 action。
	Activity string `json:"activity,omitempty"` // 推送点击，activity。
}
type PushDataConfig struct {
	ClickAction *PushDataClickAction `json:"clickAction,omitempty"` // 推送点击行为
	Badge       *PushDataBudge       `json:"badge,omitempty"`       // 推送角标
	TimeToLive  int                  // 离线消息保留时长，单位为秒（s），默认为 86,400 秒。
}

type PushDataOperation struct {
	Type       int    `json:"type,omitempty"` //  点击类型。- （默认）0：启动应用；- 1：打开 url；- 2：打开应用指定页
	OpenURL    string `json:"openUrl,omitempty"`
	OpenAction string `json:"openAction,omitempty"`
}
type PushDataEaseMob struct {
	Title            string             `json:"title,omitempty"`
	Content          string             `json:"content,omitempty"`
	SubTitle         string             `json:"subTitle,omitempty"`
	IconUrl          string             `json:"iconUrl,omitempty"`
	NeedNotification bool               `json:"needNotification,omitempty"`
	Badge            *PushDataBudge     `json:"badge,omitempty"`
	Operation        *PushDataOperation `json:"operation,omitempty"`
	ChannelId        string             `json:"channelId,omitempty"`
	ChannelName      string             `json:"channelName,omitempty"`
	ChannelLevel     int                `json:"channelLevel,omitempty"`
	AutoCancel       int                `json:"autoCancel,omitempty"`
	ExpiresTime      int                `json:"expiresTime,omitempty"`
	Sound            int                `json:"sound,omitempty"`
	Vibrate          int                `json:"vibrate,omitempty"`
	Style            int                `json:"style,omitempty"`
	BigTxt           string             `json:"bigTxt,omitempty"`
	BigPicture       string             `json:"bigPicture,omitempty"`
	Id               int                `json:"id,omitempty"`
}

type PushDataAPNs struct {
	InvalidationTime int         `json:"invalidationTime,omitempty"`
	Priority         int         `json:"priority,omitempty"`
	PushType         string      `json:"pushType,omitempty"`
	CollapseId       string      `json:"collapseId,omitempty"`
	ApnsId           string      `json:"apnsId,omitempty"`
	Badge            int         `json:"badge,omitempty"`
	Sound            string      `json:"sound,omitempty"`
	MutableContent   bool        `json:"mutableContent,omitempty"`
	ContentAvailable bool        `json:"contentAvailable,omitempty"`
	CategoryName     string      `json:"categoryName,omitempty"`
	ThreadId         string      `json:"threadId,omitempty"`
	Title            string      `json:"title,omitempty"`
	SubTitle         string      `json:"subTitle,omitempty"`
	Content          string      `json:"content,omitempty"`
	TitleLocKey      string      `json:"titleLocKey,omitempty"`
	TitleLocArgs     []string    `json:"titleLocArgs,omitempty"`
	SubTitleLocKey   string      `json:"subTitleLocKey,omitempty"`
	SubTitleLocArgs  []string    `json:"subTitleLocArgs,omitempty"`
	BodyLocKey       string      `json:"bodyLocKey,omitempty"`
	BodyLocArgs      []string    `json:"bodyLocArgs,omitempty"`
	Ext              interface{} `json:"ext,omitempty"`
	LaunchImage      string      `json:"launchImage,omitempty"`
}

type PushDataFCM struct {
	Condition             string                   `json:"condition,omitempty"`
	CollapseKey           string                   `json:"collapseKey,omitempty"`
	Priority              string                   `json:"priority,omitempty"`
	TimeToLive            string                   `json:"timeToLive,omitempty"`
	DryRun                bool                     `json:"dryRun,omitempty"`
	RestrictedPackageName string                   `json:"restrictedPackageName,omitempty"`
	Data                  interface{}              `json:"data,omitempty"`
	Notification          *PushDataFCMNotification `json:"notification,omitempty"`
}

type PushDataFCMNotification struct {
	Title            string   `json:"title,omitempty"`
	Body             string   `json:"body,omitempty"`
	AndroidChannelId string   `json:"androidChannelId,omitempty"`
	Icon             string   `json:"icon,omitempty"`
	Sound            string   `json:"sound,omitempty"`
	Tag              string   `json:"tag,omitempty"`
	Color            string   `json:"color,omitempty"`
	ClickAction      string   `json:"clickAction,omitempty"`
	TitleLocKey      string   `json:"titleLocKey,omitempty"`
	TitleLocArgs     []string `json:"titleLocArgs,omitempty"`
	BodyLocKey       string   `json:"bodyLocKey,omitempty"`
	BodyLocArgs      []string `json:"bodyLocArgs,omitempty"`
}

type PushDataXiaoMi struct {
	AppVersion             string      `json:"appVersion,omitempty"`
	AppVersionNotIn        string      `json:"appVersionNotIn,omitempty"`
	CallbackParam          string      `json:"callbackParam,omitempty"`
	CallbackType           int         `json:"callbackType,omitempty"`
	CallbackUrl            string      `json:"callbackUrl,omitempty"`
	ChannelId              string      `json:"channelId,omitempty"`
	Connpt                 string      `json:"connpt,omitempty"`
	Description            string      `json:"description,omitempty"`
	FlowControl            int         `json:"flowControl,omitempty"`
	IntentUri              string      `json:"intentUri,omitempty"`
	Jobkey                 string      `json:"jobkey,omitempty"`
	Locale                 string      `json:"locale,omitempty"`
	LocaleNotIn            string      `json:"localeNotIn,omitempty"`
	Model                  string      `json:"model,omitempty"`
	ModelNotIn             string      `json:"modelNotIn,omitempty"`
	NotifyEffect           string      `json:"notifyEffect,omitempty"`
	NotifyForeground       string      `json:"notifyForeground,omitempty"`
	NotifyId               int         `json:"notifyId,omitempty"`
	NotifyType             int         `json:"notifyType,omitempty"`
	OnlySendOnce           string      `json:"onlySendOnce,omitempty"`
	PassThrough            int         `json:"passThrough,omitempty"`
	Payload                interface{} `json:"payload,omitempty"`
	RestrictedPackageNames string      `json:"restrictedPackageNames,omitempty"`
	Sound                  string      `json:"sound,omitempty"`
	Ticker                 string      `json:"ticker,omitempty"`
	TimeToLive             int         `json:"timeToLive,omitempty"`
	TimeToSend             int         `json:"timeToSend,omitempty"`
	Title                  string      `json:"title,omitempty"`
	WebUri                 string      `json:"webUri,omitempty"`
}

type PushDataVivo struct {
	Title           string      `json:"title,omitempty"`
	Content         string      `json:"content,omitempty"`
	NotifyType      int         `json:"notifyType,omitempty"`
	TimeToLive      int         `json:"timeToLive,omitempty"`
	SkipType        int         `json:"skipType,omitempty"`
	SkipContent     string      `json:"skipContent,omitempty"`
	NetworkType     int         `json:"networkType,omitempty"`
	Classification  int         `json:"classification,omitempty"`
	ClientCustomMap interface{} `json:"clientCustomMap,omitempty"`
	RequestId       string      `json:"requestId,omitempty"`
	PushMode        int         `json:"pushMode,omitempty"`
	Extra           struct {
		Callback      string `json:"callback,omitempty"`
		CallbackParam string `json:"callbackParam,omitempty"`
	} `json:"extra,omitempty"`
}

type PushDataOppo struct {
	Title               string      `json:"title,omitempty"`
	SubTitle            string      `json:"subTitle,omitempty"`
	Content             string      `json:"content,omitempty"`
	ChannelId           string      `json:"channelId,omitempty"`
	AppMessageId        string      `json:"appMessageId,omitempty"`
	Style               int         `json:"style,omitempty"`
	ClickActionType     int         `json:"clickActionType,omitempty"`
	ClickActionActivity string      `json:"clickActionActivity,omitempty"`
	ClickActionUrl      string      `json:"clickActionUrl,omitempty"`
	ActionParameters    interface{} `json:"actionParameters,omitempty"`
	OffLine             bool        `json:"offLine,omitempty"`
	OffLineTtl          int         `json:"offLineTtl,omitempty"`
	TimeZone            string      `json:"timeZone,omitempty"`
	CallBackUrl         string      `json:"callBackUrl,omitempty"`
	CallBackParameter   string      `json:"callBackParameter,omitempty"`
	ShowTtl             int         `json:"showTtl,omitempty"`
	NotifyId            int         `json:"notifyId,omitempty"`
}

type PushDataMeizu struct {
	Title               string      `json:"title,omitempty"`
	Content             string      `json:"content,omitempty"`
	NoticeExpandType    int         `json:"noticeExpandType,omitempty"`
	NoticeExpandContent string      `json:"noticeExpandContent,omitempty"`
	ClickType           int         `json:"clickType,omitempty"`
	Url                 string      `json:"url,omitempty"`
	Parameters          interface{} `json:"parameters,omitempty"`
	Activity            string      `json:"activity,omitempty"`
	CustomAttribute     string      `json:"customAttribute,omitempty"`
	IsOffLine           bool        `json:"isOffLine,omitempty"`
	ValidTime           int         `json:"validTime,omitempty"`
	IsSuspend           bool        `json:"isSuspend,omitempty"`
	IsClearNoticeBar    bool        `json:"isClearNoticeBar,omitempty"`
	IsFixDisplay        bool        `json:"isFixDisplay,omitempty"`
	FixStartDisplayDate string      `json:"fixStartDisplayDate,omitempty"`
	FixEndDisplayDate   string      `json:"fixEndDisplayDate,omitempty"`
	Vibrate             bool        `json:"vibrate,omitempty"`
	Lights              bool        `json:"lights,omitempty"`
	Sound               bool        `json:"sound,omitempty"`
	NotifyKey           string      `json:"notifyKey,omitempty"`
	Callback            string      `json:"callback,omitempty"`
	CallbackParam       string      `json:"callbackParam,omitempty"`
	CallbackType        string      `json:"callbackType,omitempty"`
}

type PushDataHuawei struct {
	ValidateOnly bool                   `json:"validateOnly,omitempty"`
	Message      *PushDataHuaweiMessage `json:"message,omitempty"`
	Review       interface{}            `json:"review,omitempty"`
}

type PushDataHuaweiMessageButton struct {
	Name       string `json:"name,omitempty"`
	ActionType int    `json:"actionType,omitempty"`
	IntentType int    `json:"intentType,omitempty"`
	Intent     string `json:"intent,omitempty"`
	Data       string `json:"data,omitempty"`
}

type PushDataHuaweiMessage struct {
	Data         interface{} `json:"data,omitempty"`
	Notification struct {
		Title string `json:"title,omitempty"`
		Body  string `json:"body,omitempty"`
		Image string `json:"image,omitempty"`
	} `json:"notification,omitempty"`
	Android struct {
		CollapseKey       int    `json:"collapseKey,omitempty"`
		Urgency           string `json:"urgency,omitempty"`
		Category          string `json:"category,omitempty"`
		Ttl               string `json:"ttl,omitempty"`
		BiTag             string `json:"biTag,omitempty"`
		FastAppTargetType int    `json:"fastAppTargetType,omitempty"`
		Data              string `json:"data,omitempty"`
		Notification      struct {
			Title        string `json:"title,omitempty"`
			Body         string `json:"body,omitempty"`
			Icon         string `json:"icon,omitempty"`
			Color        string `json:"color,omitempty"`
			Sound        string `json:"sound,omitempty"`
			DefaultSound bool   `json:"defaultSound,omitempty"`
			Tag          string `json:"tag,omitempty"`
			ClickAction  struct {
				Type   int    `json:"type,omitempty"`
				Intent string `json:"intent,omitempty"`
				Url    string `json:"url,omitempty"`
				Action string `json:"action,omitempty"`
			} `json:"clickAction,omitempty"`
			BodyLocKey    string   `json:"bodyLocKey,omitempty"`
			BodyLocArgs   []string `json:"bodyLocArgs,omitempty"`
			TitleLocKey   string   `json:"titleLocKey,omitempty"`
			TitleLocArgs  []string `json:"titleLocArgs,omitempty"`
			MultiLangKey  []string `json:"multiLangKey,omitempty"`
			ChannelId     string   `json:"channelId,omitempty"`
			NotifySummary string   `json:"notifySummary,omitempty"`
			Image         string   `json:"image,omitempty"`
			Style         int      `json:"style,omitempty"`
			BigTitle      string   `json:"bigTitle,omitempty"`
			BigBody       string   `json:"bigBody,omitempty"`
			AutoClear     int      `json:"autoClear,omitempty"`
			NotifyId      int      `json:"notifyId,omitempty"`
			Group         string   `json:"group,omitempty"`
			Badge         struct {
				AddNum     int    `json:"addNum,omitempty"`
				BadgeClass string `json:"badgeClass,omitempty"`
				SetNum     int    `json:"setNum,omitempty"`
			} `json:"badge,omitempty"`
			AutoCancel        bool   `json:"autoCancel,omitempty"`
			When              string `json:"when,omitempty"`
			Importance        string `json:"importance,omitempty"`
			UseDefaultVibrate bool   `json:"useDefaultVibrate,omitempty"`
			UseDefaultLight   bool   `json:"useDefaultLight,omitempty"`
			VibrateConfig     string `json:"vibrateConfig,omitempty"`
			Visibility        string `json:"visibility,omitempty"`
			LightSettings     struct {
				Color struct {
					Alpha float64 `json:"alpha,omitempty"`
					Red   float64 `json:"red,omitempty"`
					Green float64 `json:"green,omitempty"`
					Blue  float64 `json:"blue,omitempty"`
				} `json:"color,omitempty"`
				LightOnDuration  string `json:"lightOnDuration,omitempty"`
				LightOffDuration string `json:"lightOffDuration,omitempty"`
			} `json:"lightSettings,omitempty"`
			ForegroundShow bool                           `json:"foregroundShow,omitempty"`
			InboxContent   []string                       `json:"inboxContent,omitempty"`
			Buttons        []*PushDataHuaweiMessageButton `json:"buttons,omitempty"`
		} `json:"notification,omitempty"`
	}
}

// PushSingleResp 推送响应
type PushSingleResp struct {
	Timestamp int64 `json:"timestamp,omitempty"`
	Data      []struct {
		Id         string `json:"id,omitempty"`
		PushStatus string `json:"pushStatus,omitempty"`
		Desc       string `json:"desc,omitempty"`
	} `json:"data,omitempty"`
	Duration int `json:"duration,omitempty"`
}

type PushMessageResp struct {
	Timestamp int64 `json:"timestamp"`
	Data      int64 `json:"data"`
	Duration  int   `json:"duration"`
}

type GetPushMessageResp struct {
	Timestamp int64 `json:"timestamp"`
	Data      struct {
		PushDataBaseMessage
		Appkey    string `json:"appkey"`
		Timestamp int64  `json:"timestamp"`
	} `json:"data"`
	Duration int `json:"duration"`
}

type PushTaskBroadcastReq struct {
	Type      string `json:"type"`      // IM
	Strategy  int    `json:"strategy"`  // 推送策略
	PushFunc  string `json:"pushFunc"`  // 推送方式（ALL）目前只支持全部推送（后续会增加组推送等）
	PushMsgId int64  `json:"pushMsgId"` // 消息ID
}

type PushTaskReq struct {
	Type        string      `json:"type"`                  // IM
	Strategy    int         `json:"strategy"`              // 推送策略
	PushFunc    string      `json:"pushFunc"`              // 推送方式（ALL）目前只支持全部推送（后续会增加组推送等）
	PushMessage interface{} `json:"pushMessage,omitempty"` // 推送消息
}

type PushListReq struct {
	Type        string      `json:"type"`                  // IM
	Strategy    int         `json:"strategy"`              // 推送策略
	PushFunc    string      `json:"pushFunc"`              // 推送类型：LABEL：标签推送。
	PushMessage interface{} `json:"pushMessage,omitempty"` // 推送消息
}

type PushListResp struct {
	Timestamp int64 `json:"timestamp"`
	Data      struct {
		TaskId int64 `json:"taskId"`
	} `json:"data"`
	Duration int `json:"duration"`
}
