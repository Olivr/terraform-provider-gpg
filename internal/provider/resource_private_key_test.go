package provider

import (
	"fmt"
	"testing"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccResourcePrivateKey = `
resource "gpg_private_key" "key" {
	name = "test"
	email = "test@test.com"
}
`

const testAccResourcePrivateKeyRsaBits = `
resource "gpg_private_key" "key" {
	name = "test"
	email = "test@test.com"
	rsa_bits = 3072
}
`

const testAccResourcePrivateKeyPassphrase = `
resource "gpg_private_key" "key" {
	name = "test"
	email = "test@test.com"
	passphrase = "this is not a secure passphrase"
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
					testAccResourceCreateKey("gpg_private_key.key", 2048, ""),
				),
			},
			{
				Config: testAccResourcePrivateKeyRsaBits,
				Check: resource.ComposeTestCheckFunc(
					testAccResourceCreateKey("gpg_private_key.key", 3072, ""),
				),
			},
			{
				Config: testAccResourcePrivateKeyPassphrase,
				Check: resource.ComposeTestCheckFunc(
					testAccResourceCreateKey("gpg_private_key.key", 2048, "this is not a secure passphrase"),
				),
			},
		},
	})
}

func testAccResourceCreateKey(id string, rsaBits uint16, passphrase string) resource.TestCheckFunc {
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

		keyLength, err := privateKeyObj.GetEntity().PrimaryKey.BitLength()
		if err != nil {
			return err
		}
		if keyLength != rsaBits {
			return fmt.Errorf("The length of the generated key (%d) doesn't match with the length requested in the Terraform configuration (%d).", keyLength, rsaBits)
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
