package main

import (
	"net/http"
	"log"
	"html/template"
	"context"
	"time"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "project/rpc"
)

type Ship struct {
	Imo string
	Name string
}

var ports = map[string]string {
	"Livorno": "8090",
	"Barcellona": "8091",
}

func arrivalHandler(writer http.ResponseWriter, request *http.Request) {
	keys := portList()
	templ, _ := template.ParseFiles("arrival.html")
	templ.Execute(writer, keys)
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
	templ, _ := template.ParseFiles("portReply.html")
	templ.Execute(writer, r.GetMessage())
	conn.Close()
}

func departureHandler(writer http.ResponseWriter, request *http.Request) {
	keys := portList()
	templ, _ := template.ParseFiles("departure.html")
	templ.Execute(writer, keys)
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
	templ, _ := template.ParseFiles("portReply.html")
	templ.Execute(writer, r.GetMessage())
	conn.Close()
}

func bunkeringRequest(writer http.ResponseWriter, request *http.Request) {
	keys := portList()
	templ, _ := template.ParseFiles("bunkering.html")
	templ.Execute(writer, keys)
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
		templ, _ := template.ParseFiles("portReply.html")
		templ.Execute(writer, r.GetErrorMessage())
	} else {
		templ, _ := template.ParseFiles("bunkeringSuccess.html")
		type Response struct {
			Port string
			Imo string
			Name string
		}
		resp := Response{Port: port, Imo: r.GetTanker().GetImo(), Name: r.GetTanker().GetName()}
		templ.Execute(writer, &resp)
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
	
	templ, _ := template.ParseFiles("portReply.html")
	templ.Execute(writer, r.GetMessage())
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

/*func loadPage(fileName string) []byte {
	body, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return body
}*/

func main() {
	log.Println("Server on")

	http.HandleFunc("/arrival", arrivalHandler)
	http.HandleFunc("/registerArrival", registerArrival)

	http.HandleFunc("/departure", departureHandler)
	http.HandleFunc("/registerDeparture", registerDeparture)

	http.HandleFunc("/bunkering", bunkeringRequest)
	http.HandleFunc("/bunkeringHandler", bunkeringHandler)
	http.HandleFunc("/bunkeringEnd", bunkeringEndHandler)
	http.ListenAndServe(":8080", nil)
}