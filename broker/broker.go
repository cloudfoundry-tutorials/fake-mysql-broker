package broker

import (
	"context"
	"fmt"
	"os"

	"github.com/pivotal-cf/brokerapi"
)

type SimpleBroker struct {
	Instances map[string]brokerapi.GetInstanceDetailsSpec
	Bindings  map[string]brokerapi.GetBindingSpec
}

func (simpleBroker *SimpleBroker) Services(ctx context.Context) ([]brokerapi.Service, error) {
	t := true
	return []brokerapi.Service{
		brokerapi.Service{
			ID:                   "mysql",
			Name:                 "mysql",
			Description:          "This service is for demonstration purposes. The same broker could advertise more than one service.",
			Bindable:             true,
			InstancesRetrievable: false,
			BindingsRetrievable:  false,
			Metadata: &brokerapi.ServiceMetadata{
				DisplayName: "mysql",
				ImageUrl:    os.Getenv("IMAGE_URL"),
			},
			Plans: []brokerapi.ServicePlan{
				brokerapi.ServicePlan{
					ID:          "free",
					Name:        "free",
					Description: "This is a plan. Plans can be used to create tiers or levels of service. For example, plans could be used to provide different amounts of cpu, memory, capacity, number of concurrent connections, network performance, etc.",
					Free:        &t,
					Bindable:    &t,
				},
			},
		},
	}, nil

}

func (simpleBroker *SimpleBroker) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	simpleBroker.Instances[instanceID] = brokerapi.GetInstanceDetailsSpec{
		ServiceID: details.ServiceID,
		PlanID:    details.PlanID,
	}
	return brokerapi.ProvisionedServiceSpec{
		IsAsync: false,
	}, nil
}

func (simpleBroker *SimpleBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	delete(simpleBroker.Instances, instanceID)
	return brokerapi.DeprovisionServiceSpec{
		IsAsync: false,
	}, nil
}

func (simpleBroker *SimpleBroker) GetInstance(ctx context.Context, instanceID string) (spec brokerapi.GetInstanceDetailsSpec, err error) {
	if instance, ok := simpleBroker.Instances[instanceID]; ok {
		return instance, nil
	}
	err = brokerapi.NewFailureResponse(fmt.Errorf("Instance Not Found. ID: %s", instanceID), 404, "get-instance")
	return
}

func (simpleBroker *SimpleBroker) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{
		IsAsync: false,
	}, nil
}

func (simpleBroker *SimpleBroker) LastOperation(ctx context.Context, instanceID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{
		State: brokerapi.Succeeded,
	}, nil
}

func (simpleBroker *SimpleBroker) Bind(ctx context.Context, instanceID string, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding, error) {
	credentials := make(map[string]string)
	credentials["host"] = "mysql-service.service.consul"
	credentials["hostname"] = "mysql-service.service.consul"
	credentials["port"] = "3306"
	credentials["name"] = "CF_99FA13ABAE21"
	credentials["database"] = "CF_99FA13ABAE21"
	credentials["username"] = "khasdfewf728s"
	credentials["password"] = "kljopuklloKJJH0khlkjl"
	credentials["database_uri"] = "mysql://khasdfewf728s:kljopuklloKJJH0khlkjl@mysql-service.service.consul:3306/CF_99FA13ABAE21?reconnect=true"
	credentials["uri"] = "mysql://khasdfewf728s:kljopuklloKJJH0khlkjl@mysql-service.service.consul:3306/CF_99FA13ABAE21?reconnect=true"
	credentials["jdbcUrl"] = "jdbc:mysql://mysql-service.service.consul:3306/CF_99FA13ABAE21?user=khasdfewf728s&password=kljopuklloKJJH0khlkjl"

	simpleBroker.Bindings[bindingID] = brokerapi.GetBindingSpec{
		Credentials: credentials,
	}
	return brokerapi.Binding{
		Credentials: credentials,
	}, nil
}

func (simpleBroker *SimpleBroker) Unbind(ctx context.Context, instanceID string, bindingID string, details brokerapi.UnbindDetails, asyncAllowed bool) (brokerapi.UnbindSpec, error) {
	delete(simpleBroker.Bindings, bindingID)
	return brokerapi.UnbindSpec{}, nil
}

func (simpleBroker *SimpleBroker) GetBinding(ctx context.Context, instanceID string, bindingID string) (spec brokerapi.GetBindingSpec, err error) {
	if binding, ok := simpleBroker.Bindings[bindingID]; ok {
		return binding, nil
	}
	err = brokerapi.NewFailureResponse(fmt.Errorf("Service Binding Not Found. ID: %s", bindingID), 404, "get-binding")
	return
}

func (simpleBroker *SimpleBroker) LastBindingOperation(ctx context.Context, instanceID string, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{
		State: brokerapi.Succeeded,
	}, nil
}
