package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"text/template"
)

// Context for templates
type Context struct {
}

// Env - Map environment variables for use in a template
func (c *Context) Env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		tmpl, err := template.New("template").Option("missingkey=error").Parse(s.Text())
		if err != nil {
			log.Fatalf("Line %q: %v\n", s.Text(), err)
		}

		if err := tmpl.Execute(os.Stdout, &Context{}); err != nil {
			panic(err)
		}
		os.Stdout.Write([]byte("\n"))
	}
}
