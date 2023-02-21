package templates

import "embed"

//go:embed *.gohtml site/layout/*.gohtml site/pages/*.gohtml
var FS embed.FS
