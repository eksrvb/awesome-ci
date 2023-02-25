package cmd_test

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fullstack-devops/awesome-ci/internal/app/awesome-ci/cmd"

	"github.com/spf13/cobra/doc"
)

var docsPath = "../../../../docs/docs/CLI"

const fmTemplate = `---
title: "%s"
---
`

func TestCreateCobraDocs(t *testing.T) {
	if _, ok := os.LookupEnv("CI"); ok {
		t.Skip()
	}
	err := doc.GenMarkdownTreeCustom(cmd.RootCmd, docsPath, filePrepender, linkHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func filePrepender(filename string) string {
	name := filepath.Base(filename)
	base := strings.TrimSuffix(name, path.Ext(name))

	// replace "_" with " " for a nicer look
	base = strings.Replace(base, "_", " ", -1)
	if base != "awesome-ci" {
		base = strings.Replace(base, "awesome-ci", "_", -1)
	}
	return fmt.Sprintf(fmTemplate, base)
}

func linkHandler(name string) string {
	base := strings.TrimSuffix(name, path.Ext(name))
	return "./" + strings.ToLower(base)
}
