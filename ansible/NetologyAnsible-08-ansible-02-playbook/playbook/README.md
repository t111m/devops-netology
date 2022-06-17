### Preparation for launch playbook
1) You need create docker container with Dockerfile
    ```bash
    $ docker build -t keqpup232/centos:repo .
    ```
2) Then run docker container centos7 
    ```bash
   $ docker run -dt --name centos7 keqpup232/centos:repo  
   ```
3) And play ansible site.yml with inventory prod.yml
    ```bash
   $ ansible-playbook --diff -i inventory/prod.yml site.yml
   ```
   1) If you need change distrib version, set that variables in group_vars
       ```yaml
       vector_version: "0.21.1"
       clickhouse_version: "22.3.3.44"
       ```
   2) Use tag vector if you need only install vector 
       ```bash
       $ ansible-playbook --diff --tags vector -i inventory/prod.yml site.yml
       ```
### Ansible actions in docker centos7
1) First tasks "Install Clickhouse"
   1) Download clickhouse_packages:
      - clickhouse-client
      - clickhouse-server
      - clickhouse-common-static
   2) Install clickhouse packages
   3) Start clickhouse service
   4) Create database
2) Second tasks "Install Vector"
   1) Download vector distrib
   2) Create directory for Vector
   3) Install unzip and tar unarchive
   4) Extract vector
   5) Create a symbolic link
   6) Handlers Check is all done



