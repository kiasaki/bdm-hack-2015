variable "do_token" {}
variable "public_key" {}
variable "private_key" {}

provider "digitalocean" {
  token = "${var.do_token}"
}

resource "digitalocean_ssh_key" "default" {
  name = "tf-bdm"
  public_key = "${file("${var.public_key}")}"
}

resource "digitalocean_droplet" "workbench" {
  image = "freebsd-10-1-x64"
  name = "workbench.bdmhack"
  region = "nyc3"
  size = "2gb"
  ssh_keys = ["${digitalocean_ssh_key.default.id}"]
  count = 1
}

resource "digitalocean_droplet" "database" {
  image = "freebsd-10-1-x64"
  name = "database${count.index+1}.bdmhack"
  region = "nyc3"
  size = "4gb"
  ssh_keys = ["${digitalocean_ssh_key.default.id}"]
  count = 5
}

resource "digitalocean_droplet" "producer" {
  image = "freebsd-10-1-x64"
  name = "producer${count.index+1}.bdmhack"
  region = "nyc3"
  size = "1gb"
  ssh_keys = ["${digitalocean_ssh_key.default.id}"]
  count = 10
}

output "workbench" {
  value = "${digitalocean_droplet.workbench.ipv4_address}"
}

output "database1" {
  value = "${digitalocean_droplet.database.0.ipv4_address}"
}
output "database2" {
  value = "${digitalocean_droplet.database.1.ipv4_address}"
}
output "database3" {
  value = "${digitalocean_droplet.database.2.ipv4_address}"
}
output "database4" {
  value = "${digitalocean_droplet.database.3.ipv4_address}"
}
output "database5" {
  value = "${digitalocean_droplet.database.4.ipv4_address}"
}

output "producer1" {
  value = "${digitalocean_droplet.producer.0.ipv4_address}"
}
output "producer2" {
  value = "${digitalocean_droplet.producer.1.ipv4_address}"
}
output "producer3" {
  value = "${digitalocean_droplet.producer.2.ipv4_address}"
}
output "producer4" {
  value = "${digitalocean_droplet.producer.3.ipv4_address}"
}
output "producer5" {
  value = "${digitalocean_droplet.producer.4.ipv4_address}"
}
output "producer6" {
  value = "${digitalocean_droplet.producer.5.ipv4_address}"
}
output "producer7" {
  value = "${digitalocean_droplet.producer.6.ipv4_address}"
}
output "producer8" {
  value = "${digitalocean_droplet.producer.7.ipv4_address}"
}
output "producer9" {
  value = "${digitalocean_droplet.producer.8.ipv4_address}"
}
output "producer10" {
  value = "${digitalocean_droplet.producer.9.ipv4_address}"
}
