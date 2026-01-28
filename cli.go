package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	Appname = "tene3rm"
)

type CmdArgs struct {
	LogOutput string
	DryRun    bool
	Args      []string
}

func ParseArgs() (*CmdArgs, error) {
	c := CmdArgs{}

	flag.Usage = flagHelpMessage
	flag.StringVar(&c.LogOutput, "log-output", logOutputStderr, "log message output")
	flag.BoolVar(&c.DryRun, "dry-run", false, "dry run")
	flag.Parse()
	c.Args = flag.Args()

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return &c, nil
}

func flagHelpMessage() {
	cmd := Appname
	fmt.Fprintln(os.Stderr, fmt.Sprintf("%s carefully deletes files", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s [OPTIONS] [files...]", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s sample.txt", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
}

func (c *CmdArgs) Validate() error {
	if len(c.Args) < 1 {
		return fmt.Errorf("Must need files")
	}

	switch c.LogOutput {
	case logOutputStdout, logOutputStderr:
	default:
		return fmt.Errorf("log-output is invalid: %s", c.LogOutput)
	}

	for _, file := range c.Args {
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			return fmt.Errorf("%s file doesn't exist", file)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("%s permission denied", file)
		}
	}

	return nil
}
