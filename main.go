package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/kolide/kit/logutil"
	"github.com/kolide/launcher/pkg/osquery/table"
	"github.com/osquery/osquery-go"
)

func main() {
	socket := flag.String("socket", "", "Path to osquery socket file")
	flag.Parse()
	if *socket == "" {
		log.Fatalf(`Usage: %s --socket SOCKET_PATH`, os.Args[0])
	}

	server, err := osquery.NewExtensionManagerServer("launcher-tables", *socket)
	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	client, err := osquery.NewClient(*socket, 10*time.Second)
	if err != nil {
		log.Fatalf("Error creating osquery client: %s\n", err)
	}

	tables := table.PlatformTables(client, logutil.NewCLILogger(false), "")

	for _, t := range tables {
		server.RegisterPlugin(t)
	}

	log.Println("start server")
	server.Run()
}
