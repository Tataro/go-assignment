package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"gitlab.com/upaphong/go-assignment/adapters/http"
	"gitlab.com/upaphong/go-assignment/engine"
	"gitlab.com/upaphong/go-assignment/providers/database"
)

func main() {
	provider := database.NewProvider()
	e := engine.NewEngine(provider)

	adapter := http.NewHTTPAdapter(e)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
	provider.Close()
}
