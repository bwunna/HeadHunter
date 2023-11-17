package v1

import (
	"HeadHunter/internal/controller"
	"HeadHunter/internal/services/userService"
	"HeadHunter/pkg/employmentService"
	"context"
)

// api

type GrpcServer struct {
	cnt         *controller.Controller
	usersClient *userService.Client
}

// constructor for api

func NewGrpcServer(cnt *controller.Controller, usersClient *userService.Client) *GrpcServer {
	return &GrpcServer{
		cnt:         cnt,
		usersClient: usersClient,
	}
}

// sending request in controller to register company by its name

func (s GrpcServer) RegisterCompany(_ context.Context, request *employmentService.CompanyRequest) (*employmentService.BasicResponse, error) {

	err := s.cnt.AddCompanyByName(request.Name)

	if err != nil {
		return &employmentService.BasicResponse{Message: MessageNotExecuted, Code: CodeNotExecuted}, err
	}

	return &employmentService.BasicResponse{Message: MessageOK, Code: CodeOK}, err
}

// sending request in controller to register new employee

func (s GrpcServer) RegisterEmployee(_ context.Context, request *employmentService.EmployeeRequest) (*employmentService.BasicResponse, error) {

	err := s.cnt.AddEmployee(request)
	if err != nil {
		return &employmentService.BasicResponse{
			Code:    int32(CodeNotExecuted),
			Message: MessageNotExecuted,
		}, err
	}

	return &employmentService.BasicResponse{
		Message: MessageOK,
		Code:    CodeOK,
	}, err
}

// sending request to controller to add the company

func (s GrpcServer) CreateNewDepartment(_ context.Context, request *employmentService.DepartmentRequest) (*employmentService.BasicResponse, error) {
	// creating department for company
	err := s.cnt.AddDepartmentToCompany(request)
	if err != nil {
		return &employmentService.BasicResponse{Message: MessageNotExecuted, Code: CodeNotExecuted}, err
	}

	return &employmentService.BasicResponse{Message: MessageOK, Code: CodeOK}, nil
}

// sending request to controller to update salary of the employee

func (s GrpcServer) UpdateEmployeeSalaryByEmail(_ context.Context, request *employmentService.ByEmailRequest) (*employmentService.BasicResponse, error) {
	err := s.cnt.UpdateSalaryByEmail(request.Email)
	if err != nil {
		return &employmentService.BasicResponse{Message: MessageNotExecuted, Code: CodeNotExecuted}, err
	}

	ctx := context.Background()
	info, err := s.cnt.GetEmployeeInfoByEmail(request.Email)
	if err != nil {
		return &employmentService.BasicResponse{Message: MessageNOTFOUND, Code: codeNOTFOUND}, err
	}
	resp, err := s.usersClient.AddEmployee(ctx, info)
	if err != nil {
		return &employmentService.BasicResponse{Message: resp.Message, Code: resp.Code}, err
	}
	return &employmentService.BasicResponse{Message: MessageOK, Code: CodeOK}, nil
}

// sending request to controller to employ the person

func (s GrpcServer) EmployPerson(_ context.Context, request *employmentService.EmploymentRequest) (*employmentService.BasicResponse, error) {
	// employing person
	err := s.cnt.EmployPerson(request.Email, request.DepartmentName, request.CompanyName)
	if err != nil {
		return &employmentService.BasicResponse{Message: MessageNotExecuted, Code: CodeNotExecuted}, err
	}

	ctx := context.Background()
	info, err := s.cnt.GetEmployeeInfoByEmail(request.Email)
	if err != nil {
		return &employmentService.BasicResponse{Message: MessageNOTFOUND, Code: codeNOTFOUND}, err
	}
	resp, err := s.usersClient.AddEmployee(ctx, info)
	if err != nil {
		return &employmentService.BasicResponse{Message: resp.Message, Code: resp.Code}, err
	}

	return &employmentService.BasicResponse{Message: MessageOK, Code: CodeOK}, nil
}

// sending request to controller to get info about employee

func (s GrpcServer) GetEmployeeInfoByEmail(_ context.Context, request *employmentService.ByEmailRequest) (*employmentService.EmployeeInfo, error) {
	info, err := s.cnt.GetEmployeeInfoByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	return info, nil

}
