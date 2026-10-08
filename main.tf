module "disks" {
  source = "./modules/disks"

  disk_names = var.extra_disk_names
  disk_type  = var.extra_disk_type
  disk_size  = var.extra_disk_size
  zone       = "ru-central1-a"
}

module "compute" {
  source = "./modules/compute"

  vm_specification = var.vm_specification
  ssh_public_key   = var.ssh_public_key
  extra_disk_ids   = module.disks.disk_ids
  subnet_name      = "default-ru-central1-a"
  zone             = "ru-central1-a"

  depends_on = [module.disks]
}
