package main

import (
	"github.com/codegangsta/cli"
)

var cmds map[string][]cli.Command

// ensure the cmds has command list for this name
func mustHaveCmds(n string) {
	if _, ok := cmds[n]; !ok {
		cmds[n] = make([]cli.Command, 0, 0)
	}
}

// add a sub-command to command list of a certain name
// which would be used by the certain command as
// its `Subcommands` parameter
func appendCmds(n string, cmd cli.Command) {

	// initialize cmds
	if cmds == nil {
		cmds = make(map[string][]cli.Command)
	}

	// ensure the command list of the name exists
	mustHaveCmds(n)

	// append the sub-command to the command list
	cmds[n] = append(cmds[n], cmd)
}
