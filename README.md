# DPWIMS-project
## How to run the project:
First you have to run a MQTT broker on the 1883 port, for example you can use [Mosquitto](https://mosquitto.org/).

Then in the project directory run:
```
go run server.go
```
Wait the response _Server on_ (the first time can take a little).

In another terminal go to the `ports` directory and run:
```
go run portManagement.go
```
After the response, the program is ready.
