// minimal example CLI used for binary size checking

package main

import (
	"github.com/urfave/cli/v2"
)

func main() {
	_ = (&cli.App{}).Run([]string{""})
}
