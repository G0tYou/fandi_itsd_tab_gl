package http

import (
	conf "app/config"
	d "app/domain"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type EmployeeHandler struct {
	se d.ServiceEmployee
}

// GetEmployee  handles the request to get a employee by NIK
func (h *EmployeeHandler) GetEmployee(c echo.Context) error {
	nik := c.Param("nik")

	employee, err := h.se.GetEmployee(c.Request().Context(), nik)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Status: false, Message: err.Error()})
	} else if employee.NIK == "" {
		return c.JSON(http.StatusOK, ResponseError{Status: false, Message: fmt.Sprintf(conf.ErrorMessageNIKNotFound, nik)})
	}

	return c.JSON(http.StatusOK, ResponseSuccess{Status: true, Data: employee})
}

// EditEmployee  handles the request to edit a employee by NIK
func (h *EmployeeHandler) EditEmployee(c echo.Context) error {
	nik := c.Param("nik")

	var e *d.Employee

	if err := c.Bind(&e); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, ResponseError{Message: http.StatusText(http.StatusBadRequest)})
	}

	employee, err := h.se.EditEmployee(c.Request().Context(), nik, *e)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Status: false, Message: err.Error()})
	} else if employee.NIK == "" {
		return c.JSON(http.StatusOK, ResponseError{Status: false, Message: fmt.Sprintf(conf.ErrorMessageNIKNotFound, nik)})
	}

	return c.JSON(http.StatusCreated, ResponseSuccess{Status: true, Message: fmt.Sprintf(conf.SuccessMessageUpdateEmployee, employee.Name, e.Name)})
}
