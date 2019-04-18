package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func clearTodos() error {
	client := http.Client{}
	//Clear all todos (already tested)
	delReq, err := http.NewRequest("DELETE", "http://localhost:8080/todos", nil)
	if err != nil {
		return err
	}
	client.Do(delReq)
	return nil
}

func getTodoList() (*string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/todos", nil)
	if err != nil { //
		return nil, err
	}
	//rr := executeRequest(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("unexpected error " + err.Error())
	}
	respBodyStr := new(string)
	*respBodyStr = string(respBody)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Expected status " + strconv.Itoa(http.StatusOK) +
			" received status " + strconv.Itoa(resp.StatusCode) +
			" response " + *respBodyStr)
	}

	return respBodyStr, nil
}

func TestEmptyModel(t *testing.T) {

	respBodyStr, err := getTodoList()
	if err != nil {
		t.Fatalf("Error " + err.Error())
	}
	if *respBodyStr != "[]" {
		t.Fatal("Expected result string \"[]\" received result string " + *respBodyStr)
	}
}

func TestDeleteAllTodos(t *testing.T) {

	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/todos", nil)
	if err != nil {
		t.Fatal("unexpected error " + err.Error())
	}
	response, err := client.Do(req)
	if err != nil {
		t.Fatal("unexpected error " + err.Error())
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("unexpected error " + err.Error())
	}

	var responsemap map[string]string
	err = json.Unmarshal(body, &responsemap)
	if responsemap["result"] != "success" {
		t.Fatal("expected result = \"success\" received = " + responsemap["result"])
	}

	//test for empty list
	respBodyStr, err := getTodoList()
	if err != nil {
		t.Fatalf("Error " + err.Error())
	}
	if *respBodyStr != "[]" {
		t.Fatal("Expected result string \"[]\" received result string " + *respBodyStr)
	}

}

func TestNonExistentTodo(t *testing.T) {

	//Clear all todos (already tested)
	err := clearTodos()
	if err != nil {
		t.Fatal("Clear all todos error " + err.Error())
	}

	//Get non - existent todo
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/todo/1", nil)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatal("Expected status " + strconv.Itoa(http.StatusNotFound) +
			" Received status " + strconv.Itoa(resp.StatusCode))
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error while reading response " + err.Error())
	}

	var responseMap map[string]string
	err = json.Unmarshal(response, &responseMap)
	if err != nil {
		t.Fatal("error while unmarshalling response " + err.Error())
	}

	if responseMap["error"] != "error not found" {
		t.Fatal("Expceted error = \"not found\" Received = \"" + responseMap["message"] + "\"")
	}
}

func TestInvalidTodoRequest(t *testing.T) {

	//Clear all todos (already tested)
	err := clearTodos()
	if err != nil {
		t.Fatal("Clear all todos error " + err.Error())
	}

	//Get non - existent todo
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/todo/abc", nil)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatal("Expected status " + strconv.Itoa(http.StatusNotFound) +
			" Received status " + strconv.Itoa(resp.StatusCode))
	}

}

func createTodo(t *testing.T, payload []byte, expectedStatusCode int) map[string]interface{} {

	client := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/todo", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal("unexpected error " + err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("unexpected error " + err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != expectedStatusCode {
		t.Fatal("Expected status code " + strconv.Itoa(expectedStatusCode) +
			" Recd " + strconv.Itoa(resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("unexpected error " + err.Error())
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(respBody, &responseMap)

	return responseMap

}

func TestCreateTodo(t *testing.T) {

	//Clear all todos (already tested)
	err := clearTodos()
	if err != nil {
		t.Fatal("Clear all todos error " + err.Error())
	}

	//Create todo
	// todopayload := "{ \"message\" : \"message1\" }"
	todopayload := []byte(`{"message":"message1"}`)
	responseMap := createTodo(t, todopayload, http.StatusCreated)
	if responseMap["message"] != "message1" {
		t.Fatal("expected message to be message1 received = ", responseMap["message"])
	}
	if responseMap["ID"] != 1.0 {
		t.Fatal("expected ID = 1 received = ", responseMap["ID"])
	}

	todopayload2 := []byte(`{"message":"message2"}`)
	responseMap2 := createTodo(t, todopayload2, http.StatusCreated)
	if responseMap2["message"] != "message2" {
		t.Fatal("expected message to be message2 received = ", responseMap2["message"])
	}
	if responseMap2["ID"] != 2.0 {
		t.Fatal("expected ID = 2 received = ", responseMap2["ID"])
	}

}

func TestCreateTodoBadRequest(t *testing.T) {

	//Clear all todos (already tested)
	err := clearTodos()
	if err != nil {
		t.Fatal("Clear all todos error " + err.Error())
	}

	badpayload := []byte(`{"message":1}`) //pass int(as message value) instead of string
	responseMap := createTodo(t, badpayload, http.StatusBadRequest)
	if responseMap["error"] != "bad request" {
		t.Fatal("expected error = \"bad request\" received = ", responseMap["error"])
	}

	badpayload = []byte(`{"message":"msg1"`) //Note the missing end paranthesis
	responseMap = createTodo(t, badpayload, http.StatusBadRequest)
	if responseMap["error"] != "bad request" {
		t.Fatal("expected error = \"bad request\" received = ", responseMap["error"])
	}

}

func TestExistentProduct(t *testing.T) {

	//Clear all todos (already tested)
	err := clearTodos()
	if err != nil {
		t.Fatal("Clear all todos error " + err.Error())
	}

	//presuming this is already tested in createtodo test, we dont check for proper response post create
	todopayload := []byte(`{"message":"message1"}`)
	responseMap := createTodo(t, todopayload, http.StatusCreated)

	//test existent product
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/todo/1", nil)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Expected status " + strconv.Itoa(http.StatusOK) +
			" Received status " + strconv.Itoa(resp.StatusCode))
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error while reading response " + err.Error())
	}

	// var responseMap map[string]string
	err = json.Unmarshal(response, &responseMap)
	if err != nil {
		t.Fatal("error while unmarshalling response " + err.Error())
	}

	if responseMap["message"] != "message1" {
		t.Fatal("Expceted message = \"message1\" Received = \"", responseMap["message"], "\"")
	}

	if responseMap["ID"] != 1.0 {
		t.Fatal("Expecting ID = 1 Received = ", responseMap["ID"])
	}

}

