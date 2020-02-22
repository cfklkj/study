package run

import (
	"context"
	"fmt"
	"os/exec"
)

//通过管道同步获取日志的函数

func (c *DoRunInfo) commandCheck(ctx context.Context, name, command string, arg []string) {
	var cmd *exec.Cmd
	switch len(arg) {
	case 0:
		cmd = exec.CommandContext(ctx, command)
	case 1:
		cmd = exec.CommandContext(ctx, command, arg[0])
	case 2:
		cmd = exec.CommandContext(ctx, command, arg[0], arg[1])
	case 3:
		cmd = exec.CommandContext(ctx, command, arg[0], arg[1], arg[2])
	case 4:
		cmd = exec.CommandContext(ctx, command, arg[0], arg[1], arg[2], arg[3])
	default:
		fmt.Println(arg)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(name, "err", err)
		c.PrintMsg(name, "unInstall")
		return
	}
	err := cmd.Wait()
	if err != nil {
		fmt.Println(name, "err", err)
	} else {
		c.PrintMsg(name, "installed")
	}
}
