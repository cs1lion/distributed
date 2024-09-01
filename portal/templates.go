package portal

import "html/template"

var rootTemplate *template.Template

func ImportTemplate() error {
	var err error
	rootTemplate, err = template.ParseFiles(
		"../../portal/student.html",
		"../../portal/students.html")
	if err != nil {
		return err
	}
	return nil
}
