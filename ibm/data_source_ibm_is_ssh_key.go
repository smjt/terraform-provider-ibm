package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
)

func dataSourceIBMISSSHKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISSSHKeyRead,

		Schema: map[string]*schema.Schema{
			isKeyName: {
				Type:     schema.TypeString,
				Required: true,
			},
			isKeyType: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isKeyFingerprint: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isKeyLength: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			isVPCResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},
		},
	}
}

func dataSourceIBMISSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	keyC := compute.NewKeyClient(sess)

	name := d.Get(isKeyName).(string)

	keys, _, err := keyC.List("")
	if err != nil {
		return err
	}

	for _, key := range keys {
		if key.Name == name {
			d.SetId(key.ID.String())
			d.Set(isKeyName, key.Name)
			d.Set(isKeyType, key.Type)
			d.Set(isKeyFingerprint, key.Fingerprint)
			d.Set(isKeyLength, key.Length)
			controller, err := getBaseController(meta)
			if err != nil {
				return err
			}
			if sess.Generation == 1 {
				d.Set(isVPCResourceControllerURL, controller+"/vpc/compute/sshKeys")
			} else {
				d.Set(isVPCResourceControllerURL, controller+"/vpc-ext/compute/sshKeys")
			}

			return nil
		}
	}
	return fmt.Errorf("No SSH Key found with name %s", name)
}
