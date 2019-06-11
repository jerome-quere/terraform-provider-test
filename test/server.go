package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

type Server struct {
	Address *string
}

var nextId = 0
var servers = map[string]*Server{}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	id := fmt.Sprintf("%d", nextId)
	nextId += 1

	address := d.Get("address").(string)
	servers[id] = &Server{
		Address: &address,
	}
	d.SetId(id)
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	if server, exist := servers[d.Id()]; exist {
		if server.Address != nil {
			d.Set("address", server.Address)
		} else {
			d.Set("address", nil)
		}
	} else {
		d.SetId("")
	}
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	if server, exist := servers[d.Id()]; exist {
		address := d.Get("address").(string)
		server.Address = &address
	} else {
		d.SetId("")
	}
	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	delete(servers, d.Id())
	return nil
}
