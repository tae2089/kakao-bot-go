package kakao

type Profile struct {
	Nickname string `json:"nickname"`
	ImageURL string `json:"imageUrl,omitempty"`
}
