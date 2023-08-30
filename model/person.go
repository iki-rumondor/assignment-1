package model


type Person struct{
	Name string `json:"student_name"`
	Code string `json:"student_code"`
	Address string `json:"student_address"`
	Occupation string `json:"student_occupation"`
	JoiningReason string `json:"joining_reason"`
}

type People struct{
	People []Person `json:"participants"`
}

