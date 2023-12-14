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
	keys := make([]string, 0)
	for k, _ := range ports {
		keys = append(keys, k)
	}
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
	log.Println(r.GetMessage())
	conn.Close()
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
	http.ListenAndServe(":8080", nil)
}