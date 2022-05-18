package data

import (
	"assignment1/model"
	"fmt"
	"strconv"
)

func SearchStudents(id string) {
	var temp int
	nimId, _ := strconv.Atoi(id)

	students := []model.Students{
		{
			Nim: 4,
			Profile: model.Profile{
				Nama:      "Alvin",
				Alamat:    "Yogyakarta",
				Pekerjaan: "Mahasiswa",
				Alasan:    "Cepat",
			},
		},
		{
			Nim: 1,
			Profile: model.Profile{
				Nama:      "Saut",
				Alamat:    "Samosir",
				Pekerjaan: "Mahasiswa",
				Alasan:    "Biar mudah",
			},
		},
		{
			Nim: 2,
			Profile: model.Profile{
				Nama:      "Bongson",
				Alamat:    "Parsoburan",
				Pekerjaan: "Mahasiswa",
				Alasan:    "Efisien",
			},
		},
		{
			Nim: 3,
			Profile: model.Profile{
				Nama:      "Medi",
				Alamat:    "Medan",
				Pekerjaan: "Mahasiswa",
				Alasan:    "Gabut",
			},
		},
		{
			Nim: 5,
			Profile: model.Profile{
				Nama:      "Joe",
				Alamat:    "Jakarta",
				Pekerjaan: "Mahasiswa",
				Alasan:    "Felksibel",
			},
		},
	}

	for key, data := range students {
		// fmt.Println(key)
		if nimId == data.Nim {
			temp = key
		}
		// fmt.Println(temp)
	}

	if temp > 0 {
		fmt.Println("Nama \t:", students[temp].Profile.Nama)
		fmt.Println("Alamat \t:", students[temp].Profile.Alamat)
		fmt.Println("Pekerjaan \t:", students[temp].Profile.Pekerjaan)
		fmt.Println("Alasan \t:", students[temp].Profile.Alasan)
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}
