package nixpacks

import "encoding/json"

type PlanOutput struct {
	Response   []byte
	Providers  []string          `json:"providers"`
	BuildImage string            `json:"buildImage"`
	Variables  map[string]string `json:"variables"`
	Phases     Phases            `json:"phases"`
	Start      Start             `json:"start"`
}

type Phases struct {
	Build   Phase `json:"build"`
	Install Phase `json:"install"`
	Setup   Setup `json:"setup"`
}

type Phase struct {
	DependsOn        []string `json:"dependsOn"`
	Cmds             []string `json:"cmds"`
	CacheDirectories []string `json:"cacheDirectories"`
}

type Setup struct {
	NixPkgs        []string `json:"nixPkgs"`
	NixLibs        []string `json:"nixLibs"`
	AptPkgs        []string `json:"aptPkgs"`
	NixOverlays    []string `json:"nixOverlays"`
	NixpkgsArchive string   `json:"nixpkgsArchive"`
}

type Start struct {
	Cmd      string `json:"cmd"`
	RunImage string `json:"runImage"`
}

func (b *PlanOutput) Parse() error {
	return json.Unmarshal(b.Response, b)
}
