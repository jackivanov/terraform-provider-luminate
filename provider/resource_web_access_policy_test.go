package provider

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

const resourceWebAccessPolicy_enabled = `
	resource "luminate_web_access_policy" "new-web-access-policy" {
		enabled = "true"
		name =  "resourceWebAccessPolicy_enabled"
		identity_provider_id = "local"

		user_ids = ["c352709b-29e9-430c-a861-481944d4a3ae"]
  		applications = ["aeb7d51e-0934-459d-bc35-4d06e9b9f6a1"]
	}
`

const resourceWebAccessPolicy_disabled = `
	resource "luminate_web_access_policy" "new-web-access-policy" {
		enabled = "false"
  		name =  "resourceWebAccessPolicy_disabled"
		identity_provider_id = "local"

  		user_ids = ["c352709b-29e9-430c-a861-481944d4a3ae"]
  		applications = ["aeb7d51e-0934-459d-bc35-4d06e9b9f6a1"]
	}
`

const resourceWebAccessPolicy_enabled_not_specified = `
	resource "luminate_web_access_policy" "new-web-access-policy" {
  		name =  "resourceWebAccessPolicy_enabled_not_specified"
		identity_provider_id = "local"

  		user_ids = ["c352709b-29e9-430c-a861-481944d4a3ae"]
  		applications = ["aeb7d51e-0934-459d-bc35-4d06e9b9f6a1"]
	}
`

const resourceWebAccessPolicy_conditions_specified = `
	resource "luminate_web_access_policy" "new-web-access-policy" {
  		name =  "resourceWebAccessPolicy_conditions_specified"
		identity_provider_id = "local"

  		user_ids = ["c352709b-29e9-430c-a861-481944d4a3ae"]
  		applications = ["aeb7d51e-0934-459d-bc35-4d06e9b9f6a1"]

		conditions {
    		source_ip = ["127.0.0.1/24", "1.1.1.1/16"]
    		location = ["Wallis and Futuna"]

    		managed_device {
      			opswat = true
      			symantec_cloudsoc = true
      			symantec_web_security_service = true
    		}
  		}

	}
`

const resourceWebAccessPolicy_conditions_specified_update = `
	resource "luminate_web_access_policy" "new-web-access-policy" {
		name =  "resourceWebAccessPolicy_conditions_specified"
		identity_provider_id = "local"
	
		user_ids = ["c352709b-29e9-430c-a861-481944d4a3ae"]
		applications = ["aeb7d51e-0934-459d-bc35-4d06e9b9f6a1"]
	
		conditions {
			source_ip = ["127.0.0.1/24", "1.1.1.1/16"]
			location = ["Wallis and Futuna"]
	
			managed_device {
				opswat = false
				symantec_cloudsoc = true
				symantec_web_security_service = true
			}
		}

	}
`

const resourceWebAccessPolicy_validators_specified = `
	resource "luminate_web_access_policy" "new-web-access-policy" {
  		name =  "resourceWebAccessPolicy_validators_specified"
		identity_provider_id = "local"

  		user_ids = ["c352709b-29e9-430c-a861-481944d4a3ae"]
  		applications = ["aeb7d51e-0934-459d-bc35-4d06e9b9f6a1"]
	}
`

func TestAccLuminateWebAccessPolicy(t *testing.T) {
	resourceName := "luminate_web_access_policy.new-web-access-policy"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				PreConfig: testAccSleep,
				Config:    resourceWebAccessPolicy_enabled,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "resourceWebAccessPolicy_enabled"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			{
				PreConfig: testAccSleep,
				Config:    resourceWebAccessPolicy_disabled,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "resourceWebAccessPolicy_disabled"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			{
				PreConfig: testAccSleep,
				Config:    resourceWebAccessPolicy_enabled_not_specified,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "resourceWebAccessPolicy_enabled_not_specified"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			{
				PreConfig: testAccSleep,
				Config:    resourceWebAccessPolicy_conditions_specified,
				Destroy:   false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "resourceWebAccessPolicy_conditions_specified"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.source_ip.0", "127.0.0.1/24"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.source_ip.1", "1.1.1.1/16"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.location.0", "Wallis and Futuna"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.managed_device.0.opswat", "true"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.managed_device.0.symantec_cloudsoc", "true"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.managed_device.0.symantec_web_security_service", "true"),
				),
			},
			{
				PreConfig: testAccSleep,
				Config:    resourceWebAccessPolicy_conditions_specified_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "resourceWebAccessPolicy_conditions_specified"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.source_ip.0", "127.0.0.1/24"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.source_ip.1", "1.1.1.1/16"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.location.0", "Wallis and Futuna"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.managed_device.0.opswat", "false"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.managed_device.0.symantec_cloudsoc", "true"),
					resource.TestCheckResourceAttr(resourceName, "conditions.0.managed_device.0.symantec_web_security_service", "true"),
				),
			},
			{
				PreConfig: testAccSleep,
				Config:    resourceWebAccessPolicy_validators_specified,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "resourceWebAccessPolicy_validators_specified"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
		},
	})
}