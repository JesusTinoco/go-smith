package stacksmith

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/JesusTinoco/go-smith/stacksmith/utils"
)

func TestStacksService_List(t *testing.T) {
	setup()
	defer teardown()

	listStacks := utils.GetJSON("list_stacks")

	mux.HandleFunc("/stacks/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"page": "1", "per_page": "100", "api_key": "my_api_key"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(listStacks)
	})

	pag := &PaginationParams{Page: 1, PerPage: 100}
	stacksRecieved, _, err := client.Stacks.List(pag)
	if err != nil {
		t.Errorf("Stacks.List returned error: %v", err.Error())
	}

	stacksExpected := new(StacksList)
	json.Unmarshal(listStacks, stacksExpected)
	if !reflect.DeepEqual(stacksRecieved, stacksExpected) {
		t.Errorf("Stacks.List returned %+v, want %+v", stacksRecieved, stacksExpected)
	}
}

func TestStacksService_Get(t *testing.T) {
	setup()
	defer teardown()

	stack := utils.GetJSON("stack")

	mux.HandleFunc("/stacks/stack1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"api_key": "my_api_key"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(stack)
	})

	stackRecieved, _, err := client.Stacks.Get("stack1")
	if err != nil {
		t.Errorf("Stacks.Get returned error: %v", err.Error())
	}

	stackExpected := new(Stack)
	json.Unmarshal(stack, stackExpected)
	if !reflect.DeepEqual(stackRecieved, stackExpected) {
		t.Errorf("Stacks.Get returned %+v, want %+v", stackRecieved, stackExpected)
	}

}

func TestStacksService_Create(t *testing.T) {
	setup()
	defer teardown()

	response := utils.GetJSON("stack_response")

	mux.HandleFunc("/stacks/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"api_key": "my_api_key"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	stackDefinition := new(StackDefinition)
	json.Unmarshal(utils.GetJSON("create_stack_definition"), stackDefinition)
	responseRecieved, _, err := client.Stacks.Create(stackDefinition)
	if err != nil {
		t.Errorf("Stacks.Create returned error: %v", err.Error())
	}

	responseExpected := new(StatusGeneration)
	json.Unmarshal(response, responseExpected)
	if !reflect.DeepEqual(responseRecieved, responseExpected) {
		t.Errorf("Stacks.Create returned %+v, want %+v", responseRecieved, responseExpected)
	}
}

func TestStacksService_Delete(t *testing.T) {
}

func TestStacksService_Update(t *testing.T) {
	setup()
	defer teardown()

	response := utils.GetJSON("stack_response")

	mux.HandleFunc("/stacks/stack1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testFormValues(t, r, values{"api_key": "my_api_key"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	requestDefinition := new(StackParams)
	json.Unmarshal(utils.GetJSON("update_stack_definition"), requestDefinition)
	responseRecieved, _, err := client.Stacks.Update("stack1", requestDefinition)
	if err != nil {
		t.Errorf("Stacks.Update returned error: %v", err.Error())
	}

	responseExpected := new(StatusGeneration)
	json.Unmarshal(response, responseExpected)
	if !reflect.DeepEqual(responseRecieved, responseExpected) {
		t.Errorf("Stacks.Update returned %+v, want %+v", responseRecieved, responseExpected)
	}
}

func TestStacksService_Regenerate(t *testing.T) {
	setup()
	defer teardown()

	response := utils.GetJSON("stack_response")

	mux.HandleFunc("/stacks/stack1/regenerate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"api_key": "my_api_key"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	responseRecieved, _, err := client.Stacks.Regenerate("stack1")
	if err != nil {
		t.Errorf("Stacks.Update returned error: %v", err.Error())
	}

	responseExpected := new(StatusGeneration)
	json.Unmarshal(response, responseExpected)
	if !reflect.DeepEqual(responseRecieved, responseExpected) {
		t.Errorf("Stacks.Update returned %+v, want %+v", responseRecieved, responseExpected)
	}
}

func TestStacksService_GetVulnerabilities(t *testing.T) {
}
