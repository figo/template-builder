package templating

import (
	"bytes"
	"github.com/figo/template-builder/pkg/config"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// WalkDir walks through directory, if directory name contains: {component}, replace it with "CompName"
// also replace variables in templates with actual variable values
func WalkDir(srcRoot string, dstRoot string, templateVars *config.TemplateVariables) {
	walkDir(srcRoot, dstRoot, templateVars)
}

func walkDir(srcRoot string, dstRoot string, templateVars *config.TemplateVariables) {
	files, err := ioutil.ReadDir(srcRoot)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		newFileName := strings.Replace(file.Name(), "{component}", templateVars.CompName, -1)
		newSrcRoot := srcRoot + "/" + file.Name()
		newDstRoot := dstRoot + "/" + newFileName

		if file.IsDir() {
			err = os.MkdirAll(newDstRoot, file.Mode().Perm())
			if err != nil {
				panic(err)
			}

			walkDir(newSrcRoot, newDstRoot, templateVars)
			continue
		}

		// if it is regular file, fill in variable values
		tmpl, err := template.New(file.Name()).ParseFiles(newSrcRoot)
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, templateVars)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(newDstRoot, buf.Bytes(), file.Mode().Perm())
		if err != nil {
			panic(err)
		}
	}
}
