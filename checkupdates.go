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
	stdout := ""

	// I want this script to work for both Fedora and Arch Linux
	OS := "fedora"
	// OS := "arch"

	// The command appropriate to the distro is executed
	if OS == "fedora" {
		stdout = fedora()
	} else if OS == "arch" {
		stdout = arch()
	} else {
		fmt.Println("Only Fedora and Arch Linux supported.")
	}

	// Count the number of updates by counting the occurence of the keyword updates.
	line_count := strings.Count(stdout, "updates")

	// Exit if no updates
	if line_count == 0 {
		fmt.Println("No updates available")
		return
	}

	// Send notification
	notify_str := strconv.Itoa(line_count) + " updates are available"
	notify.Notify("Updates", notify_str, "", "")

	// Print the updates to the terminal
	fmt.Println(stdout)
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
