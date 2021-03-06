package main

import (
	"context"
	"fmt"
	"gremlins/src/keeper"
	"gremlins/src/log"
	"gremlins/src/registry"
	"gremlins/src/service"
	stlog "log"
)

func main() {
	host, port := "localhost", "8000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.KeeperService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService, registry.GremlinService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		keeper.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down keeper service")
}
