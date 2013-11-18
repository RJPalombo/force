package main

import (
	"fmt"
	"strings"
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
	me, err := force.Whoami()
	if err != nil {
		ErrorAndExit(err.Error())
	} else if len(args) == 0 {
		DisplayForceRecord(me)
	} else {
		parts := strings.Split(force.Credentials.Id, "/")
		records, err := force.Query(fmt.Sprintf("select Name, AboutMe From User Where Id = '%s'", parts[len(parts)-1]))
		if err != nil {
			ErrorAndExit(err.Error())
		} else {
			if records[0]["AboutMe"] != nil {
				fmt.Printf("\nAbout %s\n\nOn twitter: @dcarroll\n\n%s\n\n", records[0]["Name"], records[0]["AboutMe"])
			} else {
				fmt.Printf("\nAbout %s\n\nOn twitter: @dcarroll\n\n", records[0]["Name"])
			}
		}	
	}
}
