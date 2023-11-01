package v1

import (
	"HeadHunter/internal/controller"
	"HeadHunter/pkg/employmentService"
	"context"
)

type GrpcServer struct {
	cnt *controller.Controller
}

func NewGrpcServer(cnt *controller.Controller) *GrpcServer {
	// constructor for api
	return &GrpcServer{cnt: cnt}
}

func (s GrpcServer) RegisterCompany(_ context.Context, request *employmentService.CompanyRequest) (*employmentService.BasicResponse, error) {
	// registering company
	message, err := s.cnt.AddCompanyByName(request)
	response := employmentService.BasicResponse{Message: message}

	return &response, err

}

func (s GrpcServer) RegisterEmployee(_ context.Context, request *employmentService.EmployeeRequest) (*employmentService.BasicResponse, error) {
	// registering employee
	message, err := s.cnt.AddEmployee(request)
	response := employmentService.BasicResponse{Message: message}

	return &response, err
}

func (s GrpcServer) CreateNewDepartment(_ context.Context, request *employmentService.DepartmentRequest) (*employmentService.BasicResponse, error) {
	// creating department for company
	message, err := s.cnt.AddDepartmentToCompany(request)
	response := employmentService.BasicResponse{Message: message}

	return &response, err
}

func (s GrpcServer) EmployPerson(_ context.Context, request *employmentService.EmploymentRequest) (*employmentService.BasicResponse, error) {
	// employing person
	message, err := s.cnt.EmployPerson(request)
	response := employmentService.BasicResponse{Message: message}

	return &response, err
}
