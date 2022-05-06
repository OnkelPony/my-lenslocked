package templates

import "embed"

//go:embed *
var FS embed.FS

//go:embed start.gohtml
var StartMessage string
