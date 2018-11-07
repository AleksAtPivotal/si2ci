provider "google" {
  credentials = "${file(var.gcp_service_account)}"
  project     = "fe-asaul"
  region      = "us-east1"
  version = "1.19"
}

resource "google_compute_instance" "concourse" {
  name           = "concourse-demoapp-initializr"
  count          = "${var.instance_count}"
  machine_type   = "${var.machine_type}"
  can_ip_forward = false
  zone           = "${var.zone}"

  boot_disk {
    initialize_params {
      image = "coreos-stable"
      type  = "${var.disk_type}"
      size  = "${var.disk_size}"
    }
  }

  network_interface {
    subnetwork = "${var.subnetwork_name}"

    access_config = {
      // Ephemeral IP
    }
  }

  tags = ["concourse"]

  metadata = {
    //user-data = "${data.ignition_config.concourse.rendered}"
    sshKeys   = "core:${file(var.public_ssh_key)}"
  }
}
/*
# Systemd unit data resource containing the unit definition
data "ignition_systemd_unit" "example" {
  name    = "example.service"
  content = "[Service]\nType=oneshot\nExecStart=/usr/bin/echo Hello World\n\n[Install]\nWantedBy=multi-user.target"
}

data "ignition_user" "core" {
  name                = "core"
  home_dir            = "/home/core/"
  ssh_authorized_keys = "${file(var.public_ssh_key)}"
}

# Ingnition config include the previous defined systemd unit data resource
data "ignition_config" "concourse" {
  systemd = [
    "${data.ignition_systemd_unit.example.id}",
  ]
}
*/