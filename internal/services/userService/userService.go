package userService

import (
	"HeadHunter/pkg/employmentService"
	"HeadHunter/pkg/userService"
	"context"
	"google.golang.org/grpc"
)

type Client struct {
	client userService.UserServiceClient
}

// constructor for user service client

func Init(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		client: userService.NewUserServiceClient(conn),
	}, nil
}

// adding employee in user service

func (c *Client) AddEmployee(_ context.Context, employee *employmentService.EmployeeInfo) (*userService.Basic, error) {

	ctx := context.Background()
	res, err := c.client.AddEmployee(ctx, ConvertToUserServiceEmployee(employee))
	return res, err
}
