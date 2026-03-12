/*
 * @Author: liweidong
 * @Date: 2026-03-12 11:52:20
 * @LastEditors: liweidong
 * @LastEditTime: 2026-03-12 11:54:26
 * @Description:
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
