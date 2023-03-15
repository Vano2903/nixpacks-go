package nixpacks

import (
	"errors"
)

type (
	BuildOptions struct {
		//Path to the directory to build
		Path string
		//Name for the built image
		Name string
		//Save output directory instead of building the image
		Output string
		//Specify an entire build plan in json used to build the image
		JsonPlan string
		//Tag for the built image
		Tag string
		//Command to install language dependencies.
		InstallCommand string
		//Add labels to the built image
		Labels []Label
		//Command to build the image,
		//it will overwrite the default build command
		BuildCommand string
		//Target platform for built image
		Platform string
		//Unique identifier to key cache.
		//defaults to the current directory
		//TODO: CacheKey string

		//Command to run when starting the container
		StartCommand string
		//Outputs nixpacks related files
		//to the current directory
		CurrentDirectory bool
		//Additional nix packages to install in the environment
		NixPackages []string
		//Additional apt packages to install in the environment
		AptPackages []string
		//Disable caching
		NoCache bool
		//Additional nix libraries to install in the environment
		NixLibraries []string
		//Environment variables to set in the container
		Envs []Env
		//Path to config file
		Config string
		//Do not error when no start command is set
		NoErrorWithoutStartCommand bool
	}

	Env struct {
		Key   string
		Value string
	}

	Label struct {
		Key   string
		Value string
	}
)

func (o *BuildOptions) Validate() error {
	if o.Path == "" {
		return errors.New("path must be specified")
	}
	return nil
}

func (o *BuildOptions) ToArgs() []string {
	var args []string
	if o.Name != "" {
		args = append(args, "--name", o.Name)
	}

	if o.Output != "" {
		args = append(args, "--out", o.Output)
	}

	if o.JsonPlan != "" {
		args = append(args, "--json-plan", o.JsonPlan)
	}

	if o.Tag != "" {
		args = append(args, "--tag", o.Tag)
	}

	if o.InstallCommand != "" {
		args = append(args, "--install-cmd", o.InstallCommand)
	}

	if len(o.Labels) != 0 {
		for _, label := range o.Labels {
			args = append(args, "--label", label.Key+"="+label.Value)
		}
	}

	if o.BuildCommand != "" {
		args = append(args, "--build-cmd", o.BuildCommand)
	}
	if o.Platform != "" {
		args = append(args, "--platform", o.Platform)
	}

	if o.StartCommand != "" {
		args = append(args, "--start-cmd", o.StartCommand)
	}

	if o.CurrentDirectory {
		args = append(args, "--current-directory")
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

	if o.NoCache {
		args = append(args, "--no-cache")
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

	if o.NoErrorWithoutStartCommand {
		args = append(args, "--no-error-without-start-cmd")
	}

	args = append(args, "--verbose")
	return args
}
