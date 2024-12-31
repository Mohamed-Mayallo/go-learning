package renderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	blog_posts "github.com/Mohamed-Mayallo/go-learning/15-fs"
)

//go:embed "templates/*"
var postTemplate embed.FS

func Render(w io.Writer, p blog_posts.Post) error {
	tmp, err := template.New("blog").Funcs(template.FuncMap{
		"slugify": func(title string) string {
			return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		},
	}).ParseFS(postTemplate, "templates/*.gohtml")

	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}
