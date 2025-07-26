package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Github.com/Gaintlord/hospital_management/internal/database"
	"Github.com/Gaintlord/hospital_management/internal/routes"
)

func main() {

	router := http.NewServeMux()

	Db := database.Database{}
	if err := Db.Connect(); err != nil {
		slog.Info("database could not be mounted")
	} else {
		databaseinstance := Db.GetDB()
		fmt.Println(databaseinstance, "okay working fine")
	}

	routes.Registeradminroute(router, &Db.Db)
	// routes.Registerreceproute(router)
	// routes.Registerdocroute(router)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	/*
		gracefully closing the server
	*/
	{
		done := make(chan os.Signal, 1)

		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			slog.Info("#### Server Up and Running ####")
			err := server.ListenAndServe()

			if err.Error() == "http: Server closed" {
				slog.Info("server has been closed")
			} else {
				slog.Info("Failed to start Server", slog.String("error", err.Error()))
			}
		}()

		<-done

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Info("failed to shutdown", slog.String("error :", err.Error()))
		} else {
			slog.Info("###### server shutdown Gracefully #####")
		}
	}
}
