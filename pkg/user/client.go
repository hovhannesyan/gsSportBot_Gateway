package user

import (
	"context"
	"fmt"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/user/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.UserClient
}

func InitServiceClient(url string) *ServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &ServiceClient{pb.NewUserClient(cc)}
}

func (svc *ServiceClient) AddUser(id int64, username, name, phone string, isAdmin bool) string {
	res, err := svc.Client.AddUser(context.Background(), &pb.AddUserRequest{
		Id:       id,
		Username: username,
		Name:     name,
		Phone:    phone,
		IsAdmin:  isAdmin,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Message
}

func (svc *ServiceClient) IsAdmin(id int64) bool {
	res, err := svc.Client.IsAdmin(context.Background(), &pb.IsAdminRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.IsAdmin
}

func (svc *ServiceClient) GetUsersById(ids ...int64) []*pb.UserData {
	res, err := svc.Client.GetUsersById(context.Background(), &pb.GetUsersByIdRequest{
		Id: ids,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Data
}
