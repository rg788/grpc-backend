package main

import (
	"context"
	pb "grpc-backend/gen/proto"
	"log"
	"net"
	"strconv"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPortServiceServer
}

//Creating new port

func (*server) CreatePort(ctx context.Context, in *pb.CreatePortRequest) (*pb.CreatePortResponse, error) {

	id := in.Port.GetId()
	name := in.Port.GetName()
	code := in.Port.GetCode()
	city := in.Port.GetCity()
	state := in.Port.GetState()
	country := in.Port.GetCountry()

	err := Createnewport(id, name, code, city, state, country)
	var res string
	if err != nil {
		res = "Port ID already exist"
		return &pb.CreatePortResponse{Result: res}, nil
	}

	res = "Port " + strconv.Itoa(int(id)) + " Successfully Created"
	return &pb.CreatePortResponse{Result: res}, nil
}

//retrieve new port
func (*server) RetreivePort(ctx context.Context, in *pb.RetrievePortRequest) (*pb.RetrievePortResponse, error) {

	id := in.GetPortId()
	Id, name, code, city, state, country := Getportdetails(id)

	return &pb.RetrievePortResponse{Id: Id, Name: name, Code: code, City: city, State: state, Country: country}, nil
}

// Updating port
func (*server) UpdatePort(ctx context.Context, in *pb.UpdatePortRequest) (*pb.UpdatePortResponse, error) {

	id := in.Port.GetId()
	if checkPortId(id) {
		Id := in.Port.GetId()
		Name := in.Port.GetName()
		Code := in.Port.GetCode()
		City := in.Port.GetCity()
		State := in.Port.GetState()
		Country := in.Port.GetCountry()
		UpdatePortDetails(Id, Name, Code, City, State, Country)

		var res string
		res = "Port " + strconv.Itoa(int(id)) + " Successfully updated"
		return &pb.UpdatePortResponse{Result: res}, nil
	} else {
		//id := in.Port.GetId()
		name := in.Port.GetName()
		code := in.Port.GetCode()
		city := in.Port.GetCity()
		state := in.Port.GetState()
		country := in.Port.GetCountry()
		Createnewport(id, name, code, city, state, country)
		var res string
		res = "Port " + strconv.Itoa(int(id)) + " Not found, New Port Successfully created "
		return &pb.UpdatePortResponse{Result: res}, nil
	}

}

//Delete port
func (*server) DeletePort(ctx context.Context, in *pb.DeletePortResquest) (*pb.DeletePortResponse, error) {
	id := in.GetPortId()
	err := DeletePortDetails(id)

	var res string
	res = "Port " + strconv.Itoa(int(id)) + " Successfully Deleted"
	return &pb.DeletePortResponse{Result: res}, err
}

//Pagination
func (* server)ListPort(in *pb.ListPortRequest,stream pb.PortService_ListPortServer)error{

	count := in.Count
	page := in.Page

	allPorts := getAllPorts(page, count)
	var i int32
	for i = 0; i < count; i++ {
		var port1 = allPorts[i]
		//fmt.Println(port1)
		res:= &pb.ListPortResponse{Port:&pb.Port{Id: port1.ID, Name: port1.Name, Code: port1.Code, City: port1.City, State: port1.State, Country: port1.Country}}
		stream.Send(res)
	}

	return nil

}

func main() {

	dbConnection()
	log.Println("Server Started")
	lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatalf("Error while listening , server %v", err)
	}
	sGRCP := grpc.NewServer()
	pb.RegisterPortServiceServer(sGRCP, &server{})
	if err := sGRCP.Serve(lis); err != nil {
		log.Fatalf("Error while runnig perService %v", err)
	}

}
