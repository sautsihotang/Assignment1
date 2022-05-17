package model

type Students struct {
	Nim     int
	Profile Profile
}

type Profile struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}
