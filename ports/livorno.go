package main

import (
	"context"
	"net"
	"log"
	"os"
	"google.golang.org/grpc"
	pb "project/rpc"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	dbm "project/ports/common"
)

var db *gorm.DB

type server struct {
	pb.UnimplementedRegisterServer
}

func (s *server) Arrival(ctx context.Context, in *pb.Ship) (*pb.Reply, error) {
	dbm.InsertNewArrival(db, in.GetImo(), in.GetName())
	return &pb.Reply{Message: "Arrival registered"}, nil
}

func (s *server) Departure(ctx context.Context, in *pb.DepartingShip) (*pb.Reply, error) {
	dbm.InsertNewDeparture(db, in.GetImo(), in.GetName(), in.GetDestination())
	return &pb.Reply{Message: "Departure registered"}, nil
}

func (s *server) Bunkering(ctx context.Context, in *pb.BunkeringRequest) (*pb.ShipReply, error) {
	result, tanker := dbm.Bunkering(db, in.GetImo())
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
	dbm.BunkeringEnd(db, in.GetImo())
	return &pb.Reply{Message: "Bunkering ended successfully"}, nil
}

func (s *server) AcquireTugs(ctx context.Context, in *pb.TugsRequest) (*pb.TugsReply, error) {
	result, tugs := dbm.AcquireTugs(db, in.GetImo(), in.GetType(), int(in.GetTugsNumber()))
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

func main() {
	os.Remove("livorno.db")
	log.Println("Port on")
	var err error
	db, err = gorm.Open(sqlite.Open("livorno.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error)
		return
	}

	dbm.CreateTables(db)

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
	dbm.SetUpBunkeringShips(db, tankers)

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
	dbm.SetUpTugs(db, tugs)

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterRegisterServer(s, &server{})
	lis.Addr()
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err.Error())
	}
}