package sandbox

import (
	"bytes"
	"text/template"
)

const mainTemplate = `
packeg main

import (
	"fmt"
	"os"
	"strings"
)

{{.UserCode}}

func main() {
    input := strings.Join(os.Args[1:], " ")
    fmt.Print(Solve(input))
}
`

type TemplateData struct {
	userCode string
}

func GenerativeMainFile(userCode string) (string, error) {
	templ, err := template.New("main").Parse(mainTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := templ.Execute(&buf, TemplateData{userCode: userCode}); err != nil {
		return "", err
	}

	return buf.String(), err
}
