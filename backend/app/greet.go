package app

import (
	"fmt"

	"github.com/Cyan903/QuaverBuddy/backend/pkg/log"
)

func (a *App) Greet(name string) string {
	log.Info.Println("Hii")
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
