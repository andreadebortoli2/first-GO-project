package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// config package can be imported by everywhere but don't import my packages to avoid import cycle(by importing everithing everywhere without logic)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
