package domain

import (
	"context"
)

type Employee struct {
	NIK        int    `json:"nik"`
	Name       string `json:"nama"`
	BirthDate  string `json:"tanggal_lahir"`
	Gender     string `json:"jenis_kelamin"`
	Department string `json:"department"`
}

type ServiceEmployee interface {
	// Listing service
	// AddEmployee is a service to get a employee by nik to the repository
	GetEmployee(ctx context.Context, u *Employee) (int, error)

	// Editing service
	// EditEmployee is a service to editing data employee from the repository
	EditEmployee(ctx context.Context, u *Employee) (string, error)
}

type RepositoryEmployee interface {
	// Listing repository
	// ReadEmployee is a repository to read a employee by nik to the dataset
	ReadEmployee(ctx context.Context, nik string) (Employee, error)

	// Editing repository
	// UpdateEmployee is a repository to update data employee from the dataset
	UpdateEmployee(ctx context.Context, nik string) error
}
