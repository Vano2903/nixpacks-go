package nixpacks

import (
	"context"
	"fmt"
	"os/exec"
)

type PlanCmd struct {
	cmd *exec.Cmd
}

func (n Nixpacks) Plan(ctx context.Context, opt PlanOptions) (*PlanCmd, error) {
	if err := opt.Validate(); err != nil {
		return nil, err
	}

	fmt.Println("Building: ", opt.Path)
	fmt.Println("args:", opt.ToArgs())

	cmd := exec.CommandContext(ctx, n.commandPath, PlanCommand, opt.Path)
	cmd.Args = append(cmd.Args, opt.ToArgs()...)

	return &PlanCmd{
		cmd: cmd,
	}, nil
}

func (c *PlanCmd) Result() (PlanOutput, error) {
	b := PlanOutput{}
	fmt.Println(c.cmd.String())
	out, err := c.cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		if err.Error() == "signal: killed" {
			return b, err
		}
	}
	b.Response = out
	b.Parse()
	b.Response = nil //! remove, only for debugging
	fmt.Printf("out: %+v\n", b)
	fmt.Println()
	return b, err

}
