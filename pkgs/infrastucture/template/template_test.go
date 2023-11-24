package template

import (
	"errors"
	"testing"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/stretchr/testify/assert"
)

func TestGetOldWithDefault(t *testing.T) {
	inputs := map[string]interface{}{"key": "value"}
	assert.Equal(t, "value", GetOldWithDefault(&inputs, "key", "default"))
	assert.Equal(t, "default", GetOldWithDefault(&inputs, "nonexistent", "default"))
	assert.Equal(t, "default", GetOldWithDefault(nil, "key", "default"))
}

func TestTemplateFuncs(t *testing.T) {
	assert.Equal(t, time.Now().Year(), templateFuncs["getYear"].(func() int)())
	assert.Equal(t, "", templateFuncs["old"].(func(inputs *map[string]interface{}, input string) string)(nil, "key"))
	assert.Equal(t, "default", templateFuncs["oldWithDefault"].(func(inputs *map[string]interface{}, input string, defaultValue string) string)(nil, "key", "default"))
}

func TestTemplateResolver(t *testing.T) {
	resolver := newSiteResolver()
	matches, err := resolver.GetFromPath("*.gohtml")
	assert.Nil(t, err)
	assert.NotEmpty(t, matches)

	matches, err = resolver.GetLayoutPath()
	assert.Nil(t, err)
	assert.NotEmpty(t, matches)

	matches, err = resolver.GetPagesPath()
	assert.Nil(t, err)
	assert.NotEmpty(t, matches)
}

func TestNewResolver(t *testing.T) {
	resolver := newResolver("dir", "layout", "page").(*TemplateResolver)
	assert.Equal(t, "dir", resolver.templatesDir)
	assert.Equal(t, "layout", resolver.layoutPath)
	assert.Equal(t, "page", resolver.pagePath)
}

func TestLoadTemplates(t *testing.T) {
	renderer := LoadTemplates()
	assert.NotNil(t, renderer)
}

func TestRenderFilesAsTemplatesWithLayouts(t *testing.T) {
	renderer := multitemplate.NewRenderer()
	renderer, err := renderFilesAsTemplatesWithLayouts(renderer, "prefix", []string{"layout"}, []string{"page"})
	assert.Nil(t, err)
	assert.NotNil(t, renderer)
}

func TestMust(t *testing.T) {
	assert.Equal(t, "ok", Must("ok", nil))
	assert.Panics(t, func() { Must("ok", errors.New("error")) })
}
