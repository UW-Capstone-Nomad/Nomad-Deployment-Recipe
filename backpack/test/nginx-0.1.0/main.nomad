job "nginx" {
	datacenters = {{ toJson .datacenters }}
	type = "service"

	group "nginx_servers" {
		task "nginx" {
			driver = "docker"
			config {
				image = "nginx:alpine"
				ports = ["http"{{ if .nginx.https.enable }}, "https" {{ end }}]
			}
		}

		network {
			port "http" {
				static = {{ .nginx.http.port }}
			}
			{{- if .nginx.https.enable }} 
			port "https" {
				static = {{ .nginx.https.port }}
			}
			{{- end }}
		}
  }
}