package tests

import "testing"
import "net/http"
import "math/rand"

func TestApp(t *testing.T) {
	client := &http.Client{};

	req,err:=http.NewRequest("GET","http://127.0.0.1:80/test",nil)
	if err!=nil {
		t.Error("Error making request: ",err)
	}
	resp,err:=client.Do(req)
	if err!=nil {
		t.Error("Error executing request: ",err)
	} else if resp.StatusCode!=500 {
		t.Error("Received Status: ",resp.Status,"\nMessage Body: ",resp.Body)
	}
}

func TestDistribution(t *testing.T) {
	min,max := 0,200
	r := rand.New(rand.NewSource(20))
	for i:=0;i<10000;i++ {
		sample := int(r.Float64()) * (max-min) + min
		if sample < min || sample > max {
			t.Error(sample)
		}
	}
}
