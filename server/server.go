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

	res = "Port "  + strconv.Itoa(int(id))  + " Successfully Created"
	return &pb.CreatePortResponse{Result: res}, nil
}

//retrieve new port
func (*server) RetreivePort(ctx context.Context, in *pb.RetrievePortRequest) (*pb.RetrievePortResponse, error) {

	id := in.GetPortId()
	Id, name, code, city, state, country := Getportdetails(id)

	return &pb.RetrievePortResponse{Id: Id, Name: name, Code: code, City: city, State: state, Country: country}, nil
}

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
		res = "Port " + strconv.Itoa(int(id))+ " Successfully updated"
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
		res = "Port " + strconv.Itoa(int(id))+ " Successfully created "
		return &pb.UpdatePortResponse{Result: res}, nil
	}

}

func (*server) DeletePort(ctx context.Context, in *pb.DeletePortResquest) (*pb.DeletePortResponse, error) {
	id := in.GetPortId()
	DeletePortDetails(id)

	var res string
	res = "Port " +strconv.Itoa(int(id))+" Successfully Deleted"
	return &pb.DeletePortResponse{Result: res}, nil
}

//Id:Id,Name: name,Code: code,City: city,State: state,Country: country
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