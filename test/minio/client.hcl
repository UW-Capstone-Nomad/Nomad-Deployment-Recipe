client {
  enabled = true
  host_volume "my-website-db" {
    path = "/opt/volumes/my-website-db"
    read_only = false
  }
}

