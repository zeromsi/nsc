package main

import (
	"github.com/nats-io/nsc/cmd"
	"log"
)

// the cli injects the version
var version = "0.0.0-dev"

func main() {
	//srv := cmd.New()
	cmd.SetToolName("nsc")
	conf, err := cmd.LoadOrInit("nats-io/nsc", cmd.NscHomeEnv)
	if err != nil {
		log.Fatal(err)
	}
	conf.SetVersion(version)
	cmd.Execute()
	//cmd.Routes(srv)
//	srv.Logger.Fatal(srv.Start(":" + cmd.ServerPort))
}
