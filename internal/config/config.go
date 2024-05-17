package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/shaynemeyer/go-bnb/internal/models"
)

// AppConfig hold the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
