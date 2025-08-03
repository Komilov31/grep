package app

import (
	"log"

	"github.com/Komilov31/grep/internal/flags"
	"github.com/Komilov31/grep/internal/grep"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run() {
	flags := flags.Parse()
	grep := grep.New(flags)

	if err := grep.ProcessProgram(); err != nil {
		log.Fatal(err)
	}
}
