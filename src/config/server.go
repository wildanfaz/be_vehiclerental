package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/wildanfaz/vehicle_rental/src/routers"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start app",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var address string = "0.0.0.0:8080"

		if port := os.Getenv("PORT"); port != "" {
			address = "0.0.0.0:" + port
		}

		srv := &http.Server{
			Addr:         address,
			WriteTimeout: time.Second * 20,
			ReadTimeout:  time.Second * 20,
			IdleTimeout:  time.Second * 100,
			Handler:      mainRoute,
		}

		fmt.Print("running on port http://", address)
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}
}
