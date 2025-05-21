package domain

import (
	"context"
)

type Employee struct {
	NIK        string `json:"nik"`
	Name       string `json:"nama"`
	BirthDate  string `json:"tanggal_lahir"`
	Gender     string `json:"jenis_kelamin"`
	Department string `json:"departemen"`
}

type ServiceEmployee interface {
	// Listing service
	// AddEmployee is a service to get a employee by nik to the repository
	GetEmployee(ctx context.Context, nik string) (Employee, error)

	// Editing service
	// EditEmployee is a service to editing data employee from the repository
	EditEmployee(ctx context.Context, nik string, e Employee) (Employee, error)
}
