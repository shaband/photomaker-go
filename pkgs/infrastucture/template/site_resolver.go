package template

import "github.com/mattn/go-zglob"

const templateDir = "./templates"
const siteLayout = "/site/layout/*.gohtml"
const sitePages = "/site/pages/*.gohtml"
const adminLayout = "/admin/layout/*.gohtml"
const adminPages = "/admin/pages/**/*.gohtml"

type TemplateResolver struct {
	templatesDir string
	layoutPath   string
	pagePath     string
}

func (loader TemplateResolver) GetFromPath(pattern string) (matches []string, err error) {
	return zglob.Glob(loader.templatesDir + pattern)
}

func (loader TemplateResolver) GetLayoutPath() (matches []string, err error) {

	return loader.GetFromPath(loader.layoutPath)
}

func (loader TemplateResolver) GetPagesPath() (matches []string, err error) {
	return loader.GetFromPath(loader.pagePath)
}

func newResolver(templateDir string, layoutPath string, pagePath string) resolver {

	return &TemplateResolver{
		templatesDir: templateDir,
		layoutPath:   layoutPath,
		pagePath:     pagePath,
	}
}
func newSiteResolver() resolver {

	return newResolver(
		templateDir,
		siteLayout,
		sitePages,
	)
}

func newAdmnResolver() resolver {

	return newResolver(
		templateDir,
		adminLayout,
		adminPages,
	)

}
