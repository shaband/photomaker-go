package template

import (
	"path/filepath"
	"testing"

	"github.com/gin-contrib/multitemplate"
)

func TestLoadTemplates(t *testing.T) {
	// Create a new instance of multitemplate.Renderer
	renderer := multitemplate.NewRenderer()

	// Call the LoadTemplates function
	result := LoadTemplates()

	// Assert that the returned renderer is not nil
	if result == nil {
		t.Errorf("LoadTemplates returned nil renderer")
	}
}

func TestRenderFilesAsTemplatesWithLayouts(t *testing.T) {
	// Create a new instance of multitemplate.Renderer
	renderer := multitemplate.NewRenderer()

	// Define test data
	prefix := "test."
	layouts := []string{"layout1", "layout2"}
	pages := []string{"page1", "page2"}

	// Call the renderFilesAsTemplatesWithLayouts function
	result, err := renderFilesAsTemplatesWithLayouts(renderer, prefix, layouts, pages)

	// Assert that the returned renderer is not nil
	if result == nil {
		t.Errorf("renderFilesAsTemplatesWithLayouts returned nil renderer")
	}

	// Assert that the error is nil
	if err != nil {
		t.Errorf("renderFilesAsTemplatesWithLayouts returned unexpected error: %v", err)
	}

	// Assert that the templates are added to the renderer correctly
	expectedPageName := "test.page1"
	expectedFiles := []string{"layout1", "layout2", "page1"}
	template := result.Lookup(expectedPageName)
	if template == nil {
		t.Errorf("renderFilesAsTemplatesWithLayouts did not add template with page name %s", expectedPageName)
	} else {
		// Assert that the template contains the expected files
		for _, file := range expectedFiles {
			if !contains(template.Files, file) {
				t.Errorf("renderFilesAsTemplatesWithLayouts did not add file %s to template %s", file, expectedPageName)
			}
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
