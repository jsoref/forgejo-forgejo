// Copyright 2022 The Forgejo Authors c/o Codeberg e.V.. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"bufio"
	"os"
	"regexp"
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

	replacer := strings.NewReplacer(
		"Gitea", "Forgejo",
		"https://docs.gitea.io/en-us/install-from-binary/", "https://forgejo.org/download/#installation-from-binary",
		"https://github.com/go-gitea/gitea/tree/master/docker", "https://forgejo.org/download/#container-image",
		"https://docs.gitea.io/en-us/install-from-package/", "https://forgejo.org/download",
		"https://code.gitea.io/gitea", "https://forgejo.org/download",
		"code.gitea.io/gitea", "Forgejo",
		`<a href="https://github.com/go-gitea/gitea/issues" target="_blank">GitHub</a>`, `<a href="https://codeberg.org/forgejo/forgejo/issues" target="_blank">Codeberg</a>`,
		"https://github.com/go-gitea/gitea", "https://codeberg.org/forgejo/forgejo",
		"https://blog.gitea.io", "https://forgejo.org/news",
		"https://docs.gitea.io/en-us/protected-tags/", "https://forgejo.org/docs/latest/user/protection/#protected-tags",
		"https://docs.gitea.io/en-us/webhooks/", "https://forgejo.org/docs/latest/user/webhooks/",
	)

	out := make([]byte, 0, 1024)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "license_desc=") {
			line = strings.Replace(line, "GitHub", "Forgejo", 1)
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			out = append(out, []byte("\n"+line+"\n")...)
		} else if strings.HasPrefix(line, "settings.web_hook_name_gitea") {
			out = append(out, []byte("\n"+line+"\n")...)
			out = append(out, []byte("settings.web_hook_name_forgejo = Forgejo\n")...)
		} else if strings.HasPrefix(line, "migrate.gitea.description") {
			re := regexp.MustCompile(`(.*Gitea)`)
			out = append(out, []byte(re.ReplaceAllString(line, "${1}/Forgejo")+"\n")...)
		} else {
			out = append(out, []byte(replacer.Replace(line)+"\n")...)
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
		localeFile, err = ini.LoadSources(ini.LoadOptions{
			IgnoreInlineComment: true,
		}, giteaLocale, renameGiteaForgejo(giteaLocale))
		if err != nil {
			panic(err)
		}
		err = localeFile.SaveTo("options/locale/locale_" + locale)
		if err != nil {
			panic(err)
		}
	}
}
