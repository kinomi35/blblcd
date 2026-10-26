package main

import (
	"blblcd/cli"
	"blblcd/model"
)

var (
	Version      string
	BuildTime    string
	Commit       string
	OriginAuthor = "Wangshihan"
	Maintainer   = "kinomi35"
)

func main() {
	cli.Execute(&model.Injection{
		Version:      Version,
		BuildTime:    BuildTime,
		Commit:       Commit,
		OriginAuthor: OriginAuthor,
		Maintainer:   Maintainer,
	})
}
