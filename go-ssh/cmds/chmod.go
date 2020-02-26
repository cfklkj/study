package cmds

import (
	"fmt"
	"strings"
)

func (c *Cmds) Cmd_showMod(path string) string {
	c.BindCheck(str_lsMod, c.showMod)
	sh := "ls -ld " + path
	c.Run(str_lsMod, sh)
	return c.getResult(str_lsMod).(string)
}

func (c *Cmds) Cmd_chmodR(mod, path string) bool {
	c.BindCheck(str_chmod, c.chmodR)
	sh := "chmod -R " + mod + " " + path
	c.Run(str_chmod, sh)
	return c.getResult(str_chmod).(bool)
}

func (c *Cmds) Cmd_chown(owner_group, path string) bool {
	sh := "chown -R " + owner_group + " " + path
	c.Run(str_chown, sh)
	mod := c.Cmd_showMod(path)
	sub := strings.Replace(owner_group, ":", " ", -1)
	if strings.Contains(mod, sub) {
		return true
	}
	return false
}

func (c *Cmds) showMod(msg string) {
	fmt.Println(msg)
	c.setResult(str_lsMod, msg)
}

func (c *Cmds) chmodR(msg string) {
	if strings.Contains(msg, "invalid") || strings.Contains(msg, "permitted") {
		c.setResult(str_chmod, false)
	} else {
		c.setResult(str_chmod, true)
	}
}
