package renderer_test

import (
	"bytes"
	"testing"

	blog_posts "github.com/Mohamed-Mayallo/go-learning/15-fs"
	renderer "github.com/Mohamed-Mayallo/go-learning/16-template"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRenderer(t *testing.T) {
	buf := bytes.Buffer{}

	post := blog_posts.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	if err := renderer.Render(&buf, post); err != nil {
		t.Fatal(err)
	}

	approvals.VerifyString(t, buf.String())
}
