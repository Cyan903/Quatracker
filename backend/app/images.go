package app

import "github.com/Cyan903/Quatracker/backend/internal/images"

func (a *App) NewImages() *images.FileLoader {
	return &images.FileLoader{}
}
