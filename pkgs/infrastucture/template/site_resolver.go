package template

import "path/filepath"

const templateDir = "./templates"
const siteLayout = "/site/layout/*.gohtml"
const sitePages = "/site/pages/*.gohtml"

type TemplateResolver struct {
	templatesDir string
	layoutPath   string
	pagePath     string
}

func (loader TemplateResolver) GetFromPath(pattern string) (matches []string, err error) {
	return filepath.Glob(loader.templatesDir + pattern)
}

func (loader TemplateResolver) GetLayoutPath() (matches []string, err error) {

	return loader.GetFromPath(loader.layoutPath)
}

func (loader TemplateResolver) GetPagesPath() (matches []string, err error) {
	return loader.GetFromPath(loader.pagePath)
}

func newResolver(templateDir string, layoutPath string, pagePath string) *TemplateResolver {

	return &TemplateResolver{
		templatesDir: templateDir,
		layoutPath:   layoutPath,
		pagePath:     pagePath,
	}
}
func newSiteResolver() *TemplateResolver {

	return newResolver(
		templateDir,
		siteLayout,
		sitePages,
	)
}
