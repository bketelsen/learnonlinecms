package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Webinar struct {
	item.Item
	WebinarSlug     string `json:"webinar_slug"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Thumbnail       string `json:"thumbnail"`
	Price           int    `json:"price"`
	Status          string `json:"status"`
	Purchased       bool   `json:"purchased"`
	Date            string `json:"date"`
	DurationMinutes int    `json:"duration_minutes"`
}

func (c *Webinar) String() string {
	return c.WebinarSlug
}

func (c *Webinar) ItemSlug() string {
	return c.WebinarSlug
}

// MarshalEditor writes a buffer of html to edit a Webinar within the CMS
// and implements editor.Editable
func (c *Webinar) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(c,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Webinar field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:

		editor.Field{
			View: editor.Input("WebinarSlug", c, map[string]string{
				"label":       "Webinar Slug",
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
			View: editor.Input("Date", c, map[string]string{
				"label":       "Date",
				"type":        "text",
				"placeholder": "Enter the Date/Time here",
			}),
		},
		editor.Field{
			View: editor.Input("DurationMinutes", c, map[string]string{
				"label":       "Length (Minutes)",
				"type":        "text",
				"placeholder": "Enter the Duration here",
			}),
		},
		editor.Field{
			View: editor.Input("Status", c, map[string]string{
				"label":       "Status",
				"type":        "text",
				"placeholder": "Enter the Status here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Webinar editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Webinar"] = func() interface{} { return new(Webinar) }
}
