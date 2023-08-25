package lib

import (
	"assignment-1/entity"
	"fmt"
)

func PrintOutput(student entity.Student) {
	fmt.Printf("Id \t\t : %d \n", student.Id)
	fmt.Printf("Nama \t\t : %s \n", student.Name)
	fmt.Printf("Alamat \t\t : %s \n", student.Address)
	fmt.Printf("Pekerjaan \t : %s \n", student.Job)
	fmt.Printf("Alasan \t\t : %s \n", student.Reason)
}