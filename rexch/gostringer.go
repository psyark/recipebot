package rexch

import "fmt"

// TODO: 実装
func (r Recipe) GoString() string {
	return fmt.Sprintf(`{
	"Title": %q,
	"Image": %q,
	"Servings": %v,
	"Ingredients": %#v,
	"Instructions": %#v,
}`, r.Title, r.Image, r.Servings, r.Ingredients, r.Instructions)
}
