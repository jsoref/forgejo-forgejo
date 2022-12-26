// Copyright 2022 The Forgejo Authors c/o Codeberg e.V.. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"bufio"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

const (
	trimPrefix   = "gitea_"
	sourceFolder = "options/locales/"
)

// returns list of locales, still containing the file extension!
func generate_locale_list() []string {
	localeFiles, _ := os.ReadDir(sourceFolder)
	locales := []string{}
	for _, localeFile := range localeFiles {
		if !localeFile.IsDir() && strings.HasPrefix(localeFile.Name(), trimPrefix) {
			locales = append(locales, strings.TrimPrefix(localeFile.Name(), trimPrefix))
		}
	}
	return locales
}

// replace all occurrences of Gitea with Forgejo
func renameGiteaForgejo(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	out := make([]byte, 0, 1024)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			out = append(out, []byte("\n"+line+"\n")...)
		} else {
			out = append(out, []byte(strings.NewReplacer(
				"Gitea", "Forgejo",
				"https://docs.gitea.io/en-us/install-from-binary/", "https://forgejo.org/download",
				"https://github.com/go-gitea/gitea/tree/master/docker", "https://forgejo.org/download",
				"https://docs.gitea.io/en-us/install-from-package/", "https://forgejo.org/download",
				"https://code.gitea.io/gitea", "https://forgejo.org/download",
				"code.gitea.io/gitea", "Forgejo",
				`<a href="https://github.com/go-gitea/gitea/issues" target="_blank">GitHub</a>`, `<a href="https://codeberg.org/forgejo/forgejo/issues" target="_blank">Codeberg</a>`,
				"https://github.com/go-gitea/gitea", "https://codeberg.org/forgejo/forgejo",
				"https://blog.gitea.io", "https://forgejo.org/news",
			).Replace(line)+"\n")...)
		}
	}
	file.Close()
	return out
}

func main() {
	locales := generate_locale_list()
	var err error
	var localeFile *ini.File
	for _, locale := range locales {
		giteaLocale := sourceFolder + "gitea_" + locale
		localeFile, err = ini.Load(giteaLocale, renameGiteaForgejo(giteaLocale))
		if err != nil {
			panic(err)
		}
		err = localeFile.SaveTo("options/locale/locale_" + locale)
		if err != nil {
			panic(err)
		}
	}
}
