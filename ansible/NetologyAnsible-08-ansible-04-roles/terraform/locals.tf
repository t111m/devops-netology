locals {
  web_instance_each_map = {
    prod = {
      "clickhouse.netology.yc" = "node-clickhouse",
      "vector.netology.yc" = "node-vector",
      "lighthouse.netology.yc" = "node-lighthouse"
    }
  }
}