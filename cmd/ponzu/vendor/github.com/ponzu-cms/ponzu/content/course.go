package content

import (
	"fmt"

	"github.com/bosssauce/reference"
	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Course struct {
	item.Item
	CourseSlug  string   `json:"course_slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	Price       int      `json:"price"`
	Status      string   `json:"status"`
	Modules     []string `json:"modules"`
	Purchased   bool     `json:"purchased"`
	Level       string   `json:"level"`
	Topic       string   `json:"topic"`
	Outline     string   `json:"outline"`
}

func (c *Course) String() string {
	return c.CourseSlug
}

func (c *Course) ItemSlug() string {
	return c.CourseSlug
}
func (c *Course) Push() []string {
	return []string{"modules"}
}

// MarshalEditor writes a buffer of html to edit a Course within the CMS
// and implements editor.Editable
func (c *Course) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(c,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Course field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:

		editor.Field{
			View: editor.Input("CourseSlug", c, map[string]string{
				"label":       "Course Slug",
				"type":        "text",
				"placeholder": "Enter the Slug here",
			}),
		},
		editor.Field{
			View: editor.Input("Title", c, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Description", c, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Outline", c, map[string]string{
				"label":       "Outline (markdown)",
				"type":        "text",
				"placeholder": "Enter the Outline (markdown)here",
			}),
		},
		editor.Field{
			View: editor.Input("Topic", c, map[string]string{
				"label":       "Topic",
				"type":        "text",
				"placeholder": "Enter the Topic here",
			}),
		},
		editor.Field{
			View: editor.File("Thumbnail", c, map[string]string{
				"label":       "Thumbnail",
				"type":        "text",
				"placeholder": "Enter the Thumbnail here",
			}),
		},
		editor.Field{
			View: editor.Input("Price", c, map[string]string{
				"label":       "Price",
				"type":        "text",
				"placeholder": "Enter the Price here",
			}),
		},
		editor.Field{
			View: editor.Input("Status", c, map[string]string{
				"label":       "Status",
				"type":        "text",
				"placeholder": "Enter the Status here",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Modules", c, map[string]string{
				"label":       "Module",
				"placeholder": "Enter the module here",
			},
				"Module",
				"{{ .module_slug }}",
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Course editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Course"] = func() interface{} { return new(Course) }
}
