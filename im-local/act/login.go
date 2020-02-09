package act

import "encoding/json"
import "../module/localMem"
import define "../define"

type Login struct {
	mem *localMem.LocalMem
}

func NewLogin() *Login {
	ret := new(Login)
	ret.mem = localMem.NewLocalMem()
	return ret
}

func (c *Login) initUser() {

}

func (c *Login) GetUserId(data string) string {
	type loginInfo struct {
		Username, Password string
	}
	var info loginInfo
	json.Unmarshal([]byte(data), &info)
	return c.getUserId(info.Username, info.Password)
}
func (c *Login) getUserId(name, pwd string) string {
	return name
}

func (c *Login) SetCon(name string, con interface{}) {
	c.mem.Hset(define.Act_user, name, con)
}
func (c *Login) GetCon(name string) interface{} {
	return c.mem.Hget(define.Act_user, name)
}
func (c *Login) DelCon(name string) {
	c.mem.Hdel(define.Act_user, name)
}

func (c *Login) GetUsers() map[interface{}]interface{} {
	users := c.mem.Hfield(define.Act_user)
	return users.Hwnd()
}
