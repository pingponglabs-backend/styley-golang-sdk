package models

import (
	"net/url"
	"path"
)

func generateNamePath(name string) string {
	return path.Join("/api/v1/models/", "name", url.PathEscape(name))
}
