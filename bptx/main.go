package main

// If failed to conn. to db, handle error well
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"wins21.co.kr/soc/aptx/engine"
)

const (
	DefaultPort     = "8080"
	DefaultInterval = 30
)

var (
	fs *flag.FlagSet
)

func init() {
	log.SetPrefix("[Realmon] ")
}

func main() {
	// Set CPU
	runtime.GOMAXPROCS(1)

	// Set flags
	fs = flag.NewFlagSet("", flag.ExitOnError)
	var (
		port     = fs.String("port", DefaultPort, "Default port number")
		interval = fs.Int("interval", DefaultInterval, "Default reload interval")
		isDebug  = fs.Bool("debug", false, "Is debug mode?")
	)
	if *isDebug {
		log.Println("Mode: debug")
	}
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])
	log.Printf("Port:%s, Interval:%d(sec)", *port, *interval)

	//	errChan := make(chan error)
	engine := engine.NewEngine(*interval)
	engine.Start()

	// Start monitoring
	startHTTPServer(port)

	// Stop
	waitForSignals()
}

func startHTTPServer(port *string) error {
	http.HandleFunc("/logs", engine.GetFiletransLog)
	err := http.ListenAndServe("localhost:"+*port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

func printHelp() {
	fmt.Println("dataserver [options]")
	fs.PrintDefaults()
}

func waitForSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalCh:
		log.Println("Signal received, shutting down...")
	}
}
