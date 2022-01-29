package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func invalidCMD(cmd string) string {
	return fmt.Sprintf(`"%s" is not a valid command. Type "help" for a list of commands`, cmd)
}

func help(commands map[string]CMD, args ...string) string {
	if len(args) == 0 && len(commands) != 0 {
		builder := strings.Builder{}
		builder.WriteString("valid commands:\n")
		for key, _ := range commands {
			builder.WriteString("\t")
			builder.WriteString(key)
			builder.WriteString("\n")
		}
		builder.WriteString("\thelp\n\tquit")
		return builder.String()
	} else if cmd, ok := commands[args[0]]; ok {
		return cmd.Help(args[1:]...)
	} else {
		return invalidCMD(args[0])
	}
}

func Loop(ps1 string, commands map[string]CMD) {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		out.WriteString(ps1)
		out.Flush()

		if !in.Scan() {
			break
		}

		args := strings.Fields(in.Text())

		if len(args) == 0 {
			continue
		} else if cmd, ok := commands[args[0]]; ok {
			out.WriteString(cmd.Run(args[1:]...))
		} else if args[0] == "help" {
			out.WriteString(help(commands, args[1:]...))
		} else if args[0] == "quit" {
			break
		} else {
			out.WriteString(invalidCMD(args[0]))
		}
		out.WriteString("\n")
	}
	out.WriteString("\n")
}
