package nixpacks

import (
	"errors"
	"fmt"
)

type PlanOptions struct {
	Path string
	//Command to install language dependencies.
	InstallCommand string
	//Command to build the image,
	//it will overwrite the default build command
	BuildCommand string
	//Command to run when starting the container
	StartCommand string
	//Additional nix packages to install in the environment
	NixPackages []string
	//Additional apt packages to install in the environment
	AptPackages []string
	//Additional nix libraries to install in the environment
	NixLibraries []string
	//Environment variables to set in the container
	Envs []Env
	//Path to config file
	Config string
}

func (p PlanOptions) Validate() error {
	if p.Path == "" {
		return errors.New("path is required")
	}
	return nil

}

func (p PlanOptions) ToArgs() []string {
	var args []string

	if p.InstallCommand != "" {
		args = append(args, fmt.Sprintf("--install-cmd %q", p.InstallCommand))
	}

	if p.BuildCommand != "" {
		args = append(args, fmt.Sprintf("--build-cmd %q", p.BuildCommand))
	}

	if p.StartCommand != "" {
		args = append(args, fmt.Sprintf("--start-cmd %q", p.StartCommand))
	}

	if len(p.NixPackages) != 0 {
		for _, pkg := range p.NixPackages {
			args = append(args, fmt.Sprintf("--pkgs %q", pkg))
		}
	}

	if len(p.AptPackages) != 0 {
		for _, pkg := range p.AptPackages {
			args = append(args, fmt.Sprintf("--apt %q", pkg))
		}
	}

	if len(p.NixLibraries) != 0 {
		for _, lib := range p.NixLibraries {
			args = append(args, fmt.Sprintf("--libs %q", lib))
		}
	}

	if len(p.Envs) != 0 {
		for _, env := range p.Envs {
			args = append(args, fmt.Sprintf("--env %q", env.Key+"="+env.Value))
		}
	}

	if p.Config != "" {
		args = append(args, fmt.Sprintf("--config %q", p.Config))
	}

	return args
}
