package nixpacks

import (
	"context"
	"os/exec"
)

type PlanCmd struct {
	cmd *exec.Cmd
}

func (n Nixpacks) Plan(ctx context.Context, opt PlanOptions) (*PlanCmd, error) {
	if err := opt.Validate(); err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(ctx, n.commandPath, PlanCommand)
	cmd.Args = append(cmd.Args, opt.ToArgs()...)
	cmd.Args = append(cmd.Args, opt.Path)

	return &PlanCmd{
		cmd: cmd,
	}, nil
}

func (c *PlanCmd) Result() (PlanOutput, error) {
	b := PlanOutput{}
	out, err := c.cmd.CombinedOutput()
	if err != nil {
		if err.Error() == "signal: killed" {
			return b, err
		}
	}
	b.Response = out
	b.Parse()
	return b, err

}
