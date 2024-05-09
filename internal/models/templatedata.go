package models

import "github.com/shaynemeyer/go-bnb/internal/forms"

// TemplateData - holds data sent to handlers
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
