package main

import (
	tzj_kefu "./src/kefu"
)

func main() {
	handle := tzj_kefu.NewHttp()
	handle.Listen()
}
