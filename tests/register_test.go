package tests

import "testing"
import "net/http"

func TestApp(t *testing.T) {
	client := &http.Client{};

	req,err:=http.NewRequest("GET","http://127.0.0.1:8081/test",nil)
	if(err!=nil) {
		t.Error("Error making request!")
	}
	resp,err:=client.Do(req)
	if(err!=nil) {
		t.Error("Error executing request!")
	} else if(resp.StatusCode!=500) {
		t.Error("Received Status: ",resp.Status)
	}
}
