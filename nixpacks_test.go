package nixpacks

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestBuild(t *testing.T) {
	t.Run("default build", func(t *testing.T) {
		n, err := NewNixpacks()
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cmd, err := n.Build(ctx, BuildOptions{Path: "/home/vano/go/nixpacks-go/testing"})
		if err != nil {
			t.Fatal(err)
		}

		out, err := cmd.Result()
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%+v\n", out)
	})

	t.Run("build with broken code", func(t *testing.T) {
		n, err := NewNixpacks()
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cmd, err := n.Build(ctx, BuildOptions{Path: "/home/vano/go/nixpacks-go/testing-broken"})
		if err != nil {
			t.Fatal(err)
		}

		out, err := cmd.Result()
		if err == nil {
			t.Fatal(err)
		}

		t.Log(out.BuildError)
	})

	t.Run("build with shutdown", func(t *testing.T) {
		n, err := NewNixpacks()
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cmd, err := n.Build(ctx, BuildOptions{Path: "/home/vano/go/nixpacks-go/testing"})
		if err != nil {
			t.Fatal(err)
		}

		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("canceling")
			cancel()
		}()

		_, err = cmd.Result()
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("build with args", func(t *testing.T) {
		n, err := NewNixpacks()
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cmd, err := n.Build(ctx, BuildOptions{Path: "/home/vano/go/nixpacks-go/testing", Name: fmt.Sprintf("test-%s", time.Now().String())})
		if err != nil {
			t.Fatal(err)
		}
		out, err := cmd.Result()
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(out.ImageName)
	})
}

func TestPlan(t *testing.T) {
	n, err := NewNixpacks()
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, err := n.Plan(ctx, PlanOptions{Path: "/home/vano/go/nixpacks-go/testing"})
	if err != nil {
		t.Fatal(err)
	}

	out, err := cmd.Result()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", out)
}
