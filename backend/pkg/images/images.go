package images

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

func routes() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/home/hwk/.local/share/Steam/steamapps/common/Quaver/Songs/27013 - 763/HOME.jpg")
	})
}

func Serve(port int) error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Info.Printf("Running on :%d", port)
	return srv.ListenAndServe()
}
