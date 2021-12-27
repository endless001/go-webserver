package main

import "os"

type Env struct {
	Version string
}


func getEnv() *Env {
	v := os.Getenv("VERSION")
	return &Env{Version: v}
}

