package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePhoneBook() *schema.Resource {
	return &schema.Resource{
		Create: resourcePhoneBookCreate,
		Read:   resourcePhoneBookRead,
		Update: resourcePhoneBookUpdate,
		Delete: resourcePhoneBookDelete,

		Schema: map[string]*schema.Schema{
			"contact": {
				Type:     schema.TypeSet,
				Set:      contactHash,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_name": {
							Type:     schema.TypeString,
							Required: true,
							StateFunc: func(i interface{}) string {
								str := i.(string)
								return strings.ToUpper(str)
							},
						},
						"first_name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

type Contact struct {
	lastName  string
	firstName string
}

type PhoneBook struct {
	contacts map[int]*Contact
}

var nextId = 0
var phoneBooks = map[string]*PhoneBook{}

func resourcePhoneBookCreate(d *schema.ResourceData, m interface{}) error {
	id := fmt.Sprintf("%d", nextId)
	nextId += 1

	phoneBook := &PhoneBook{
		contacts: map[int]*Contact{},
	}

	contacts := d.Get("contact").(*schema.Set)

	//*****************************************************************************
	//                                 ISSUE HERE
	//
	// In our test we only have one contact. We should never enter this condition
	//*****************************************************************************
	if contacts.Len() != 1 {
		for _, rawContact := range contacts.List() {
			fmt.Printf("Got contact with hash %q => %v\n", contactHash(rawContact), rawContact)
		}
		panic("Should only have one contact in set")
	}

	for _, rawContact := range contacts.List() {
		phoneBook.contacts[contactHash(rawContact)] = &Contact{
			lastName:  rawContact.(map[string]interface{})["last_name"].(string),
			firstName: rawContact.(map[string]interface{})["first_name"].(string),
		}
	}

	phoneBooks[id] = phoneBook
	d.SetId(id)
	return resourcePhoneBookRead(d, m)
}

func resourcePhoneBookRead(d *schema.ResourceData, m interface{}) error {
	if phoneBook, exist := phoneBooks[d.Id()]; exist {
		contacts := schema.NewSet(contactHash, nil)
		for _, c := range phoneBook.contacts {
			contacts.Add(map[string]interface{}{
				"last_name":  c.lastName,
				"first_name": c.firstName,
			})
		}
		d.Set("contact", contacts)
	} else {
		d.SetId("")
	}
	return nil
}

func resourcePhoneBookUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePhoneBookDelete(d *schema.ResourceData, m interface{}) error {
	delete(phoneBooks, d.Id())
	return nil
}

func contactHash(d interface{}) int {
	r := d.(map[string]interface{})
	s := fmt.Sprintf("%s-%s", r["last_name"], r["first_name"])
	return schema.HashString(s)
}
