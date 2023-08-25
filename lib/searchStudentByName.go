package lib

import (
	"assignment-1/data"
	"assignment-1/entity"
	"strings"
	"errors"
)

func SearchStudentByName(name string) (entity.Student, error)  {
	for _, student := range data.Students {
		if strings.ToLower(student.Name) == strings.ToLower(name) {
			return student, nil
		}
	}

	return entity.Student{}, errors.New("Data tidak ditemukan")
}