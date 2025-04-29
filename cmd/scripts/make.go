package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var files = map[string]string{
	"internal/{{.Name}}/handler.go": `package {{.Name}}

// Handler functions for {{.PascalName}} module.
`,
	"internal/{{.Name}}/service.go": `package {{.Name}}

// Service provides business logic for the {{.PascalName}} module.
`,
	"internal/{{.Name}}/middleware.go": `package {{.Name}}

// Middleware functions for {{.PascalName}} module.
`,
	"internal/{{.Name}}/routes.go": `package {{.Name}}

// Routes for {{.PascalName}} module.
`,
	"internal/{{.Name}}/database/{{.Name}}/model.go": `package {{.Name}}

// Model definitions for {{.PascalName}} module.
`,
	"internal/{{.Name}}/view/index.blade":  `<!-- Index view for {{.PascalName}} module -->`,
	"internal/{{.Name}}/view/header.blade": `<!-- Header for {{.PascalName}} module -->`,
	"internal/{{.Name}}/view/footer.blade": `<!-- Footer for {{.PascalName}} module -->`,
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
	pascal := toPascalCase(module)

	data := ModuleData{
		Name:       module,
		PascalName: pascal,
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
	var result strings.Builder
	t.Execute(&result, data)
	return result.String()
}

func toPascalCase(input string) string {
	parts := strings.Split(input, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
