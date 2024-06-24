package tf6muxprovider

import (
	log "github.com/sirupsen/logrus"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
)

//nolint:unused // Test provider server, executed by test framework
func main() {
	provider, err := New()

	if err != nil {
		log.Fatalf("unable to create provider: %s", err)
	}

	err = tf6server.Serve("registry.terraform.io/hashicorp/corner", provider)

	if err != nil {
		log.Fatalf("unable to serve provider: %s", err)
	}
}
