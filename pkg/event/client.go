package event

import (
	"context"
	"fmt"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/event/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.EventClient
}

func InitServiceClient(url string) *ServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &ServiceClient{pb.NewEventClient(cc)}
}

func (svc *ServiceClient) AddDiscipline(name string) int64 {
	res, err := svc.Client.AddDiscipline(context.Background(), &pb.AddDisciplineRequest{
		Name: name,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Id
}

func (svc *ServiceClient) DeleteDiscipline(id int64) bool {
	res, err := svc.Client.DeleteDiscipline(context.Background(), &pb.DeleteDisciplineRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Success
}

func (svc *ServiceClient) GetDisciplines(ids ...int64) []*pb.DisciplineData {
	res, err := svc.Client.GetDisciplines(context.Background(), &pb.GetDisciplinesRequest{
		Id: ids,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Disciplines
}

func (svc *ServiceClient) AddEvent(disciplineId int64, place, date, startTime, endTime, price, limit, description string) int64 {
	res, err := svc.Client.AddEvent(context.Background(), &pb.AddEventRequest{
		DisciplineId: disciplineId,
		Place:        place,
		Date:         date,
		StartTime:    startTime,
		EndTime:      endTime,
		Price:        price,
		Limit:        limit,
		Description:  description,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Id
}

func (svc *ServiceClient) DeleteEvent(id int64) bool {
	res, err := svc.Client.DeleteEvent(context.Background(), &pb.DeleteEventRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Success
}

func (svc *ServiceClient) GetEvents(disciplineIds ...int64) []*pb.EventData {
	res, err := svc.Client.GetEvents(context.Background(), &pb.GetEventsRequest{
		DisciplineId: disciplineIds,
	})
	if err != nil {
		fmt.Println(err)
	}
	return res.Events
}
