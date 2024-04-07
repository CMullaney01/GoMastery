package microservice

import (
	"context"
	"log"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

func Microservice() {

	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)

	apiServer := NewApiServer(svc)
	log.Fatal(apiServer.Start(":3000"))

}
