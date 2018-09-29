package main

import (
	"snippetbox.org/pkg/models"
)

type App struct {
	Database  *models.Database
	HTMLDir   string
	StaticDir string
}
