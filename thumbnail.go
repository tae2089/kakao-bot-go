package kakao

type Thumbnail struct {
	ImageURL   string `json:"imageUrl"`
	FixedRatio bool   `json:"fixedRatio,omitempty"`
	Link       Link   `json:"link,omitempty"`
}

// BasicCard 또는 CommerceCard에서 true 설정 시, 버튼이 가로로 배열되며 최대 2개로 제한됩니다.
// 케로셀 내에서는 모든 이미지가 정사각형 (1:1) 혹은 모든 이미지가 와이드형 (2:1)으로 통일되어야 합니다.
// true: 이미지 영역을 1:1 비율로 두고 이미지의 원본 비율을 유지합니다. 이미지가 없는 영역은 흰색으로 노출합니다.
// false: 이미지 영역을 2:1 비율로 두고 이미지의 가운데를 크롭하여 노출합니다.

func NewThumnail(imageUrl string, fixedRatio bool, link Link) *Thumbnail {
	return &Thumbnail{
		ImageURL:   imageUrl,
		FixedRatio: fixedRatio,
		Link:       link,
	}
}
