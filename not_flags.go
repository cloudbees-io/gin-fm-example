package main

type notFlags struct {
	Message     string `json:"message"`
	FontColor   string `json:"fontColor"`
	FontSize    int    `json:"fontSize"`
	ShowMessage bool   `json:"showMessage"`
}

func NewNotFlags(message, fontColor string, fontSize int, showMessage bool) notFlags {
	return notFlags{
		Message:     message,
		FontColor:   fontColor,
		FontSize:    fontSize,
		ShowMessage: showMessage,
	}
}

func (n notFlags) GetMessage() string {
	return n.Message
}

func (n notFlags) GetFontColor() string {
	return n.FontColor
}

func (n notFlags) GetFontSize() int {
	return n.FontSize
}

func (n notFlags) IsShowMessage() bool {
	return n.ShowMessage
}
