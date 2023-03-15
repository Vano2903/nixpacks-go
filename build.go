package nixpacks

import (
	"context"
	"fmt"
	"os/exec"
)

func (n Nixpacks) Build(ctx context.Context, opt BuildOptions) (*BuildCmd, error) {
	if err := opt.Validate(); err != nil {
		return nil, err
	}

	fmt.Println("Building: ", opt.Path)
	fmt.Println("args:", opt.ToArgs())

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
	fmt.Println(c.cmd.String())
	out, err := c.cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		if err.Error() == "signal: killed" {
			return n, err
		}
	}
	n.Response = out
	n.IsBrokenImage = err != nil
	n.Parse()
	fmt.Printf("out: %+v\n", n)
	fmt.Println()
	return n, err
}
