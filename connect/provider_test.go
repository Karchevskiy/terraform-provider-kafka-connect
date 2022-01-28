package connect

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testProvider *schema.Provider
var testProviders map[string]*schema.Provider

func init() {
	testProvider = Provider()
	testProviders = map[string]*schema.Provider{
		"kafka-connect": testProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	connectVar := "KAFKA_CONNECT_URL"
	urlValue := os.Getenv(connectVar)
	if urlValue == "" {
		t.Fatalf("%s env var must be set", connectVar)
	}

	connectToken := "KAFKA_CONNECT_TOKEN"
	tokenValue := os.Getenv(connectToken)
	if tokenValue == "" {
		t.Fatalf("%s env var must be set", connectToken)
	}
}
