package blog_posts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blog_posts "github.com/Mohamed-Mayallo/go-learning/15-fs"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fsMap := fstest.MapFS{
		"hello1.md": {Data: []byte(firstBody)},
		"hello2.md": {Data: []byte(secondBody)},
	}

	posts, err := blog_posts.NewPostsFromFS(fsMap)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fsMap) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fsMap))
	}

	got := posts[0]
	want := blog_posts.Post{Title: "Post 1", Description: "Description 1", Body: `Hello
World`, Tags: []string{"tdd", "go"}}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got blog_posts.Post, want blog_posts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
