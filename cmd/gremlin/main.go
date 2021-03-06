package main

import (
	"context"
	"flag"
	"fmt"
	"gremlins/src/gremlin"
	"gremlins/src/log"
	"gremlins/src/registry"
	"gremlins/src/service"
	stlog "log"
)

func main() {
	host := "localhost"
	var port string = "7000"
	flag.StringVar(&port, "port", "7000", "port to bind to")
	flag.Parse()
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GremlinService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		gremlin.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down gremlinservice")
}
