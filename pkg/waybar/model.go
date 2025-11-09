package waybar

const (
	MAX_LENGTH = 35
)

type ModuleOutput struct {
	Text  string `json:"text"`
	Class string `json:"class"`
}
