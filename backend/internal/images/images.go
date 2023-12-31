package images

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Cyan903/Quatracker/backend/internal/database"
	"github.com/Cyan903/Quatracker/backend/pkg/log"
)

type FileLoader struct {
	http.Handler
}

func (h *FileLoader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mid := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/id/"), ".id")

	// Is proper URL?
	if !strings.HasPrefix(r.URL.Path, "/id/") || !strings.HasSuffix(r.URL.Path, ".id") {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	// Is valid ID?
	id, err := strconv.Atoi(mid)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warning.Println("ServeHTTP received invalid ID", mid)

		return
	}

	// Fetch image
	img, err := serve(id)

	if err != nil {
		if errors.Is(err, database.ErrBGNotFound) {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Warning.Println("ServeHTTP internal error")

		return
	}

	w.Write(img)
}
