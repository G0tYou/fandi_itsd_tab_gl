package service

import (
	d "app/domain"
	"context"
	"encoding/json"
	"io/fs"
	"os"
)

type serviceEmployee struct {
}

// NewService creates an service with the necessary dependencies
func NewServiceUser() d.ServiceEmployee {
	return &serviceEmployee{}
}

// Write the listing service below
// GetEmployee is a method to get a employee from dataset
func (s *serviceEmployee) GetEmployee(ctx context.Context, nik string) (d.Employee, error) {
	var e d.Employee

	es, err := readEmployeesJsonByNik(nik)
	if err != nil {
		return e, err
	}

	for _, employee := range es {
		if employee.NIK == nik {
			e = employee
			break
		}
	}

	return e, nil
}

// Write the editing service below
// EditEmployee is a method to update data employee to the dataset
func (s *serviceEmployee) EditEmployee(ctx context.Context, nik string, ee d.Employee) (d.Employee, error) {
	var e d.Employee

	es, err := readEmployeesJsonByNik(nik)
	if err != nil {
		return e, err
	}

	for i, employee := range es {
		if employee.NIK == nik {
			if ee.Name != "" {
				e.Name = employee.Name
				es[i].Name = ee.Name
			}
			if ee.BirthDate != "" {
				es[i].BirthDate = ee.BirthDate
			}
			if ee.Gender != "" {
				es[i].Gender = ee.Gender
			}
			if ee.Department != "" {
				es[i].Department = ee.Department
			}

			e.NIK = employee.NIK

			data, err := json.MarshalIndent(es, "", "  ")
			if err != nil {
				return e, err
			}

			os.WriteFile("employee.json", data, fs.ModePerm)
			break
		}
	}

	return e, err
}

func readEmployeesJsonByNik(nik string) ([]d.Employee, error) {
	var es []d.Employee

	data, err := os.ReadFile("employee.json")
	if err != nil {
		return es, err
	}

	if err := json.Unmarshal(data, &es); err != nil {
		return es, err
	}

	return es, nil
}
