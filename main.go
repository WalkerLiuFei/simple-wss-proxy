package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"simple-ws-proxy/websocketproxy"
)

var (
	flagBackend = flag.String("backend", "", "Backend URL for proxying")
)

func main() {

	u, err := url.Parse("wss://stream.bybit.com/contract/usdt/public/v3")
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(":80", websocketproxy.NewProxy(u))
	if err != nil {
		log.Fatalln(err)
	}

}
