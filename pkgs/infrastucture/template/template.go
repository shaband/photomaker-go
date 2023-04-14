package template

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func LoadTemplates() multitemplate.Renderer {
	siteResolver := newSiteResolver()
	siteLayouts := Must(siteResolver.GetLayoutPath())
	sitePages := Must(siteResolver.GetPagesPath())
	r := Must(renderFilesAsTemplatesWithLayouts(multitemplate.NewRenderer(), "site.", siteLayouts, sitePages))

	// adminLayouts, err := filepath.Glob(templatesDir + "/layouts/admin/*.gohtml")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// admins, err := filepath.Glob(templatesDir + "/admins/*.html")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // Generate our templates map from our adminLayouts/ and admins/ directories
	// for _, admin := range admins {
	// 	layoutCopy := make([]string, len(adminLayouts))
	// 	copy(layoutCopy, adminLayouts)
	// 	files := append(layoutCopy, admin)
	// 	r.AddFromFiles(filepath.Base(admin), files...)
	// }
	return r
}

func renderFilesAsTemplatesWithLayouts(renderer multitemplate.Renderer, prefix string, layouts []string, pages []string) (multitemplate.Renderer, error) {
	for _, page := range pages {
		LayoutCopy := make([]string, len(layouts))
		copy(LayoutCopy, layouts)
		files := append(LayoutCopy, page)
		renderer.AddFromFilesFuncs(prefix+filepath.Base(page), templateFuncs, files...)

	}
	return renderer, nil
}

func Must[Result any](ok Result, err error) Result {

	if err != nil {
		panic(err)
	}
	return ok
}
