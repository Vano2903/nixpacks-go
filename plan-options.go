package nixpacks

import (
	"errors"
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

func (o PlanOptions) ToArgs() []string {
	var args []string

	if o.InstallCommand != "" {
		args = append(args, "--install-cmd", o.InstallCommand)
	}

	if o.BuildCommand != "" {
		args = append(args, "--build-cmd", o.BuildCommand)
	}

	if o.StartCommand != "" {
		args = append(args, "--start-cmd", o.StartCommand)
	}

	if len(o.NixPackages) != 0 {
		for _, pkg := range o.NixPackages {
			args = append(args, "--pkgs", pkg)
		}
	}

	if len(o.AptPackages) != 0 {
		for _, pkg := range o.AptPackages {
			args = append(args, "--apt", pkg)
		}
	}

	if len(o.NixLibraries) != 0 {
		for _, lib := range o.NixLibraries {
			args = append(args, "--libs", lib)
		}
	}

	if len(o.Envs) != 0 {
		for _, env := range o.Envs {
			args = append(args, "--env", env.Key+"="+env.Value)
		}
	}

	if o.Config != "" {
		args = append(args, "--config", o.Config)
	}
	return args
}
