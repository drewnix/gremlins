package service

import (
	"context"
	"fmt"
	"gremlins/src/registry"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start(ctx context.Context, host, port string, reg registry.Registration, RegisterHandlersFunc func(e *echo.Echo)) (context.Context, error) {
	e := echo.New()
	RegisterHandlersFunc(e)

	ctx = startService(ctx, reg.ServiceName, e, host, port)
	err := registry.RegisterService(reg, e)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, e *echo.Echo, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	srvPort := ":" + port

	go func() {
		e.Use(middleware.Logger())
		// Start server
		e.Logger.Fatal(e.Start(srvPort))
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%v:%v", host, port))
		if err != nil {
			log.Println(err)
		}
		// srv.Shutdown(ctx)
		e.Logger.Fatal(ctx)
		cancel()
	}()
	return ctx
}
