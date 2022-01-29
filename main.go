package main

import (
	"strings"

	"github.com/m-mcdonagh/shop-and-shop/shell"
)

func main() {
	commands := map[string]shell.CMD{
		"test": shell.CMD{
			Help: func(args ...string) string {
				return "HELP " + strings.Join(args, " ")
			},
			Run: func(args ...string) string {
				return "RUN " + strings.Join(args, " ")
			},
		},
	}

	shell.Loop("$ ", commands)
}
