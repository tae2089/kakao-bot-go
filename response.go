package kakao

type Response struct {
	Version  string   `json:"version"`
	Template Template `json:"template"`
}

type Template struct {
	Blocks   []Block   `json:"outputs"`
	Replicas []Replica `json:"quickReplies,omitempty"`
}

type Block interface {
	ToString()
}

type Replica struct {
	Label       string `json:"label"`
	Action      string `json:"action"`
	MessageText string `json:"messageText"`
	BlockID     string `json:"blockId"`
	Extra       string `json:"extra"`
}
