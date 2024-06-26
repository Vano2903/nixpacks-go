package nixpacks

import (
	"context"
	"os/exec"
)

type PlanCmd struct {
	cmd *exec.Cmd
	err error
}

func (n Nixpacks) Plan(ctx context.Context, opt PlanOptions) *PlanCmd {
	if err := opt.Validate(); err != nil {
		return &PlanCmd{
			cmd: nil,
			err: err,
		}
	}

	cmd := exec.CommandContext(ctx, n.commandPath, PlanCommand)
	cmd.Args = append(cmd.Args, opt.ToArgs()...)
	cmd.Args = append(cmd.Args, opt.Path)

	return &PlanCmd{
		cmd: cmd,
		err: nil,
	}
}

func (c *PlanCmd) Error() error {
	return c.err
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
