output "subscription_output" {
  value = [data.bytes_order.example.subscription_id, data.bytes_order.example.friendly_name]
}