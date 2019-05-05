package main

import (
	"flag"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	rootFlag     = flag.String("root", ".", "Web root")
	passwordFlag = flag.String("password", "duplex:agreements", "A comma-separated list of <user>:<password> pairs.")
	bindFlag     = flag.String("bind", "0.0.0.0:8080", "Where to bind. (0.0.0.0:<port> means listen on any interface, 127.0.0.1:<port> means listen only on localhost.)")
)

func getAccounts() gin.Accounts {
	accounts := gin.Accounts{}
	if *passwordFlag == "" {
		return nil
	}
	for _, line := range strings.Split(*passwordFlag, ",") {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}
		accounts[parts[0]] = parts[1]
	}
	return accounts
}

func main() {
	flag.Parse()

	r := gin.Default()
	accounts := getAccounts()
	if accounts != nil {
		r.Use(gin.BasicAuth(accounts))
	}
	r.StaticFS("/", http.Dir(*rootFlag))
	r.Run(*bindFlag)
}
