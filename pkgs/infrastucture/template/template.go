package template

import (
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/mattn/go-zglob"
)

func LoadTemplates() multitemplate.Renderer {
	siteResolver := newSiteResolver()
	siteLayouts := Must(siteResolver.GetLayoutPath())
	sitePages := Must(siteResolver.GetPagesPath())
	r := Must(renderFilesAsTemplatesWithLayouts(multitemplate.NewRenderer(), "site.", siteLayouts, sitePages))

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++=")
	adminLayouts, err := filepath.Glob("./templates/admin/layout/*.gohtml")
	if err != nil {
		panic(err.Error())
	}

	admins, err := zglob.Glob(`./templates/admin/pages/**/*.gohtml`)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("========================adminLayouts===================")
	fmt.Println(adminLayouts)

	// // Generate our templates map from our adminLayouts/ and admins/ directories
	for _, admin := range admins {
		layoutCopy := make([]string, len(adminLayouts))
		copy(layoutCopy, adminLayouts)
		fmt.Println("========================layoutCopy===================")
		fmt.Println(layoutCopy)
		files := append(layoutCopy, admin)
		fmt.Println("========================admin page===================")
		fmt.Println(admin)

		fmt.Println("========================files(layout+ pages)===================")
		fmt.Println("================================================")
		dir_name := filepath.Base(filepath.Dir(admin))
		fmt.Println("admin."+dir_name+"."+filepath.Base(admin))
		r.AddFromFilesFuncs("admin."+dir_name+"."+filepath.Base(admin), templateFuncs, files...)
	}
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
