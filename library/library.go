package library

import "fmt"

func iniprivate(){
	fmt.Println("Ini Private")
}

func Inipublic(){
	fmt.Println("Ini Public")
}

func Aksesprivate(){
	iniprivate()
}