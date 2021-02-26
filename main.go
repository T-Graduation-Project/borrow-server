package main

import (
	_ "github.com/T-Graduation-Project/borrow-server/router"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := ghttp.GetServer()
	// s.EnablePProf()
	s.Run()
}
