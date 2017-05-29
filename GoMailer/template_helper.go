package main

import (
	"bytes"
	"html/template"
)

// ParseTemplate returns parsed template with data in string format
// If there is an error, it will return response with error data
func ParseTemplate(templateFileName string, data interface{}) (content string, err error) {

	// ParseFiles creates a new 	Template and parses the template definitions from
	// the named files. The returned template's name will have the (base) name and
	// (parsed) contents of the first file. There must be at least one file.
	// If an error occurs, parsing stops and the returned *Template is nil.
	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	// The zero value for Buffer is an empty buffer ready to use.
	buf := new(bytes.Buffer)

	// Execute applies a parsed template to the specified data object,
	// writing the output to wr.
	// If an error occurs executing the template or writing its output,
	// execution stops, but partial results may already have been written to
	// the output writer.
	// A template may be executed safely in parallel.
	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
