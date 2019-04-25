package main

import (
	"expvar"
	"fmt"
)

// Version of the app
var Version = ""

// GitCommit hash
var GitCommit = "HEAD"

func printVersion() {
	version := Version
	if version == "" {
		version = GitCommit
	}
	fmt.Printf(`feedpushr (%s)
Copyright (C) 2018 Nunux, Org.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Nicolas Carlier.`, version)
}

func init() {
	expvar.NewString("version").Set(Version)
	expvar.NewString("rev").Set(GitCommit)
}
