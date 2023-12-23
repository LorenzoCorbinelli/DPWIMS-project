package main

import (
	"context"
	"net"
	"net/http"
	"fmt"
	"log"
	"os"
	"google.golang.org/grpc"
	pb "project/rpc"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	dbm "project/ports/common"
)

type server struct {
	pb.UnimplementedRegisterServer
	Db *gorm.DB
}

func createPort(name string, portConnection string, bunkeringShips []dbm.BunkeringShips, tugs []dbm.Tugs) {
	// communication of name and port connection at the server
	requestURL := fmt.Sprintf("http://localhost:8080/registerPort?name=%s&portConnection=%s", name, portConnection)
	req, reqErr := http.NewRequest(http.MethodPost, requestURL, nil)
	if reqErr != nil {
		log.Fatal(reqErr.Error())
		return
	}
	http.DefaultClient.Do(req)

	dbName := fmt.Sprintf("%s.db", name)
	os.Remove(dbName)
	log.Println(name + " on")
	port := server{}
	var err error
	port.Db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error)
		return
	}
	dbm.CreateTables(port.Db)
	dbm.SetUpBunkeringShips(port.Db, bunkeringShips)
	dbm.SetUpTugs(port.Db, tugs)

	portConnection = fmt.Sprintf(":%s", portConnection)
	lis, err := net.Listen("tcp", portConnection)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterRegisterServer(s, &port)
	lis.Addr()
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (s *server) Arrival(ctx context.Context, in *pb.Ship) (*pb.Reply, error) {
	dbm.InsertNewArrival(s.Db, in.GetImo(), in.GetName())
	return &pb.Reply{Message: "Arrival registered"}, nil
}

func (s *server) Departure(ctx context.Context, in *pb.DepartingShip) (*pb.Reply, error) {
	dbm.InsertNewDeparture(s.Db, in.GetImo(), in.GetName(), in.GetDestination())
	return &pb.Reply{Message: "Departure registered"}, nil
}

func (s *server) Bunkering(ctx context.Context, in *pb.BunkeringRequest) (*pb.ShipReply, error) {
	result, tanker := dbm.Bunkering(s.Db, in.GetImo())
	if result == -1 {	// the ship (client) is not in this port
		return &pb.ShipReply{ErrorMessage: "The ship that requested a bunkering operation is not in this port", Ship: nil}, nil
	}
	if result == 0 {	// bunkering ships unavailable
		return &pb.ShipReply{ErrorMessage: "All the bunkering ships are unavailable", Ship: nil}, nil
	}
	ship := pb.Ship{Imo: tanker.Imo, Name: tanker.Name}
	return &pb.ShipReply{ErrorMessage: "", Ship: &ship}, nil
}

func (s *server) BunkeringEnd(ctx context.Context, in *pb.BunkeringRequest) (*pb.Reply, error) {
	dbm.BunkeringEnd(s.Db, in.GetImo())
	return &pb.Reply{Message: "Bunkering ended successfully"}, nil
}

func (s *server) AcquireTugs(ctx context.Context, in *pb.TugsRequest) (*pb.TugsReply, error) {
	result, tugs := dbm.AcquireTugs(s.Db, in.GetImo(), in.GetType(), int(in.GetTugsNumber()))
	if result == -1 {	// the ship (client) is not in this port and has requested tugs for a departure
		return &pb.TugsReply{ErrorMessage: "The ship that requested tugs for a departure is not in this port", Ships: nil}, nil
	}
	if result == 0 {	// not enough available tugs
		return &pb.TugsReply{ErrorMessage: "Not enough available tugs for the request", Ships: nil}, nil
	}
	ships := make([]*pb.Ship, 0)
	for _, tug := range tugs {
		ships = append(ships, &pb.Ship{Imo: tug.Imo, Name: tug.Name})
	}
	return &pb.TugsReply{ErrorMessage: "", Ships: ships}, nil
}

func (s *server) ReleaseTugs(ctx context.Context, in *pb.ReleaseTugsRequest) (*pb.Reply, error) {
	dbm.ReleaseTugs(s.Db, in.GetImoList())
	return &pb.Reply{Message: "Tugs released successfully"}, nil
}

func main() {
	tankers := make([]dbm.BunkeringShips, 0)
	tankers = append(tankers, dbm.BunkeringShips{
		Imo: "9487744",
		Name: "Elba",
		Available: true,
	})
	tankers = append(tankers, dbm.BunkeringShips{
		Imo: "9304485",
		Name: "Gorgona",
		Available: true,
	})
	tankers = append(tankers, dbm.BunkeringShips{
		Imo: "9365207",
		Name: "Giglio",
		Available: true,
	})

	tugs := make([]dbm.Tugs, 0)
	tugs = append(tugs, dbm.Tugs{
		Imo: "9443750",
		Name: "Costante neri",
		Available: true,
	})
	tugs = append(tugs, dbm.Tugs{
		Imo: "9443748",
		Name: "Corrado neri",
		Available: true,
	})
	tugs = append(tugs, dbm.Tugs{
		Imo: "9842968",
		Name: "Calafuria",
		Available: true,
	})
	tugs = append(tugs, dbm.Tugs{
		Imo: "9779252",
		Name: "Antignano",
		Available: true,
	})

	go createPort("Livorno", "8090", tankers, tugs)

	tankers = make([]dbm.BunkeringShips, 0)
	tankers = append(tankers, dbm.BunkeringShips{
		Imo: "9280378",
		Name: "Spabunker veintidos",
		Available: true,
	})
	tankers = append(tankers, dbm.BunkeringShips{
		Imo: "9301172",
		Name: "Petrobay",
		Available: true,
	})
	tankers = append(tankers, dbm.BunkeringShips{
		Imo: "9391177",
		Name: "Greenoil",
		Available: true,
	})
	tankers = append(tankers, dbm.BunkeringShips{
		Imo: "9398955",
		Name: "Spabunker cincuenta",
		Available: true,
	})
	
	tugs = make([]dbm.Tugs, 0)
	tugs = append(tugs, dbm.Tugs{
		Imo: "9881328",
		Name: "Azabra",
		Available: true,
	})
	tugs = append(tugs, dbm.Tugs{
		Imo: "9390771",
		Name: "Eliseo vazquez",
		Available: true,
	})
	tugs = append(tugs, dbm.Tugs{
		Imo: "9439723",
		Name: "Montclar",
		Available: true,
	})
	go createPort("Barcellona", "8091", tankers, tugs)

	// just to keep the ports active
	c := make(chan int)
	wait := <-c
	fmt.Println(wait)
}