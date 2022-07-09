///usr/bin/true; exec /usr/bin/env go run "$0" "$@"

package main

import (
	"fmt"
	"os/exec"

	"github.com/martinlindhe/notify"
)

func main() {
	notify.Notify("Updates", "foo", "", "")
	app := "echo"
	arg0 := "bar"
	
	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}
