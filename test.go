package force

import (
	"fmt"
	"strconv"
)

var cmdTest = &Command{
	Usage: "test (all | classname...)",
	Short: "Run apex tests",
	Long: `
Run apex tests

Test Options
  -namespace=<namespace>     Select namespace to run test from

Examples:

  force test all
  force test Test1 Test2 Test3
  force test -namespace=ns Test4 
`,
}

func init() {
	cmdTest.Run = runTests
}

var (
	namespaceTestFlag = cmdTest.Flag.String("namespace", "", "namespace to run tests in")
)

func runTests(cmd *Command, args []string) {
	if len(args) < 1 {
		ErrorAndExit("must specify tests to run")
	}
	force, _ := ActiveForce()
	output, err := force.Partner.RunTests(args, *namespaceTestFlag)
	success := false
	if err != nil {
		ErrorAndExit(err.Error())
	} else {
		//working on a better way to do this - catches when no class are found and ran
		if output.NumberRun == 0 {
			fmt.Println("Test classes specified not found")
		} else {
			var percent string
			fmt.Println("Coverage:")
			fmt.Println()
			for index := range output.NumberLocations {
				if output.NumberLocations[index] != 0 {
					percent = strconv.Itoa(((output.NumberLocations[index]-output.NumberLocationsNotCovered[index])/output.NumberLocations[index])*100) + "%"
				} else {
					percent = "0%"
				}
				fmt.Println("  " + percent + "   " + output.Name[index])
			}
			fmt.Println()
			fmt.Println()
			fmt.Println("Results:")
			fmt.Println()
			for index := range output.SMethodNames {
				fmt.Println("  [PASS]  " + output.SClassNames[index] + "::" + output.SMethodNames[index])
			}

			for index := range output.FMethodNames {
				fmt.Println("  [FAIL]  " + output.FClassNames[index] + "::" + output.FMethodNames[index] + ": " + output.FMessage[index])
				fmt.Println("    " + output.FStackTrace[index])
			}
			fmt.Println()
			fmt.Println()

			success = len(output.FMethodNames) == 0
		}

		// Handle notifications
		notifySuccess("test", success)
	}
}
