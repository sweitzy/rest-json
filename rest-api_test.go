package main

import (
	"testing"
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
	"encoding/json"
)

func TestSanity(t *testing.T) {

	var resp *http.Response
	var err error

/*
	// Test deleting
	fmt.Println("TEST DELETE")
	var req *http.Request
	req, err = http.NewRequest("DELETE", "http://example.com/2", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
*/


	// Test that we can get to index page.
	resp, err = http.Get("http://localhost:8081")
	if err != nil {
		t.Error(err)
	}

	// Test that response is JSON.
	if resp.Header["Content-Type"][0] != "application/json; charset=UTF-8" {
		t.Error("Expected content-type JSON, got", resp.Header["Content-Type"])
	}
	

	// Test that index page is correct.
	page, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Error(err)
	}
//	fmt.Printf("%s", page)
	if string(page) != `{"todos":"/todos/"}\n` {
		t.Error("Index was not correct.")
	}

	// Test that we can get list of TODOs.
	resp, err = http.Get("http://localhost:8081/todos")
	if err != nil {
		t.Error(err)
	}

/*
	// Test that number of pre-defined TODOs is correct
	todo_read := []Todo;
	err = json.NewDecoder(resp.Body).Decode(&todo_read)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Type: %T Value: %v", todo_read, todo_read)
*/

/*
	// Test that number of pre-defined TODOs is correct
	page, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s", page)
*/

	// Test that we can create a TODO.
	// Similar command-lne example:
	//   curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8081/todos
	todo := Todo{Name: "Pass interview"}
	fmt.Println("DEBUG", todo)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(todo)
	fmt.Println(b)
t.Error("DEBUG")
	resp, err = http.Post("http://localhost:8081/todos", "application/json", b)
	if err != nil {
		t.Error(err)
	}

/*
	page, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s", page)
*/
	
	// Test that new created TODO has proper data.
	todo_read := Todo{};
	err = json.NewDecoder(resp.Body).Decode(&todo_read)
	if err != nil {
		t.Error(err)
	}
	if todo != todo_read {
		t.Error("Created TODO was not proper")
	}
	fmt.Printf("Type: %T Value: %v", todo_read, todo_read)

/*
	resp, err = json.Unmarshal(body, &todo)

	page, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s", page)
	if string(page) != "Welcome!\n" {
		t.Error("Index was not Welcome!")
	}
*/
}
