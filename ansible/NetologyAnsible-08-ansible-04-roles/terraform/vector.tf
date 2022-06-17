resource "local_file" "vector" {
  content = <<-DOC
    [sources.my_source_id]
    type = "syslog"
    address = "0.0.0.0:514"
    mode = "tcp"
    path = "/usr/lib/systemd/system/syslog.socket"

    [sinks.my_sink_id]
    type = "clickhouse"
    inputs = [ "my_source_id" ]
    endpoint = "http://${yandex_compute_instance.vm_for_each["clickhouse.netology.yc"].network_interface.0.nat_ip_address}:8123"
    database = "logs"
    table = "access_logs"
    skip_unknown_fields = true
    DOC
  filename = "../playbook/roles/vector-role/templates/vector.toml"

  depends_on = [
    local_file.inventory
  ]
}