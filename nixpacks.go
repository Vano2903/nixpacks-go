package nixpacks

import "os/exec"

const (
	PlanCommand   = "plan"
	BuildCommand  = "build"
	DetectCommand = "detect"
)

type (
	Nixpacks struct {
		commandPath string
	}

	NixpacksOptions struct {
		//Path to the nixpacks binary
		CommandPath string
	}
)

func NewNixpacks(opt ...NixpacksOptions) (*Nixpacks, error) {
	if len(opt) > 0 {
		n := &Nixpacks{}
		if opt[0].CommandPath != "" {
			n.commandPath = opt[0].CommandPath
		} else {
			foundPath, err := exec.LookPath("nixpacks")
			if err != nil {
				return &Nixpacks{}, err
			}
			n.commandPath = foundPath
		}
		return n, nil
	}

	path, err := exec.LookPath("nixpacks")
	if err != nil {
		return &Nixpacks{}, err
	}
	return &Nixpacks{
		commandPath: path,
	}, nil
}
