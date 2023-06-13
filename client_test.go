package onedev

import (
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewClient("https://code.onedev.io")
	_, _, err := client.GetProjects(&ProjectQueryOptions{})
	if err != nil {
		t.Errorf("Expected projects, received %v", err)
	}
}

func TestClientBasicAuth(t *testing.T) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	url := os.Getenv("URL")
	client := NewClient(url, SetBasicAuth(user, password))
	_, _, err := client.GetProjects(&ProjectQueryOptions{})
	if err != nil {
		t.Errorf("Expected projects, received %v", err)
	}
}

func TestClientBasicAuthWithToken(t *testing.T) {
	user := os.Getenv("USER")
	password := os.Getenv("TOKEN")
	url := os.Getenv("URL")
	client := NewClient(url, SetBasicAuth(user, password))
	_, _, err := client.GetProjects(&ProjectQueryOptions{})
	if err != nil {
		t.Errorf("Expected projects, received %v", err)
	}
}

func TestClientToken(t *testing.T) {
	token := os.Getenv("TOKEN")
	url := os.Getenv("URL")
	client := NewClient(url, SetToken(token))
	_, _, err := client.GetProjects(&ProjectQueryOptions{})
	if err != nil {
		t.Errorf("Expected projects, received %v", err)
	}
}
