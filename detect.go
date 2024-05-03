package nixpacks

import (
	"context"
	"os/exec"
)

type DetectCmd struct {
	cmd *exec.Cmd
	err error
}

func (n Nixpacks) Detect(ctx context.Context, opt DetectOptions) *DetectCmd {
	if err := opt.Validate(); err != nil {
		return &DetectCmd{
			cmd: nil,
			err: err,
		}
	}

	cmd := exec.CommandContext(ctx, n.commandPath, DetectCommand)
	cmd.Args = append(cmd.Args, opt.ToArgs()...)
	cmd.Args = append(cmd.Args, opt.Path)

	return &DetectCmd{
		cmd: cmd,
		err: nil,
	}
}

func (c *DetectCmd) Error() error {
	return c.err
}

func (c *DetectCmd) Result() (DetectOutput, error) {
	b := DetectOutput{}
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
