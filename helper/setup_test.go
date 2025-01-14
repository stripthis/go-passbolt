package helper

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/speatzle/go-passbolt/api"
)

var client *api.Client

func TestMain(m *testing.M) {
	url := os.Getenv("REG_URL")
	fmt.Printf("Registering with url: %v\n", url)
	userID, token, err := ParseInviteUrl(url)
	if err != nil {
		panic(fmt.Errorf("Unable to Parse Invite URL: %w", err))
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	hc := &http.Client{Transport: tr}

	rc, err := api.NewClient(hc, "", "https://localhost", "", "")
	if err != nil {
		panic(fmt.Errorf("Creating Registration Client: %w", err))
	}

	ctx := context.TODO()

	privkey, err := SetupAccount(ctx, rc, userID, token, "password123")
	if err != nil {
		panic(fmt.Errorf("Setup Account: %w", err))
	}

	c, err := api.NewClient(hc, "", "https://localhost", privkey, "password123")
	if err != nil {
		panic(fmt.Errorf("Setup Client: %w", err))
	}

	c.Login(ctx)
	if err != nil {
		panic(fmt.Errorf("Login Client: %w", err))
	}

	client = c

	os.Exit(m.Run())
}
