package template

import (
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func LoadTemplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	siteResolver := newSiteResolver()
	r = Must(renderFilesAsTemplatesWithLayouts(r,
		"site.",
		Must(siteResolver.GetLayoutPath()),
		Must(siteResolver.GetPagesPath()),
		))
	adminResolver := newAdmnResolver()
	r = Must(renderFilesAsTemplatesWithLayouts(
		r,
		"admin.",
		Must(adminResolver.GetLayoutPath()),
		Must(adminResolver.GetPagesPath()),
		))
	r = Must(renderFilesAsTemplatesWithLayouts(
		r,
		"admin.",
		[]string{},
		[]string{
			"./templates/admin/auth/login.gohtml",
		}),
	)

	return r
}

func renderFilesAsTemplatesWithLayouts(renderer multitemplate.Renderer, prefix string, layouts []string, pages []string) (multitemplate.Renderer, error) {
	for _, page := range pages {
		LayoutCopy := make([]string, len(layouts))
		copy(LayoutCopy, layouts)
		files := append(LayoutCopy, page)
		pageName := prefix + filepath.Base(filepath.Dir(page)) + "." + filepath.Base(page)
		fmt.Println(pageName)
		renderer.AddFromFilesFuncs(pageName, templateFuncs, files...)

	}
	return renderer, nil
}

func Must[Result any](ok Result, err error) Result {

	if err != nil {
		panic(err)
	}
	return ok
}
