package main

import (
	"net/http"
	"log"
	"html/template"
	"context"
	"time"
	"strconv"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "project/rpc"
)

type Ship struct {
	Imo string
	Name string
}

var ports = make(map[string]string)

func registerNewPort(writer http.ResponseWriter, request *http.Request) {
	portName := request.URL.Query().Get("name")
	portConnection := request.URL.Query().Get("portConnection")
	ports[portName] = portConnection
}

func arrivalHandler(writer http.ResponseWriter, request *http.Request) {
	keys := portList()
	templ, _ := template.ParseFiles("layout.html", "arrival.html")
	templ.ExecuteTemplate(writer, "layout", keys)
}

func registerArrival(writer http.ResponseWriter, request *http.Request) {
	port := request.PostFormValue("port")
	ship := pb.Ship {
		Imo: request.PostFormValue("imo"),
		Name: request.PostFormValue("shipName"),
	}
	conn := portConnection(port)

	c := pb.NewRegisterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Arrival(ctx, &ship)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	templ, _ := template.ParseFiles("layout.html", "portReply.html")
	templ.ExecuteTemplate(writer, "layout", r.GetMessage())
	conn.Close()
}

func departureHandler(writer http.ResponseWriter, request *http.Request) {
	keys := portList()
	templ, _ := template.ParseFiles("layout.html", "departure.html")
	templ.ExecuteTemplate(writer, "layout", keys)
}

func registerDeparture(writer http.ResponseWriter, request *http.Request) {
	port := request.PostFormValue("port")
	ship := pb.DepartingShip {
		Imo: request.PostFormValue("imo"),
		Name: request.PostFormValue("shipName"),
		Destination: request.PostFormValue("destination"),
	}
	conn := portConnection(port)

	c := pb.NewRegisterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Departure(ctx, &ship)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	templ, _ := template.ParseFiles("layout.html", "portReply.html")
	templ.ExecuteTemplate(writer, "layout", r.GetMessage())
	conn.Close()
}

func bunkeringRequest(writer http.ResponseWriter, request *http.Request) {
	keys := portList()
	templ, _ := template.ParseFiles("layout.html", "bunkering.html")
	templ.ExecuteTemplate(writer, "layout", keys)
}

func bunkeringHandler(writer http.ResponseWriter, request *http.Request) {
	port := request.PostFormValue("port")
	bunkeringReq := pb.BunkeringRequest {
		Imo: request.PostFormValue("imo"),
	}
	conn := portConnection(port)

	c := pb.NewRegisterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Bunkering(ctx, &bunkeringReq)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	
	if r.GetErrorMessage() != "" {
		templ, _ := template.ParseFiles("layout.html", "portReply.html")
		templ.ExecuteTemplate(writer, "layout", r.GetErrorMessage())
	} else {
		templ, _ := template.ParseFiles("layout.html", "bunkeringSuccess.html")
		type Response struct {
			Port string
			Imo string
			Name string
		}
		resp := Response{Port: port, Imo: r.GetShip().GetImo(), Name: r.GetShip().GetName()}
		templ.ExecuteTemplate(writer, "layout", &resp)
	}
	conn.Close()
}

func bunkeringEndHandler(writer http.ResponseWriter, request *http.Request) {
	port := request.PostFormValue("port")
	bunkeringReq := pb.BunkeringRequest {
		Imo: request.PostFormValue("imo"),
	}
	conn := portConnection(port)

	c := pb.NewRegisterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.BunkeringEnd(ctx, &bunkeringReq)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	
	templ, _ := template.ParseFiles("layout.html", "portReply.html")
	templ.ExecuteTemplate(writer, "layout", r.GetMessage())
	conn.Close()
}

func tugsRequest(writer http.ResponseWriter, request *http.Request) {
	keys := portList()
	templ, _ := template.ParseFiles("layout.html", "tugs.html")
	templ.ExecuteTemplate(writer, "layout", keys)
}

func tugsHandler(writer http.ResponseWriter, request *http.Request) {
	port := request.PostFormValue("port")
	n, _ := strconv.ParseInt(request.PostFormValue("tugsNumber"), 10, 32)
	tugsReq := pb.TugsRequest {
		Imo: request.PostFormValue("imo"),
		Type: request.PostFormValue("requestType"),
		TugsNumber: int32(n),
	}
	conn := portConnection(port)

	c := pb.NewRegisterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AcquireTugs(ctx, &tugsReq)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	
	if r.GetErrorMessage() != "" {
		templ, _ := template.ParseFiles("layout.html", "portReply.html")
		templ.ExecuteTemplate(writer, "layout", r.GetErrorMessage())
	} else {
		templ, _ := template.ParseFiles("layout.html", "tugInfo.html")
		type Response struct {
			Port string
			Ships []*pb.Ship
		}
		resp := Response{Port: port, Ships: r.GetShips()}
		templ.ExecuteTemplate(writer, "layout", &resp)
	}
	conn.Close()
}

func releaseTugsHandler(writer http.ResponseWriter, request *http.Request) {
	// retrieve the imos
	i := 0
	imoList := make([]string, 0)
	for {
		imo := request.PostFormValue(strconv.Itoa(i))
		if imo != "" {
			imoList = append(imoList, imo)
			i++
		} else {
			break
		}
	}

	port := request.PostFormValue("port")
	tugsReq := pb.ReleaseTugsRequest {
		ImoList: imoList,
	}
	conn := portConnection(port)

	c := pb.NewRegisterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ReleaseTugs(ctx, &tugsReq)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	
	templ, _ := template.ParseFiles("layout.html", "portReply.html")
	templ.ExecuteTemplate(writer, "layout", r.GetMessage())
	conn.Close()
}

func portList() []string{
	keys := make([]string, 0)

	for k, _ := range ports {
		keys = append(keys, k)
	}
	return keys
}

func portConnection(port string) *grpc.ClientConn {
	conn, err := grpc.Dial(":" + ports[port], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return conn
}

func main() {
	log.Println("Server on")

	http.HandleFunc("/registerPort", registerNewPort)

	http.HandleFunc("/arrival", arrivalHandler)
	http.HandleFunc("/registerArrival", registerArrival)

	http.HandleFunc("/departure", departureHandler)
	http.HandleFunc("/registerDeparture", registerDeparture)

	http.HandleFunc("/bunkering", bunkeringRequest)
	http.HandleFunc("/bunkeringHandler", bunkeringHandler)
	http.HandleFunc("/bunkeringEnd", bunkeringEndHandler)

	http.HandleFunc("/tugs", tugsRequest)
	http.HandleFunc("/tugsHandler", tugsHandler)
	http.HandleFunc("/releaseTugs", releaseTugsHandler)
	http.ListenAndServe(":8080", nil)
}