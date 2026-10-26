package model

type VideoItem struct {
	Comment          int64  `json:"comment"`
	Typeid           int    `json:"typeid"`
	Play             int64  `json:"play"`
	Pic              string `json:"pic"`
	Subtitle         string `json:"subtitle"`
	Description      string `json:"description"`
	Copyright        string `json:"copyright"`
	Title            string `json:"title"`
	Review           int    `json:"review"`
	Author           string `json:"author"`
	Mid              int64  `json:"mid"`
	Created          int64  `json:"created"`
	Length           string `json:"length"`
	VideoReview      int64  `json:"video_review"`
	Aid              int64  `json:"aid"`
	Bvid             string `json:"bvid"`
	HideClick        bool   `json:"hide_click"`
	IsPay            int    `json:"is_pay"`
	IsUnionVideo     int    `json:"is_union_video"`
	IsSteinsGate     int    `json:"is_steins_gate"`
	IsLivePlayback   int    `json:"is_live_playback"`
	IsLessonVideo    int    `json:"is_lesson_video"`
	IsLessonFinished int    `json:"is_lesson_finished"`
	LessonUpdateInfo string `json:"lesson_update_info"`
	JumpURL          string `json:"jump_url"`
	Meta             any    `json:"meta"`
	IsAvoided        int    `json:"is_avoided"`
	SeasonID         int    `json:"season_id"`
	Attribute        int    `json:"attribute"`
	IsChargingArc    bool   `json:"is_charging_arc"`
	Vt               int    `json:"vt"`
	EnableVt         int    `json:"enable_vt"`
	VtDisplay        string `json:"vt_display"`
	PlaybackPosition int    `json:"playback_position"`
}

type VideoListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		List struct {
			Vlist []VideoItem `json:"vlist"`
			Slist []any       `json:"slist"`
		} `json:"list"`
		Page struct {
			Pn    int `json:"pn"`
			Ps    int `json:"ps"`
			Count int `json:"count"`
		} `json:"page"`
		EpisodicButton struct {
			Text string `json:"text"`
			URI  string `json:"uri"`
		} `json:"episodic_button"`
		IsRisk      bool `json:"is_risk"`
		GaiaResType int  `json:"gaia_res_type"`
		GaiaData    any  `json:"gaia_data"`
	} `json:"data"`
}

type SearchVideoItem struct {
	Type             string   `json:"type"`
	ID               int64    `json:"id"`
	Author           string   `json:"author"`
	Mid              int64    `json:"mid"`
	Typeid           string   `json:"typeid"`
	Typename         string   `json:"typename"`
	Arcurl           string   `json:"arcurl"`
	Aid              int64    `json:"aid"`
	Bvid             string   `json:"bvid"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	Pic              string   `json:"pic"`
	Play             int64    `json:"play"`
	VideoReview      int64    `json:"video_review"`
	Favorites        int64    `json:"favorites"`
	Tag              string   `json:"tag"`
	Review           int64    `json:"review"`
	Pubdate          int64    `json:"pubdate"`
	Senddate         int64    `json:"senddate"`
	Duration         string   `json:"duration"`
	Badgepay         bool     `json:"badgepay"`
	HitColumns       []string `json:"hit_columns"`
	ViewType         string   `json:"view_type"`
	IsPay            int      `json:"is_pay"`
	IsUnionVideo     int      `json:"is_union_video"`
	RecTags          any      `json:"rec_tags"`
	NewRecTags       []any    `json:"new_rec_tags,omitempty"`
	Like             int64    `json:"like"`
	Upic             string   `json:"upic"`
	Corner           string   `json:"corner"`
	Cover            string   `json:"cover"`
	Desc             string   `json:"desc"`
	Url              string   `json:"url"`
	RecReason        string   `json:"rec_reason"`
	Danmaku          int64    `json:"danmaku"`
	BizData          any      `json:"biz_data"`
	IsChargeVideo    int      `json:"is_charge_video"`
	Vt               int      `json:"vt"`
	EnableVt         int      `json:"enable_vt"`
	VtDisplay        string   `json:"vt_display"`
	Subtitle         string   `json:"subtitle"`
	EpisodeCountText string   `json:"episode_count_text"`
	ReleaseStatus    int      `json:"release_status"`
	IsIntervene      int      `json:"is_intervene"`
	Area             int      `json:"area"`
	Style            int      `json:"style"`
	CateName         string   `json:"cate_name"`
	IsLiveRoomInline int      `json:"is_live_room_inline"`
	LiveStatus       int      `json:"live_status"`
	LiveTime         string   `json:"live_time"`
	Online           int      `json:"online"`
	RankIndex        int      `json:"rank_index"`
	RankOffset       int      `json:"rank_offset"`
	Roomid           int64    `json:"roomid"`
	ShortID          int64    `json:"short_id"`
	SpreadID         int64    `json:"spread_id"`
	Tags             string   `json:"tags"`
	Uface            string   `json:"uface"`
	Uid              int64    `json:"uid"`
	Uname            string   `json:"uname"`
	UserCover        string   `json:"user_cover"`
	ParentAreaID     int      `json:"parent_area_id"`
	ParentAreaName   string   `json:"parent_area_name"`
	WatchedShow      any      `json:"watched_show"`
	CnyFlag          int      `json:"cny_flag"`
	CardStatus       int      `json:"card_status"`
}

type SearchResp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Ttl     int        `json:"ttl"`
	Data    SearchData `json:"data"`
}

type SearchData struct {
	Seid           string            `json:"seid"`
	Page           int               `json:"page"`
	Pagesize       int               `json:"pagesize"`
	NumResults     int               `json:"numResults"`
	NumPages       int               `json:"numPages"`
	SuggestKeyword string            `json:"suggest_keyword"`
	RqtType        string            `json:"rqt_type"`
	ExpList        any               `json:"exp_list"`
	IsHitWebInf    bool              `json:"is_hit_web_inf"`
	EggHit         int               `json:"egg_hit"`
	Result         []SearchVideoItem `json:"result"`
	ShowColumn     int               `json:"show_column"`
	InBlackKey     int               `json:"in_black_key"`
	InWhiteKey     int               `json:"in_white_key"`
	Next           int               `json:"next"`
}
