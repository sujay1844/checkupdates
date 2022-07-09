///usr/bin/true; exec /usr/bin/env go run "$0" "$@"

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/martinlindhe/notify"
)

func main() {
	OS := "fedora"
	stdout := ""

	if OS == "fedora" {
		stdout = fedora()
	} else if OS == "arch" {
		stdout = arch()
	}

	fmt.Println(stdout)
	line_count := strings.Count(stdout, "updates")
	if line_count == 0 {
		fmt.Println("No updates available")
		return
	}
	notify_str := strconv.Itoa(line_count) + " updates are available"
	notify.Notify("Updates", notify_str, "", "")
}

func fedora() string {
	app := "dnf"
	arg0 := "check-update"
	arg1 := "-q"
	cmd := exec.Command(app, arg0, arg1)
	stdout, _ := cmd.Output()
	return string(stdout)
}

func arch() string {
	app := "check-updates"
	cmd := exec.Command(app)
	stdout, _ := cmd.Output()
	return string(stdout)
}
