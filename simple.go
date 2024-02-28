package kakao

type SimpleTextBlock struct {
	SimpleText simpleText `json:"simpleText"`
}

type simpleText struct {
	Text string `json:"text"`
}

func (s *SimpleTextBlock) ToString() {
	// do nothing
}

func NewSimpleTextBlock(text string) *SimpleTextBlock {
	return &SimpleTextBlock{
		SimpleText: simpleText{
			Text: text,
		},
	}
}

type SimpleImageBlock struct {
	SimpleImage simpleImage `json:"simpleImage"`
}
type simpleImage struct {
	ImageURL string `json:"imageUrl"`
	AltText  string `json:"altText"`
}

func (s *SimpleImageBlock) ToString() {
	// do nothing
}

func NewSimpleImageBlock(imageUrl string, altText string) *SimpleImageBlock {
	return &SimpleImageBlock{
		SimpleImage: simpleImage{
			ImageURL: imageUrl,
			AltText:  altText,
		},
	}
}
