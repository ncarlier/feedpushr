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

Copyright (C) 2018 Nicolas Carlier

This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it under certain conditions:
GNU General Public License v3.0+ (see LICENSE or https://www.gnu.org/licenses/gpl-3.0.txt).
`, version)
}

func init() {
	expvar.NewString("version").Set(Version)
	expvar.NewString("rev").Set(GitCommit)
}
