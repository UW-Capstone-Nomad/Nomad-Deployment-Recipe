job "prometheus" {
  datacenters = ["dc1"]
  type        = "service"
  update {
    max_parallel     = 1
    min_healthy_time = "10s"
    healthy_deadline = "3m"
    auto_revert      = false
    canary           = 0
  }
  group "monitoring" {
    count = 1
    restart {
      attempts = 10
      interval = "5m"
      delay    = "25s"
      mode     = "delay"
    }
    network {
      port "prometheus_ui" {
        to = 9090
      }
      port "grafana_ui" {
        to = 3000
      }
    }

    service {
      name = "prometheus-ui"
      #tags = ["urlprefix-/prometheus"]
      tags = ["urlprefix-/prometheus strip=/prometheus"]
      port = "prometheus_ui"
      check {
        name     = "prometheus_ui port alive"
        type     = "tcp"
        interval = "10s"
        timeout  = "2s"
      }
    }

    service {
      name = "grafana-ui"
      port = "grafana_ui"
      tags = ["urlprefix-/grafana strip=/grafana"]
      check {
        name     = "grafana-ui port alive"
        type     = "tcp"
        interval = "10s"
        timeout  = "2s"
      }
    }
    ephemeral_disk { size = 1000 }
    task "grafana" {
      artifact {
        source      = "https://gist.githubusercontent.com/angrycub/046cee11bd3d8c4ab9a3819646c9660c/raw/c699095c2cb25b896e2c709da588b668ce82f8b5/prometheus_nomad.json"
        destination = "local/provisioning/dashboards/dashs"
      }
      template {
        change_mode = "noop"
        destination = "local/provisioning/dashboards/file_provider.yml"
        data        = <<EOH
apiVersion: 1
providers:
- name: 'default'
  orgId: 1
  folder: ''
  type: file
  disableDeletion: false
  updateIntervalSeconds: 10 #how often Grafana will scan for changed dashboards
  options:
    path: {{ env "NOMAD_TASK_DIR" }}/provisioning/dashboards/dashs
EOH

      }
      template {
        change_mode = "noop"
        destination = "local/provisioning/datasources/prometheus_datasource.yml"
        data        = <<EOH
apiVersion: 1
datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    url: http://{{ env "NOMAD_ADDR_prometheus_ui" }}
EOH
      }
      env {
        GF_SERVER_ROOT_URL    = "http://${NOMAD_ADDR_grafana_ui}"
        GF_PATHS_PROVISIONING = "/${NOMAD_TASK_DIR}/provisioning"
      }
      driver = "docker"
      config {
        image = "grafana/grafana:6.1.4"
        ports = ["grafana_ui"]
      }
    }

    task "prometheus" {
      template {
        change_mode = "noop"
        destination = "local/prometheus.yml"
        data        = <<EOH
---
global:
  scrape_interval: 15s
EOH
      }

      driver = "docker"
      config {
        image = "prom/prometheus:v2.9.1"
        args = [
          "--web.external-url=http://${NOMAD_ADDR_prometheus_ui}/prometheus",
          "--config.file=${NOMAD_TASK_DIR}/prometheus.yml",
        ]
        ports = ["prometheus_ui"]
      }
      resources {
        cpu    = 500
        memory = 256
      }
    }
  }
}
