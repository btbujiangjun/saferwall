// Copyright 2018 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package exiftool

import (
	"log"
	"strings"

	"github.com/saferwall/saferwall/pkg/utils"
	strcase "github.com/stoewer/go-strcase"
)

const (
	// Command to invoke exiftool scanner
	Command = "exiftool"
)

// Scan a file using TRiD Scanner
// This will execute trid command line tool and read the output from stdout
func Scan(FilePath string) (map[string]string, error) {

	Args := []string{FilePath}

	output, err := utils.ExecCommand(Command, Args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return ParseOutput(output), nil
}

// ParseOutput convert exiftool output into JSON
func ParseOutput(exifout string) map[string]string {

	var ignoreTags = []string{
		"Directory",
		"File Name",
		"File Permissions",
	}

	lines := strings.Split(exifout, "\n")

	if utils.StringInSlice("File not found", lines) {
		return nil
	}

	datas := make(map[string]string, len(lines))

	for _, line := range lines {
		keyvalue := strings.Split(line, ":")
		if len(keyvalue) != 2 {
			continue
		}
		if !utils.StringInSlice(strings.TrimSpace(keyvalue[0]), ignoreTags) {
			datas[strings.TrimSpace(strcase.UpperCamelCase(keyvalue[0]))] = strings.TrimSpace(keyvalue[1])
		}
	}

	return datas
}