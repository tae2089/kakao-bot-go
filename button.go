package kakao

type Button struct {
	Label       string `json:"label"`
	Action      string `json:"action"`
	WebLink     string `json:"webLinkUrl,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Message     string `json:"messageText,omitempty"`
	Block       string `json:"blockId,omitempty"`
	Extra       string `json:"extra,omitempty"`
}

type ButtonType struct {
	Action       string
	Value        string
	BlockMessage string
}

type ButtonTypeFunc func(value string) ButtonType

func WithWebLink(value string) ButtonType {
	return ButtonType{
		Action: "webLink",
		Value:  value,
	}
}

func WithPhoneNumber(phoneNumber string) ButtonType {
	return ButtonType{
		Action: "phone",
		Value:  phoneNumber,
	}
}

func WithMessage(messageText string) ButtonType {
	return ButtonType{
		Action: "message",
		Value:  messageText,
	}
}

func WithBlock(blockId, message string) ButtonType {
	return ButtonType{
		Action:       "block",
		Value:        blockId,
		BlockMessage: message,
	}
}

func NewButton(label string, f ButtonTypeFunc) *Button {
	var button Button = Button{}
	var v string
	button.Label = label
	buttonType := f(v)
	button.Action = buttonType.Action
	switch buttonType.Action {
	case "webLink":
		button.WebLink = buttonType.Value
	case "phone":
		button.PhoneNumber = buttonType.Value
	case "message":
		button.Message = buttonType.Value
	case "block":
		button.Block = buttonType.Value
		button.Message = buttonType.BlockMessage
	}
	return &button
}
