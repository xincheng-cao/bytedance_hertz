// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"
)

func main() {
	h := server.Default(server.WithHostPorts("0.0.0.0:13001"))
	/*
		h.Use(basic_auth.BasicAuth(map[string]string{
			"chao": "yang",
		}))
	*/

	// SetLevel sets the level of logs below which logs will not be output.
	hlog.SetLevel(hlog.LevelDebug)

	f, err := os.Create("./hertz.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// SetOutput sets the output of default logger. By default, it is stderr.
	hlog.SetOutput(f)

	register(h)
	h.Spin()
}
