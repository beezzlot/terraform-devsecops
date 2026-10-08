output "disk_ids" {
  description = "IDs of additional disks"
  value       = module.disks.disk_ids
}

output "private_ips" {
  description = "Private IP addresses of VMs"
  value       = module.compute.private_ips
}
