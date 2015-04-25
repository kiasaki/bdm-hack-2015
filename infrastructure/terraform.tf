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
  name = "database.${count.index+1}.bdmhack"
  region = "nyc3"
  size = "2gb"
  ssh_keys = ["${digitalocean_ssh_key.default.id}"]
  count = 5
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
