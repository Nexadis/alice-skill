package api

var Version = "1.0"

type Message string

var CanNothing Message = "Извините, я пока ничего не умею"

type Response struct {
	Text Message `json:"text"`
}

type API struct {
	Response *Response `json:"response,omitempty"`
	Version  string    `json:"version"`
}
