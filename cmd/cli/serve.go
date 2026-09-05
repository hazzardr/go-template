package cli

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	RunE: func(_ *cobra.Command, _ []string) error {
		r := chi.NewRouter()
		r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		srv := &http.Server{
			Addr:              ":8080",
			Handler:           r,
			ReadHeaderTimeout: 5 * time.Second,
		}
		slog.Info("starting server", "addr", srv.Addr)
		return srv.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
