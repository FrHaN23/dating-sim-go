package types

type Res struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
