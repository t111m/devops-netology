Role Name
=========

Lighthouse install role

Requirements
------------

Role install nginx and git by dependence

Role Variables
--------------

version: "master"

Dependencies
------------

geerlingguy.nginx
geerlingguy.git

Example Playbook
----------------

- name: Install Lighthouse
  hosts: lighthouse
  tags: lighthouse
  roles:
    - lighthouse-role

License
-------

MIT

Author Information
------------------


