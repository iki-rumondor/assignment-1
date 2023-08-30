package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/iki-rumondor/assignment-1/model"
)

func main() {

	person, res := getPersonByCode(&os.Args[1])

	printPersonInformation(person, &res)

}

// Fungsi ini akan mengecek person berdasarkan code tertentu didalam struct people kemudian akan mengembalikan pointer dari struct person.
func getPersonByCode(code *string) (*model.Person, bool) {
	people := convertDataToModel()

	for _, person := range people.People {
		if person.Code == *code {
			return &person, true
		}
	}

	return &model.Person{}, false

}

// Fungsi ini akan mengkonversi data JSON ke dalam struct People di dalam folder model. Fungsi ini akan mengembalikan pointer dari struct people.
func convertDataToModel() (people *model.People) {
	jsonPath := filepath.Join("data", "participants.json")
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &people); err != nil {
		log.Fatal(err)
	}

	return
}

// Fungsi ini akan mengecek kondisi yang diberikan, apabila true akan mencetak informasi dari person yang dimasukkan, apabila false akan mencetak pesan data tidak ditemukan.
func printPersonInformation(person *model.Person, conditon *bool) {
	switch *conditon {
	case true:
		fmt.Println("ID :", person.Code)
		fmt.Println("nama :", person.Name)
		fmt.Println("alamat :", person.Address)
		fmt.Println("pekerjaan :", person.Occupation)
		fmt.Println("alasan :", person.JoiningReason)
	default:
		fmt.Println("Not Found Person")
	}
}
