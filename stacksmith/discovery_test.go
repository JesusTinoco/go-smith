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

func TestStacksService_ListFlavors(t *testing.T) {
	setup()
	defer teardown()

	listFlavors := utils.GetJSON("all_flavors")

	mux.HandleFunc("/flavors/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"page": "1", "per_page": "100"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(listFlavors)
	})

	pag := &PaginationParams{Page: 1, PerPage: 100}
	flavorsRecieved, _, err := client.Discovery.FlavorsList(pag)
	if err != nil {
		t.Errorf("Discovery.FlavorsList returned error: %v", err.Error())
	}

	flavorsExpected := new(Flavors)
	json.Unmarshal(listFlavors, flavorsExpected)
	if !reflect.DeepEqual(flavorsRecieved, flavorsExpected) {
		t.Errorf("Discovery.FlavorsList returned %+v, want %+v", flavorsRecieved, flavorsExpected)
	}
}
