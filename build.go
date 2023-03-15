package nixpacks

import (
	"context"
	"os/exec"
)

func (n Nixpacks) Build(ctx context.Context, opt BuildOptions) (*BuildCmd, error) {
	if err := opt.Validate(); err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(ctx, n.commandPath, BuildCommand, opt.Path)
	cmd.Args = append(cmd.Args, opt.ToArgs()...)

	return &BuildCmd{
		cmd: cmd,
	}, nil
}

type BuildCmd struct {
	cmd *exec.Cmd
}

func (c *BuildCmd) Result() (BuildOutput, error) {
	n := BuildOutput{}
	out, err := c.cmd.CombinedOutput()
	if err != nil {
		if err.Error() == "signal: killed" {
			return n, err
		}
	}
	n.Response = out
	n.IsBrokenImage = err != nil
	n.Parse()
	return n, err
}
