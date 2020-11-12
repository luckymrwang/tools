package main

import "fmt"

// Loader defines a content loader
type Loader interface {
    Load(src string) string
}

// WebLoader is a web content loader
type WebLoader struct{}

// Load loads the content of a page
func (w *WebLoader) Load(src string) string {
    return fmt.Sprintf("I loaded this page %s", src)
}

func main() {
    webLoader := &WebLoader{}
    loadContent(webLoader)
}

func loadContent(loader Loader) {
    loader.Load("google.com")
}
