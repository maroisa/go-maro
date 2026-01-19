package web

import "embed"

//go:embed templates/*.html
var TemplatesFS embed.FS

//go:embed assets/*
var AssetsFS embed.FS
