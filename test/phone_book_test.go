package main

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders = map[string]terraform.ResourceProvider{
	"test": Provider(),
}

func TestAccTestPhoneBook(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccTestServerStep1,
				Check: resource.ComposeTestCheckFunc(
					func(state *terraform.State) error {
						fmt.Println(state)
						return nil
					},
					resource.TestCheckResourceAttr("test_phone_book.base", "contact.#", "1"),
					resource.TestCheckResourceAttr("test_phone_book.base", "contact.206327822.last_name", "DOE"),
					resource.TestCheckResourceAttr("test_phone_book.base", "contact.206327822.first_name", "john"),
				),
			},
		},
	})
}

var TestAccTestServerStep1 = `
resource "test_phone_book" "base" {
	contact {
      last_name = "doe"
      first_name = "john"
    }
}
`
