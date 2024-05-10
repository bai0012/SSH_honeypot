package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

var shellProgram = []string{"sh"}

type readLiner interface {
	ReadLine() (string, error)
}

type commandContext struct {
	args    []string
	stdin   readLiner
	stdout  io.Writer
	stderr  io.Writer
	pty     bool
	user    string
}

type command interface {
	execute(context commandContext) (uint32, error)
}

var commands = make(map[string]command)

func init() {
	// Initialize commands in a single place
	registerCommands()
}

func registerCommands() {
	commands["sh"] = cmdShell{}
	commands["true"] = cmdTrue{}
	commands["false"] = cmdFalse{}
	commands["echo"] = cmdEcho{}
	commands["cat"] = cmdCat{}
	commands["su"] = cmdSu{}
	commands["ifconfig"] = cmdIfconfig{}
	commands["uname"] = cmdUname{}
}

func executeProgram(context commandContext) (uint32, error) {
	if len(context.args) == 0 {
		return 0, nil // Early return for no args
	}
	cmd, exists := commands[context.args[0]]
	if !exists {
		fmt.Fprintf(context.stderr, "%v: command not found\n", context.args[0])
		return 127, nil // Changed to return nil error after printing
	}
	return cmd.execute(context)
}

type cmdShell struct{}

func (c cmdShell) execute(context commandContext) (uint32, error) {
	prompt := getPrompt(context) // Simplified prompt selection
	var lastStatus uint32
	for {
		fmt.Fprint(context.stdout, prompt) // Error handling omitted for brevity
		line, err := context.stdin.ReadLine()
		if err != nil {
			return lastStatus, err // Return on read errors
		}
		if len(line) == 0 {
			continue // Skip empty lines
		}
		args := strings.Fields(line)
		if args[0] == "exit" {
			return parseExitStatus(args, lastStatus) // Simplified exit handling
		}
		newContext := context
		newContext.args = args
		lastStatus, err = executeProgram(newContext) // Simplify error handling
		if err != nil {
			return lastStatus, err
		}
	}
}

// Additional functions for cmdShell
func getPrompt(context commandContext) string {
	if context.pty {
		if context.user == "root" {
			return "# "
		}
		return "$ "
	}
	return ""
}

func parseExitStatus(args []string, lastStatus uint32) (uint32, error) {
	if len(args) > 1 {
		status, err := strconv.ParseUint(args[1], 10, 32)
		if err != nil {
			return 255, nil // Return 255 on parse error
		}
		return uint32(status), nil
	}
	return lastStatus, nil
}

type cmdTrue struct{}

func (cmdTrue) execute(context commandContext) (uint32, error) {
	return 0, nil
}

type cmdFalse struct{}

func (cmdFalse) execute(context commandContext) (uint32, error) {
	return 1, nil
}

type cmdEcho struct{}

func (cmdEcho) execute(context commandContext) (uint32, error) {
	_, err := fmt.Fprintln(context.stdout, strings.Join(context.args[1:], " "))
	return 0, err
}

type cmdCat struct{}

func (cmdCat) execute(context commandContext) (uint32, error) {
    numberLine := false
    files := []string{}

    for _, arg := range context.args[1:] {
        if arg == "-n" {
            numberLine = true
        } else {
            files = append(files, arg)
        }
    }

    if len(files) > 0 {
        for _, file := range files {
			if file == "/proc/cpuinfo" {
                _, err := fmt.Fprint(context.stdout, CpuInfo)
                return 0, err
            } else {
            if _, err := fmt.Fprintf(context.stderr, "%v: %v: No such file or directory\n", context.args[0], file); err != nil {
                return 0, err
            }
		}
        }
        return 1, nil
    }

    var line string
    var err error
    lineNumber := 1
    for err == nil {
        line, err = context.stdin.ReadLine()
        if err == nil {
            if numberLine {
                _, err = fmt.Fprintf(context.stdout, "%6d  %s\n", lineNumber, line)
                lineNumber++
            } else {
                _, err = fmt.Fprintln(context.stdout, line)
            }
        }
    }
    return 0, err
}


type cmdSu struct{}

func (cmdSu) execute(context commandContext) (uint32, error) {
	newContext := context
	newContext.user = "root"
	if len(context.args) > 1 {
		newContext.user = context.args[1]
	}
	newContext.args = shellProgram
	return executeProgram(newContext)
}

type cmdUname struct{}

func (cmdUname) execute(context commandContext) (uint32, error) {
    currentTime := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")
    output := fmt.Sprintf("Linux Server 6.8.0-31-generic #31-Ubuntu SMP PREEMPT_DYNAMIC %s x86_64 x86_64 x86_64 GNU/Linux", currentTime)
    _, err := fmt.Fprintln(context.stdout, output)
    return 0, err
}

func init() {
    commands["uname"] = cmdUname{}
}

type cmdIfconfig struct{}

func (cmdIfconfig) execute(context commandContext) (uint32, error) {
	output := IfconfigOutput
	_, err := fmt.Fprintln(context.stdout, output)
	return 0, err
}