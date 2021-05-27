client {
  enabled = true
  host_volume "minio-data" {
    path = "/opt/volumes/minio-data"
    read_only = false
  }
}
