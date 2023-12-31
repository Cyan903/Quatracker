package images

import (
	"errors"
	"os"

	"github.com/Cyan903/Quatracker/backend/internal/database"
	"github.com/Cyan903/Quatracker/backend/pkg/log"
)

func serve(id int) ([]byte, error) {
	img, err := database.GetImage(id)

	if err != nil {
		if errors.Is(err, database.ErrBGNotFound) {
			return []byte{}, err
		}

		log.Error.Println("could not serve image from", id)
		log.Error.Println("server error", err)

		return []byte{}, err
	}

	file, err := os.ReadFile(img)

	if err != nil {
		log.Warning.Println("could not read", img)
		log.Warning.Println("read error", err)

		return []byte{}, err
	}

	return file, nil
}
