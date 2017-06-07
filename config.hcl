option {
  rollback = true
  security_group = [
    "sg-000001",
    "sg-000002"
  ]
}

instance {
  tags {
    application = "munisystem-app"
    environment = "production"
  }
  class = "c4.xlarge"
}

# dns "route53" {
#   record = "munisystem.net"
#   ttl = 60
# }

dns "dnsimple" {
  recordID = 1
  recordName = "munisystem.net"
  ttl = 60
}
