package stacksmith

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/JesusTinoco/go-smith/stacksmith/utils"
)

func retrieveItemsList(kind string) (*ListItems, *http.Response, error) {
	switch kind {
	case "components":
		return client.Discovery.ComponentsList("")
	case "services":
		return client.Discovery.ServicesList("")
	case "runtimes":
		return client.Discovery.RuntimesList("")
	case "oses":
		return client.Discovery.OsesList("")
	}
	return nil, nil, nil
}

func TestDiscovery_ListItems(t *testing.T) {
	setup()
	defer teardown()

	var items = []struct {
		name   string
		method string
	}{
		{"components", "ComponentsList"},
		{"services", "ServicesList"},
		{"runtimes", "RuntimesList"},
		{"oses", "OsesList"},
	}

	listItems := utils.GetJSON("all_items")

	itemsExpected := new(ListItems)
	json.Unmarshal(listItems, itemsExpected)

	for _, item := range items {
		mux.HandleFunc(fmt.Sprintf("/%s/", item.name), func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			w.Header().Set("Content-Type", "application/json")
			w.Write(listItems)
		})

		itemsRecieved, _, err := retrieveItemsList(item.name)

		if err != nil {
			t.Errorf("Discovery.%s returned error: %v", item.method, err.Error())
		}

		if !reflect.DeepEqual(itemsRecieved, itemsExpected) {
			t.Errorf("Discovery.%s returned %+v, want %+v", item.method, itemsRecieved, itemsExpected)
		}
	}
}

func retrieveFlavors(kind string) (*Flavors, *http.Response, error) {
	pag := &PaginationParams{Page: 1, PerPage: 100}
	switch kind {
	case "FlavorsList":
		return client.Discovery.FlavorsList(pag)
	case "GetFlavorsFrom":
		return client.Discovery.GetFlavorsFrom("test", pag)
	}
	return nil, nil, nil
}

func TestDiscoveryService_ListFlavors(t *testing.T) {
	setup()
	defer teardown()

	var items = []struct {
		url    string
		method string
	}{
		{"/flavors", "FlavorsList"},
		{"/components/test/flavors", "GetFlavorsFrom"},
	}

	listFlavors := utils.GetJSON("all_flavors")

	flavorsExpected := new(Flavors)
	json.Unmarshal(listFlavors, flavorsExpected)

	for _, item := range items {
		mux.HandleFunc(item.url, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testFormValues(t, r, values{"page": "1", "per_page": "100"})
			w.Header().Set("Content-Type", "application/json")
			w.Write(listFlavors)
		})

		flavorsRecieved, _, err := retrieveFlavors(item.method)
		if err != nil {
			t.Errorf("Discovery.%s returned error: %v", item.method, err.Error())
		}

		if !reflect.DeepEqual(flavorsRecieved, flavorsExpected) {
			t.Errorf("Discovery.%s returned %+v, want %+v", item.method, flavorsRecieved, flavorsExpected)
		}
	}
}

func TestDiscoverService_GetComponent(t *testing.T) {
	setup()
	defer teardown()

	component := utils.GetJSON("component")

	mux.HandleFunc("/components/apache", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json")
		w.Write(component)
	})

	componentRecieved, _, err := client.Discovery.GetComponent("apache")
	if err != nil {
		t.Errorf("Discovery.GetComponent returned error: %v", err.Error())
	}

	componentExpected := new(Item)
	json.Unmarshal(component, componentExpected)
	if !reflect.DeepEqual(componentRecieved, componentExpected) {
		t.Errorf("Discovery.Component returned %+v, want %+v", componentRecieved, componentExpected)
	}
}
