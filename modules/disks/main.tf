resource "yandex_compute_disk" "extra_disks" {
  for_each = toset(var.disk_names)

  name = each.value
  type = var.disk_type
  size = var.disk_size
  zone = var.zone
}
