resource "local_file" "inventory" {
  content = <<-DOC
    ---
    clickhouse:
      hosts:
        node-clickhouse:
          ansible_host: ${yandex_compute_instance.vm_for_each["clickhouse.netology.yc"].network_interface.0.nat_ip_address}
    vector:
      hosts:
        node-vector:
          ansible_host: ${yandex_compute_instance.vm_for_each["vector.netology.yc"].network_interface.0.nat_ip_address}
    lighthouse:
      hosts:
        node-lighthouse:
          ansible_host: ${yandex_compute_instance.vm_for_each["lighthouse.netology.yc"].network_interface.0.nat_ip_address}
    DOC
  filename = "../playbook/inventory/prod.yml"

  depends_on = [
    yandex_compute_instance.vm_for_each
  ]
}