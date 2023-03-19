package connection

import (
	"context"
	"fmt"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/connection/pb"
	"google.golang.org/grpc"
	"strconv"
)

type ServiceClient struct {
	Client pb.ConnectionClient
}

func InitServiceClient(url string) *ServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &ServiceClient{pb.NewConnectionClient(cc)}
}

func (svc *ServiceClient) GetSet(setFor, setOf string, id int64) []string {
	res, err := svc.Client.GetSet(context.Background(), &pb.GetSetRequest{
		Set: &pb.SetInfo{
			SetFor: setFor,
			SetOf:  setOf,
			Id:     strconv.FormatInt(id, 10),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Items
}

func (svc *ServiceClient) DeleteSet(setFor, setOf string, id int64) bool {
	res, err := svc.Client.DeleteSet(context.Background(), &pb.DeleteSetRequest{
		Set: &pb.SetInfo{
			SetFor: setFor,
			SetOf:  setOf,
			Id:     strconv.FormatInt(id, 10),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Success
}

func (svc *ServiceClient) AddToSet(setFor, setOf string, id int64, items ...string) bool {
	res, err := svc.Client.AddToSet(context.Background(), &pb.AddToSetRequest{
		Set: &pb.SetInfo{
			SetFor: setFor,
			SetOf:  setOf,
			Id:     strconv.FormatInt(id, 10),
		},
		Items: items,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Success
}

func (svc *ServiceClient) RemoveFromSet(setFor, setOf string, id int64, items ...string) bool {
	res, err := svc.Client.RemoveFromSet(context.Background(), &pb.RemoveFromSetRequest{
		Set: &pb.SetInfo{
			SetFor: setFor,
			SetOf:  setOf,
			Id:     strconv.FormatInt(id, 10),
		},
		Items: items,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Success
}
