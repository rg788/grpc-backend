package main

import (
	//"database/sql"
	//"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	//srv "grpc-backend/server"
	//"os"
	//"os/signal"
	"context"
	pb "grpc-backend/gen/proto"

	_ "github.com/lib/pq"
)

/* func dbConnection(){
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to DB")

	connStr := "user=postgres dbname=test1 password=password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
} */
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

	Createnewport(id, name, code, city, state, country)
	return &pb.CreatePortResponse{Result: "Successfully Created A port"}, nil
}

//retrieve new port
func (*server) RetreivePort(ctx context.Context, in *pb.RetrievePortRequest) (*pb.RetrievePortResponse, error) {

	id := in.GetPortId()
	Id, name, code, city, state, country := Getportdetails(id)

	return &pb.RetrievePortResponse{Id: Id, Name: name, Code: code, City: city, State: state, Country: country}, nil
}

//Id:Id,Name: name,Code: code,City: city,State: state,Country: country
func main() {

	//dbConnection()

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

//func createnewport(id, name, code, city, state, country string) {}
