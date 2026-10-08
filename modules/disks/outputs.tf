output "disk_ids" {
  description = "Additional disk IDs indexed by disk name"
  value       = { for name, disk in yandex_compute_disk.extra_disks : name => disk.id }
}
