package nixpacks

import "encoding/json"

type DetectOutput struct {
	Response   []byte
	Providers  []string          `json:"providers"`
	BuildImage string            `json:"buildImage"`
	Variables  map[string]string `json:"variables"`
	Phases     Phases            `json:"phases"`
	Start      Start             `json:"start"`
}

func (b *DetectOutput) Parse() error {
	return json.Unmarshal(b.Response, b)
}
