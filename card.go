package kakao

type TextCardBlock struct {
	TextCard TextCard `json:"textCard"`
}

type BasicCardBlock struct {
	BasicCard BasicCard `json:"basicCard"`
}

type CommerceCardBlock struct {
	CommerceCard CommerceCard `json:"commerceCard"`
}

type BasicCard struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Buttons     []Button  `json:"buttons,omitempty"`
}

type TextCard struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Buttons     []Button `json:"buttons,omitempty"`
}

type CommerceCard struct {
	Title         string      `json:"title,omitempty"`
	Description   string      `json:"description,omitempty"`
	Price         int         `json:"price"`
	Currency      string      `json:"currency,omitempty"`
	Discount      int         `json:"discount,omitempty"`
	DiscountRate  int         `json:"discountRate,omitempty"`
	DiscountPrice int         `json:"discountPrice,omitempty"`
	Thumbnails    []Thumbnail `json:"thumbnails"`
	Profile       Profile     `json:"profile,omitempty"`
	Buttons       []Button    `json:"buttons,omitempty"`
}

func NewTextCardBlock(title string, description string, buttons []Button) *TextCardBlock {
	return &TextCardBlock{
		TextCard: TextCard{
			Title:       title,
			Description: description,
			Buttons:     buttons,
		},
	}
}

func NewBasicCardBlock(title string, description string, thumbnail Thumbnail, buttons []Button) *BasicCardBlock {
	return &BasicCardBlock{
		BasicCard: BasicCard{
			Title:       title,
			Description: description,
			Thumbnail:   thumbnail,
			Buttons:     buttons,
		},
	}
}

func NewCommerceCardBlock(title string, description string, price int, currency string, discount int, discountRate int, discountPrice int, thumbnails []Thumbnail, profile Profile, buttons []Button) *CommerceCardBlock {
	return &CommerceCardBlock{
		CommerceCard: CommerceCard{
			Title:         title,
			Description:   description,
			Price:         price,
			Currency:      currency,
			Discount:      discount,
			DiscountRate:  discountRate,
			DiscountPrice: discountPrice,
			Thumbnails:    thumbnails,
			Profile:       profile,
			Buttons:       buttons,
		},
	}
}

func (t *TextCardBlock) AddButton(button ...Button) {
	t.TextCard.Buttons = append(t.TextCard.Buttons, button...)
}

func (t *TextCardBlock) ToString() {
	// do nothing
}

func (b *BasicCard) AddButton(button ...Button) {
	b.Buttons = append(b.Buttons, button...)
}

func (b *BasicCard) ToString() {
	// do nothing
}

func (c *CommerceCard) AddButton(button ...Button) {
	c.Buttons = append(c.Buttons, button...)
}

func (c *CommerceCard) ToString() {
	// do nothing
}
