package subscriptions

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-bytes/client"
)

func datasourceOrder() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceOrderRead,

		Schema: map[string]*schema.Schema{
			"order_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"contract_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"po_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	orderID := d.Get("order_id").(string)
	order, err := c.GetOrderDetails(orderID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to get order with id %s: %s", orderID, err))
	}

	d.SetId(fmt.Sprintf("%d", order.ID))
	d.Set("id", order.ID)
	d.Set("contract_name", order.ContractName)
	d.Set("create_date", order.CreateDate)

	if len(order.Items) > 0 {
		d.Set("subscription_id", order.Items[0].SubscriptionID)
		d.Set("friendly_name", order.Items[0].FriendlyName)
		d.Set("po_number", order.Items[0].PONumber)
	}

	return nil
}
