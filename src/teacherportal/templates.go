package teacherportal

import "html/template"

var rootTemplate *template.Template

func ImportTemplates() error {
	var err error
	rootTemplate, err = template.ParseFiles(
		"src/teacherportal/students.gohtml",
		"src/teacherportal/student.gohtml",
	)

	if err != nil {
		return err
	}
	return nil
}
