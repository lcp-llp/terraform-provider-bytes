package subscriptions

import (
	"context"
	"fmt"
	"terraform-provider-bytes/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSubscription() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSubscriptionCreate,
		ReadContext:   resourceSubscriptionRead,
		DeleteContext: resourceSubscriptionDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"po_number": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"default_admin": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
				ForceNew: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSubscriptionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(*client.Client)

	// Prepare the JSON Data for API Payload
	subscriptionDetails := client.SubscriptionDetails{
		FriendlyName: d.Get("friendly_name").(string),
		PONumber:     d.Get("po_number").(string),
		PrincipalID:  d.Get("default_admin").(string),
	}

	// Call the function create the subscription with payload
	subscription, err := c.CreateSubscription(subscriptionDetails)

	// If error, print it out
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Something wrong with Provider to Delete record: %s", err),
			Detail:   fmt.Sprintf("Something wrong with Provider to create record: %s", err),
		})
		return diags
	}

	d.SetId(fmt.Sprintf("%d", subscription.ID))
	d.Set("contract_name", subscription.ContractName)
	d.Set("create_date", subscription.CreateDate)

	if len(subscription.Items) > 0 {
		d.Set("subscription_id", subscription.Items[0].SubscriptionID)
		d.Set("friendly_name", subscription.Items[0].FriendlyName)
		d.Set("po_number", subscription.Items[0].PONumber)
		d.Set("default_admin", subscription.Items[0].PrincipalID)
	}
	// Call the resourceARecordRead function
	resourceSubscriptionRead(ctx, d, m)

	return nil
}
func resourceSubscriptionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
func resourceSubscriptionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// No-op, do nothing when deleting, not currently supported by Bytes API
	return nil
}
