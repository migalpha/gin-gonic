package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gin-gonic/pkg/setting"
	"gin-gonic/routers"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "api",
				Usage: "Launch gin-gonic API",
				Action: func(c *cli.Context) error {
					router := routers.InitRouter()

					s := &http.Server{
						Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
						Handler:        router,
						ReadTimeout:    setting.ReadTimeout,
						WriteTimeout:   setting.WriteTimeout,
						MaxHeaderBytes: 1 << 20,
					}

					s.ListenAndServe()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
