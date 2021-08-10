package provider

import (
	"context"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePrivateKey() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description:   "The resource `gpg_private_key` generates a GPG private/public key pair in ASCII-armored format.",
		CreateContext: resourcePrivateKeyCreate,
		ReadContext:   resourcePrivateKeyRead,
		DeleteContext: resourcePrivateKeyDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name attached to the key.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"email": {
				Description: "Email attached to the key.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"rsa_bits": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of bits to use when generating RSA key.",
				ForceNew:    true,
				Default:     4096,
			},

			"private_key": {
				Type:        schema.TypeString,
				Description: "Generated private key in ASCII-armored format.",
				Computed:    true,
				Sensitive:   true,
			},

			"public_key": {
				Type:        schema.TypeString,
				Description: "Generated public key in ASCII-armored format.",
				Computed:    true,
			},

			"fingerprint": {
				Type:        schema.TypeString,
				Description: "Public key fingerprint.",
				Computed:    true,
			},
		},
	}
}

func resourcePrivateKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	rsa_bits := d.Get("rsa_bits").(int)
	name := d.Get("name").(string)
	email := d.Get("email").(string)

	key, err := crypto.GenerateKey(name, email, "rsa", rsa_bits)
	if err != nil {
		return diag.FromErr(err)
	}

	publicKey, err := key.GetArmoredPublicKey()
	if err != nil {
		return diag.FromErr(err)
	}

	privateKey, err := key.Armor()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(key.GetFingerprint())
	d.Set("private_key", string(privateKey))
	d.Set("public_key", publicKey)
	d.Set("fingerprint", key.GetFingerprint())

	return nil
}

func resourcePrivateKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return nil
}

func resourcePrivateKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	d.SetId("")
	return nil
}
