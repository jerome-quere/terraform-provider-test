package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders = map[string]terraform.ResourceProvider{
	"test": Provider(),
}

func TestAccTestServer(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccTestServerStep1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("test_server.base", "address", "1.1.1.1"),
				),
			},
			{
				Config: TestAccTestServerStep2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("test_server.base", "address", ""),
				),
			},
		},
	})
}

var TestAccTestServerStep1 = `
resource "test_server" "base" {
	address = "1.1.1.1"
}
`

var TestAccTestServerStep2 = `
resource "test_server" "base" {
}
`
