package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Module struct {
	item.Item

	ModuleSlug   string `json:"module_slug"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Thumbnail    string `json:"thumbnail"`
	Exercise     string `json:"exercise"`
	ExercisePath string `json:"exercise_path"`
	VideoId      int    `json:"video_id"`
	Minutes      int    `json:"minutes"`
	Order        int    `json:"order"`
}

func (c *Module) String() string {
	return c.ModuleSlug
}

func (c *Module) ItemSlug() string {
	return c.ModuleSlug
}

// MarshalEditor writes a buffer of html to edit a Module within the CMS
// and implements editor.Editable
func (m *Module) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(m,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Module field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("ModuleSlug", m, map[string]string{
				"label":       "ModuleSlug",
				"type":        "text",
				"placeholder": "Enter the ModuleSlug here",
			}),
		},
		editor.Field{
			View: editor.Input("Order", m, map[string]string{
				"label":       "Order",
				"type":        "text",
				"placeholder": "Enter the Order here",
			}),
		},
		editor.Field{
			View: editor.Input("Title", m, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Description", m, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.File("Thumbnail", m, map[string]string{
				"label":       "Thumbnail",
				"type":        "text",
				"placeholder": "Enter the Thumbnail here",
			}),
		},
		editor.Field{
			View: editor.Input("Minutes", m, map[string]string{
				"label":       "Minutes",
				"type":        "text",
				"placeholder": "Enter the length in Minutes here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Exercise", m, map[string]string{
				"label":       "Exercise (Markdown Format)",
				"placeholder": "Enter the Exercise here",
			}),
		},
		editor.Field{
			View: editor.Input("ExercisePath", m, map[string]string{
				"label":       "ExercisePath",
				"type":        "text",
				"placeholder": "Enter the ExercisePath here",
			}),
		},
		editor.Field{
			View: editor.Input("VideoId", m, map[string]string{
				"label":       "Video ID",
				"type":        "text",
				"placeholder": "Enter the Video ID here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Module editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Module"] = func() interface{} { return new(Module) }
}
