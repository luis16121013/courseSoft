package models

type User struct {
	Id int64 `json:"id"`
	Nombre string `json:"nombre"`
	Apellido int64 `json:"apellido"`
	Dni string `json:"dni"`
	TypeUser string `json:"typeUser"`
}