func TestDeleteProduct(t *testing.T) {

	//Clear all todos (already tested)
	err := clearTodos()
	if err != nil {
		t.Fatal("Clear all todos error " + err.Error())
	}

	//presuming this is already tested in createtodo test, we dont check for proper response post create
	todopayload := []byte(`{"message":"message1"}`)
	responseMap := createTodo(t, todopayload, http.StatusCreated)

	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/todo/1", nil)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Expected status " + strconv.Itoa(http.StatusOK) +
			" Received status " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error while reading response " + err.Error())
	}

	err = json.Unmarshal(response, &responseMap)
	if err != nil {
		t.Fatal("error while unmarshalling response " + err.Error())
	}

	if responseMap["result"] != "success" {
		t.Fatal("Expceted result = \"success\" Received = \"", responseMap["result"], "\"")
	}

	//test todo for non-existence
	req, err = http.NewRequest("GET", "http://localhost:8080/todo/1", nil)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	resp2, err := client.Do(req)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusNotFound {
		t.Fatal("Expected status " + strconv.Itoa(http.StatusNotFound) +
			" Received status " + strconv.Itoa(resp2.StatusCode))
	}

}

func TestUpdateProduct(t *testing.T) {

	//Clear all todos (already tested)
	err := clearTodos()
	if err != nil {
		t.Fatal("Clear all todos error " + err.Error())
	}

	//CREATE
	//presuming this is already tested in createtodo test, we dont check for proper response post create
	todopayload := []byte(`{"message":"message1"}`)
	responseMap := createTodo(t, todopayload, http.StatusCreated)

	//UPDATE
	client := http.Client{}
	updatepayload := []byte(`{"message":"message1-modified"}`)
	req, err := http.NewRequest("PUT", "http://localhost:8080/todo/1", bytes.NewReader(updatepayload))
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Expected status " + strconv.Itoa(http.StatusOK) +
			" Received status " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error while reading response " + err.Error())
	}

	err = json.Unmarshal(response, &responseMap)
	if err != nil {
		t.Fatal("error while unmarshalling response " + err.Error())
	}

	if responseMap["message"] != "message1-modified" {
		t.Fatal("Expceted message = \"message1-modified\" Received = \"", responseMap["message"], "\"")
	}

	// DOUBLE CHECK
	//test todo for update

	req, err = http.NewRequest("GET", "http://localhost:8080/todo/1", nil)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	resp2, err := client.Do(req)
	if err != nil {
		t.Fatal("Unexpected error " + err.Error())
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		t.Fatal("Expected status " + strconv.Itoa(http.StatusOK) +
			" Received status " + strconv.Itoa(resp2.StatusCode))
	}
	response2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		t.Fatal("error while reading response " + err.Error())
	}

	err = json.Unmarshal(response2, &responseMap)
	if err != nil {
		t.Fatal("error while unmarshalling response " + err.Error())
	}

	if responseMap["message"] != "message1-modified" {
		t.Fatal("Expceted message = \"message1-modified\" Received = \"", responseMap["message"], "\"")
	}

}
