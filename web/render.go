package web

import (
	"fmt"
	// "html/template"
	"net/http"
	"path/filepath"

	"github.com/arschles/go-bindata-html-template"

	"github.com/oxtoacart/bpool"
)

//go:generate go-bindata -pkg web includes layouts

var (
	templates map[string]*template.Template
	bufpool   *bpool.BufferPool
	loaded    = false
)

func raw(x string) interface{} { return template.HTML(x) }

// renderTemplate is a wrapper around template.ExecuteTemplate.
// It writes into a bytes.Buffer before writing to the http.ResponseWriter to catch
// any errors resulting from populating the template.
func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	loadTemplates()

	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist", name)
	}

	// Create a buffer to temporarily write to and check if any errors were encountered.
	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		return err
	}

	// The X-Frame-Options HTTP response header can be used to indicate whether
	// or not a browser should be allowed to render a page in a <frame>,
	// <iframe> or <object> . Sites can use this to avoid clickjacking attacks,
	// by ensuring that their content is not embedded into other sites.
	w.Header().Set("X-Frame-Options", "deny")
	// Set the header and write the buffer to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}

func loadTemplates() {
	if loaded {
		return
	}

	templates = make(map[string]*template.Template)

	bufpool = bpool.NewBufferPool(64)

	layoutTemplates := map[string][]string{
		"layouts/index.html": {
			"includes/index.html",
		},
		"layouts/outside.html": {
			"includes/register.html",
			"includes/login.html",
		},
		"layouts/inside.html": {
			"includes/authorize.html",
			"includes/account.html",
		},
		"layouts/profile.html": {
			"includes/clients.html",
			"includes/new_client.html",
			"includes/client_detail.html",
		},
	}

	funcMap := template.FuncMap{
		"raw": raw,
	}

	for layout, includes := range layoutTemplates {
		for _, include := range includes {
			files := []string{include, layout}

			templates[filepath.Base(include)] = template.Must(template.New("", Asset).Funcs(funcMap).ParseFiles(files...))
		}
	}

	loaded = true
}
