package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"context"
	"time"
	"strconv"
	"strings"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "project/rpc"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Ship struct {
	Imo string
	Name string
}

var ports = make(map[string]string)

func registerNewPort(client mqtt.Client, msg mqtt.Message) {
	payload := strings.Split(string(msg.Payload()), ":")
	ports[payload[0]] = payload[1]	// payload[0] is the port name and payload[1] is the port connection
}

func portDisconnected(client mqtt.Client, msg mqtt.Message) {
	payload := strings.Split(string(msg.Payload()), ":")
	delete(ports, payload[0])	// a port has disconnected and so I remove that port from the map
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	templ, _ := template.ParseFiles("layout.html", "index.html")
	templ.ExecuteTemplate(writer, "layout", nil)
}

func operationSelected(writer http.ResponseWriter, request *http.Request) {
	operation := request.PostFormValue("operation")
	url := fmt.Sprintf("http://localhost:8080/%s", operation)
	http.Redirect(writer, request, url, http.StatusSeeOther)
}

func arrivalHandler(writer http.ResponseWriter, request *http.Request) {
	ports := portList()
	templ, _ := template.ParseFiles("layout.html", "arrival.html")
	templ.ExecuteTemplate(writer, "layout", ports)
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
	ports := portList()
	templ, _ := template.ParseFiles("layout.html", "departure.html")
	templ.ExecuteTemplate(writer, "layout", ports)
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
	ports := portList()
	templ, _ := template.ParseFiles("layout.html", "bunkering.html")
	templ.ExecuteTemplate(writer, "layout", ports)
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
	ports := portList()
	templ, _ := template.ParseFiles("layout.html", "tugs.html")
	templ.ExecuteTemplate(writer, "layout", ports)
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
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("server")
	mqttClient := mqtt.NewClient(opts)
	token := mqttClient.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
		return
	}
	sub := mqttClient.Subscribe("ports/register", 0, registerNewPort)
	if sub.Wait() && sub.Error() != nil {
		log.Fatal(sub.Error())
		return
	}
	sub = mqttClient.Subscribe("ports/disconnection", 0, portDisconnected)
	if sub.Wait() && sub.Error() != nil {
		log.Fatal(sub.Error())
		return
	}

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/operations", operationSelected)

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

	log.Println("Server on")

	http.ListenAndServe(":8080", nil)
}