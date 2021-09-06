package provider

import (
	"fmt"
	"testing"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccResourcePrivateKey = `
resource "gpg_private_key" "key_1" {
	name = "test"
	email = "test@test.com"
}
resource "gpg_private_key" "key_2" {
	name = "test"
	email = "test@test.com"
	passphrase = "this is not a secure passphrase"
	rsa_bits = 3072
}
`

func TestAccResourcePrivateKey(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePrivateKey,
				Check: resource.ComposeTestCheckFunc(
					testAccResourceCreateKey("gpg_private_key.key_1", ""),
				),
			},
			{
				Config: testAccResourcePrivateKey,
				Check: resource.ComposeTestCheckFunc(
					testAccResourceCreateKey("gpg_private_key.key_2", "this is not a secure passphrase"),
				),
			},
		},
	})
}

func testAccResourceCreateKey(id string, passphrase string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[id]
		if !ok {
			return fmt.Errorf("Not found: %s", id)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		// crypto.NewPlainMessageFromString("Verified message")
		privateKeyObj, err := crypto.NewKeyFromArmored(rs.Primary.Attributes["private_key"])
		if err != nil {
			return err
		}

		if passphrase != "" {
			privateKeyObj, err = privateKeyObj.Unlock([]byte(passphrase))
			if err != nil {
				return err
			}
		}

		signingKeyRing, err := crypto.NewKeyRing(privateKeyObj)
		if err != nil {
			return err
		}

		if ok := signingKeyRing.CanEncrypt(); !ok {
			return fmt.Errorf("The primary key cannot encrypt. There is a problem.")
		}

		return nil
	}
}
