terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.27.0"
    }
  }
}

provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
}

resource "google_container_cluster" "primary" {
  name                = "runners-cluster"
  location            = var.zone
  initial_node_count  = 3
  deletion_protection = false
  node_config {
    disk_size_gb = 10
    machine_type = "e2-micro"
  }
}

resource "google_sql_database" "database" {
  name     = "runners_db"
  instance = google_sql_database_instance.instance.name
}

resource "google_sql_database_instance" "instance" {
  name                = "database-instance"
  root_password       = var.root_password
  deletion_protection = false
  database_version    = "POSTGRES_15"

  settings {
    tier              = "db-f1-micro"
    disk_type         = "PD_HDD"
    availability_type = "ZONAL"
    edition           = "ENTERPRISE"
    backup_configuration {
      enabled = false
    }
    data_cache_config {
      data_cache_enabled = false
    }
  }
}
