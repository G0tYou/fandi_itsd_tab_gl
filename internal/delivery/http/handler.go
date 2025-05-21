package http

import (
	d "app/domain"

	"github.com/labstack/echo"
)

type ResponseSuccess struct {
	Status  bool        `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type ResponseError struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// registers Employee routes with the provided Echo instance.
func NewUserHandler(e *echo.Echo, se d.ServiceEmployee) {
	handler := &EmployeeHandler{se}

	r := e.Group("/api/v1")

	//listing
	r.GET("/info/:nik", handler.GetEmployee)

	//editing
	r.PATCH("/update/:nik", handler.EditEmployee)

}
