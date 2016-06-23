package main

import (
	"flag"
	"fmt"
	"github.com/AlKoFDC/eva.support/logger"
	"github.com/AlKoFDC/eva.support/slack"
	"os"
	"os/signal"
	"syscall"
)

// The name of the bot, that it will react to.
const eva = "eva"

// Flags for API token.
// Set alternative name for bot.
var name = flag.String("n", eva, "alternative `name` for bot")

// Get token from command line.
var token = flag.String("t", "", "slack bot `token`")

// Get token from file.
var tokenFile = flag.String("f", "", "`file` with slack bot token")

// Print unknown messages.
var printUnknownMessages = flag.Bool("u", false, "print unknown messages")

// Help
var help = flag.Bool("h", false, "show help")

// Usage
var printUsage = func(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Error: %s", err))
	}
	fmt.Fprintln(os.Stderr, fmt.Sprintf("Usage: %s ([-t ]<token>|-f <file>) [-n <name>] [-u] [-h]", os.Args[0]))
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if *help {
		printUsage(nil)
		os.Exit(0)
	}
	tf := *tokenFile
	t := *token

	if tf == "" && t == "" && len(os.Args) != 2 {
		printUsage(fmt.Errorf("need to provide either token or file with token"))
		os.Exit(1)
	}
	if tf == "" && t == "" {
		t = os.Args[1]
	}
	if t == "" {
		// TODO Read token from File
		printUsage(fmt.Errorf("reading token from file not implemented yet"))
		os.Exit(1)
	}

	// Connect to slack.
	wsHandler, err := slack.Connect(t)
	if err != nil {
		logger.Error.Println("Error while connecting:", err)
		os.Exit(1)
	}
	defer wsHandler.Close()
	logger.Standard.Println(fmt.Sprintf("Connected as %s.", wsHandler.ID))

	// Set the name.
	wsHandler.Name = *name

	// Set the print flag.
	wsHandler.PrintUnknown = *printUnknownMessages

	go wsHandler.Handle()

	// Handle SIGINT and SIGTERM for graceful shutdowns.
	systemInterruptChannel := make(chan os.Signal)
	defer close(systemInterruptChannel)
	signal.Notify(systemInterruptChannel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case interruptSignal := <-systemInterruptChannel:
		logger.Standard.Println("Finishing because of signal:", interruptSignal)
		return
	}
}
