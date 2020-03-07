package main

import (
	"fmt"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry-tutorials/fake-mysql-broker/broker"
	"github.com/pivotal-cf/brokerapi"
)

func statusAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	logger := lager.NewLogger("fake-mysql-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO))
	logger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))

	servicebroker := &broker.SimpleBroker{
		Instances: map[string]brokerapi.GetInstanceDetailsSpec{},
		Bindings:  map[string]brokerapi.GetBindingSpec{},
	}

	brokerCredentials := brokerapi.BrokerCredentials{
		Username: "admin",
		Password: "secret",
	}
	brokerAPI := brokerapi.New(servicebroker, logger, brokerCredentials)
	http.HandleFunc("/health", statusAPI)
	http.Handle("/", brokerAPI)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("\n\nStarting Fake MySQL Broker on 0.0.0.0:" + port)
	logger.Fatal("http-listen", http.ListenAndServe("0.0.0.0:"+port, nil))
}
