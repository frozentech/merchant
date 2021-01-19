package model

// Response ...
type Response struct {
	Success bool        `json:"success"`
	Error   *Error      `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

// Error ...
type Error struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Records ...
type Records struct {
	Count  int         `json:"int"`
	Page   int         `json:"page"`
	Record interface{} `json:"record"`
}

// SetResult ....
func (me *Response) SetResult(v interface{}) {
	me.Result = v
}
