package models

type AdminStatus int

const (
	AdminStatusDel    AdminStatus = iota - 1
	AdminStatusNormal AdminStatus = iota
	AdminStatusDisable
)
