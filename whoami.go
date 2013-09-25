package main

import (
	"fmt"
)

var cmdWhoami = &Command{
	Run:   runWhoami,
	Usage: "whoami [account]",
	Short: "Show information about the active account",
	Long: `
Show information about the active account

Examples:

  force whoami
`,
}

func runWhoami(cmd *Command, args []string) {
	force, _ := ActiveForce()
	user, err := force.Get("User", force.Credentials.Id)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		DisplayForceObject(user)
	}
}
