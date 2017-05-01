package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Post struct {
	item.Item

	Title       string   `json:"title"`
	BlogSlug    string   `json:"blog_slug"`
	HeaderImage string   `json:"header_image"`
	Lede        string   `json:"lede"`
	Body        string   `json:"body"`
	Tags        []string `json:"tags"`
}

func (p *Post) String() string {
	return p.BlogSlug
}

func (p *Post) ItemSlug() string {
	return p.BlogSlug
}

// MarshalEditor writes a buffer of html to edit a Post within the CMS
// and implements editor.Editable
func (p *Post) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Post field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", p, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Lede", p, map[string]string{
				"label":       "Lede",
				"type":        "text",
				"placeholder": "Enter the Lede here",
			}),
		},
		editor.Field{
			View: editor.Input("BlogSlug", p, map[string]string{
				"label":       "BlogSlug",
				"type":        "text",
				"placeholder": "Enter the BlogSlug here",
			}),
		},
		editor.Field{
			View: editor.File("HeaderImage", p, map[string]string{
				"label":       "Header Image",
				"type":        "text",
				"placeholder": "Enter the Header Image here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Body", p, map[string]string{
				"label":       "Body",
				"type":        "text",
				"placeholder": "Enter the Body here",
			}),
		},
		editor.Field{
			View: editor.Tags("Tags", p, map[string]string{
				"label":       "Tags",
				"type":        "text",
				"placeholder": "Enter the Tags here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Post editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Post"] = func() interface{} { return new(Post) }
}
