package nixpacks

import "bytes"

type DetectOutput struct {
	Response  []byte
	Providers []string `json:"providers"`
}

func (b *DetectOutput) Parse() error {
	b.Response = bytes.TrimSpace(b.Response)
	if len(b.Response) == 0 {
		return nil
	}
	providers := bytes.Split(b.Response, []byte(", "))
	for _, p := range providers {
		b.Providers = append(b.Providers, string(p))
	}
	return nil
}
