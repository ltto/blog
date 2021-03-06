package tools

import (
	"blog/system"
	"path"
	"strings"
)

type String string

func (s *String) TrimLeft(left string) *String {
	*s = String(strings.TrimLeft(string(*s), left))
	return s
}
func BaseAdd(add string) string {
	return path.Join("/", system.Conf.Server.Base, add)
}
