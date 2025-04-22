package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var files = map[string]string{
	"internal/{{.Name}}/handler.go":    `package {{.Name}}\n\n// Handler ...`,
	"internal/{{.Name}}/model.go":      `package {{.Name}}\n\ntype {{.PascalName}} struct {}`,
	"internal/{{.Name}}/repository.go": `package {{.Name}}\n\n// Repository functions`,
	"internal/{{.Name}}/service.go":    `package {{.Name}}\n\n// Business logic`,
	"internal/{{.Name}}/middleware.go": `package {{.Name}}\n\n// Middleware functions`,
	"internal/{{.Name}}/routes.go":     `package {{.Name}}\n\n// Routes setup`,
}

type ModuleData struct {
	Name       string
	PascalName string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a module name.")
		return
	}

	module := strings.ToLower(os.Args[1])
	data := ModuleData{
		Name:       module,
		PascalName: strings.Title(module),
	}

	for tmplPath, tmplContent := range files {
		path := parseTemplate(tmplPath, data)

		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			fmt.Println("Failed to create directory:", err)
			continue
		}

		if _, err := os.Stat(path); err == nil {
			fmt.Println("Skipped (exists):", path)
			continue
		}

		content := parseTemplate(tmplContent, data)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Println("Failed to create file:", err)
		} else {
			fmt.Println("Created:", path)
		}
	}
}

func parseTemplate(tmpl string, data ModuleData) string {
	t := template.Must(template.New("tmpl").Parse(tmpl))
	var result string
	buf := &stringWriter{&result}
	t.Execute(buf, data)
	return result
}

type stringWriter struct {
	s *string
}

func (w *stringWriter) Write(p []byte) (int, error) {
	*w.s += string(p)
	return len(p), nil
}
