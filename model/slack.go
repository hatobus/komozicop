package model

type ChallengeReq struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"url_verification"`
}

type ChallengeRes struct {
	Challenge string `json:"challenge"`
}

type UserMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	Ts      string `json:"ts"`
}

type PostMessage struct {
	Token     string `json:"token"`
	Channel   string `json:"channel"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
	ThreadTs  string `json:"thread_ts"`
	UserName  string `json:"username"`
}
