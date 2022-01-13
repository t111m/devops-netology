# Курсовая работа по итогам модуля "DevOps и системное администрирование"

## Задание

1. Создайте виртуальную машину Linux.
![ScreenShot](kr1.jpg)

2. Установите ufw и разрешите к этой машине сессии на порты 22 и 443, при этом трафик на интерфейсе localhost (lo) должен ходить свободно на все порты.
```shell
tim@tim:~$ sudo ufw status
Status: inactive
tim@tim:~$ sudo apt-get install ufw
Reading package lists... Done
Building dependency tree
Reading state information... Done
The following packages will be upgraded:
  ufw
1 upgraded, 0 newly installed, 0 to remove and 44 not upgraded.
Need to get 147 kB of archives.
After this operation, 3,072 B of additional disk space will be used.
Get:1 http://lu.archive.ubuntu.com/ubuntu focal-updates/main amd64 ufw all 0.36-6ubuntu1 [147 kB]
Fetched 147 kB in 13s (11.1 kB/s)
Preconfiguring packages ...
(Reading database ... 71558 files and directories currently installed.)
Preparing to unpack .../ufw_0.36-6ubuntu1_all.deb ...
Unpacking ufw (0.36-6ubuntu1) over (0.36-6) ...
Setting up ufw (0.36-6ubuntu1) ...
Processing triggers for man-db (2.9.1-1) ...
Processing triggers for rsyslog (8.2001.0-1ubuntu1.1) ...
Processing triggers for systemd (245.4-4ubuntu3.11) ...

tim@tim:~$ sudo ufw enable
Command may disrupt existing ssh connections. Proceed with operation (y|n)? y
Firewall is active and enabled on system startup
tim@tim:~$ sudo ufw status
Status: active
tim@tim:~$ sudo ufw app list
Available applications:
  OpenSSH
tim@tim:~$ sudo ufw allow OpenSSH
Rule added
Rule added (v6)
tim@tim:~$ sudo ufw allow https
Rule added
Rule added (v6)
tim@tim:~$ sudo ufw status
Status: active

To                         Action      From
--                         ------      ----
OpenSSH                    ALLOW       Anywhere
443/tcp                    ALLOW       Anywhere
OpenSSH (v6)               ALLOW       Anywhere (v6)
443/tcp (v6)               ALLOW       Anywhere (v6)

```


3. Установите hashicorp vault ([инструкция по ссылке](https://learn.hashicorp.com/tutorials/vault/getting-started-install?in=vault/getting-started#install-vault)).
```shell
tim@tim:~$ vault --version
Vault v1.9.2 (f4c6d873e2767c0d6853b5d9ffc77b0d297bfbdf)
```
4. Cоздайте центр сертификации по инструкции ([ссылка](https://learn.hashicorp.com/tutorials/vault/pki-engine?in=vault/secrets-management)) и выпустите сертификат для использования его в настройке веб-сервера nginx (срок жизни сертификата - месяц).
```shell
tim@tim:~$ export VAULT_ADDR=http://127.0.0.1:8200
tim@tim:~$ export VAULT_TOKEN=root
tim@tim:~$ vault secrets enable pki
Success! Enabled the pki secrets engine at: pki/
tim@tim:~$ vault secrets tune -max-lease-ttl=87600h pki
Success! Tuned the secrets engine at: pki/
tim@tim:~$ vault write -field=certificate pki/root/generate/internal \
>      common_name="example.com" \
>      ttl=87600h > CA_cert.crt
tim@tim:~$ vault write pki/config/urls \
>      issuing_certificates="$VAULT_ADDR/v1/pki/ca" \
>      crl_distribution_points="$VAULT_ADDR/v1/pki/crl"
Success! Data written to: pki/config/urls
tim@tim:~$ vault secrets enable -path=pki_int pki
Success! Enabled the pki secrets engine at: pki_int/
tim@tim:~$ vault secrets tune -max-lease-ttl=43800h pki_int
Success! Tuned the secrets engine at: pki_int/
tim@tim:~$ vault write -format=json pki_int/intermediate/generate/internal \
>      common_name="example.com Intermediate Authority" \
>      | jq -r '.data.csr' > pki_intermediate.csr
tim@tim:~$ vault write -format=json pki/root/sign-intermediate csr=@pki_intermed                                                                                                                                                             iate.csr \
>      format=pem_bundle ttl="43800h" \
>      | jq -r '.data.certificate' > intermediate.cert.pem
tim@tim:~$ vault write pki_int/intermediate/set-signed certificate=@intermediate                                                                                                                                                             .cert.pem
Success! Data written to: pki_int/intermediate/set-signed
tim@tim:~$ vault write pki_int/roles/example-dot-com \
>      allowed_domains="example.com" \
>      allow_subdomains=true \
>      max_ttl="740h"
Success! Data written to: pki_int/roles/example-dot-com
tim@tim:~$ vault write pki_int/issue/example-dot-com common_name="test.example.c                                                                                                                                                             om" ttl="720h"
Key                 Value
---                 -----
ca_chain            [-----BEGIN CERTIFICATE-----
MIIDpjCCAo6gAwIBAgIUK3ks87DCjns2g7WAKLmo1zdM1CQwDQYJKoZIhvcNAQEL
BQAwFjEUMBIGA1UEAxMLZXhhbXBsZS5jb20wHhcNMjIwMTEzMDc0NzE1WhcNMjcw
MTEyMDc0NzQ1WjAtMSswKQYDVQQDEyJleGFtcGxlLmNvbSBJbnRlcm1lZGlhdGUg
QXV0aG9yaXR5MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsRfeyN+F
lAe+fLhWVxDfsb23oCprvb5hgAtLeVGlMeSDJ2UeXjwrxYZzR8H9/bhIPfi+QMz5
bsFcpQL/MgLkyGNcv37+hTk/CFDzNRtnXQSIULe7h3pcZ6zg1NsBDhWPXiHdAy6H
4aL6yhRyRr4DHe0Xea429Xq2VD7nZw2WNKnrTCCaOIigezZqYfGTSTvuHdvDSvLD
jlN4JfJUv07wzk2D462QOtiWDPY+kxs0v8zo/5JA4vPocihuUIajLfTdJeLFzz/9
q4d7K1gSMe0zfPrFhLA/ahUMpwqp3krB4xGk6YnKOuvVUDAqNFknGadYVwU0cvne
EQdmg/mVDM/L0wIDAQABo4HUMIHRMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E
BTADAQH/MB0GA1UdDgQWBBSGiiB8s/wrHZ5PJp5ZJ72wDXfEzTAfBgNVHSMEGDAW
gBQdLCzsiHUdOri9k5nJxZaevsXbdzA7BggrBgEFBQcBAQQvMC0wKwYIKwYBBQUH
MAKGH2h0dHA6Ly8xMjcuMC4wLjE6ODIwMC92MS9wa2kvY2EwMQYDVR0fBCowKDAm
oCSgIoYgaHR0cDovLzEyNy4wLjAuMTo4MjAwL3YxL3BraS9jcmwwDQYJKoZIhvcN
AQELBQADggEBAK6WMAxehaBiZQ+Nv6TNqDraSX+uxwOH4fGRgH49mgZrANrhu6Wt
rV0Iki8Gsg/VH5O5MLolLwLbedXQQ5PeynoC6Kr7ruFwTXAYHcXTDG4PRuVoOcKF
yhGgfHTg9CuPK9req0u1r5haP0zWq0HfoWbpinMWm40dmx78GE2BswLzuA54pjwT
DXzHiM4mdTJco1leMR0QeCGogc4hyYU5KaUJRBRdLEh1v/Hy1Kn5tXaxUxrQrSpx
e313iVQKmCUHSS2QAyVBY7SQ1DxcZWwGzIBmrnZ4SLxaoMf0c3CVR3DUfavp31v7
xstLR68wnDtj7BgcrCI9ipkH/2xUCZXvIzQ=
-----END CERTIFICATE-----]
certificate         -----BEGIN CERTIFICATE-----
MIIDZjCCAk6gAwIBAgIUGROOtSpDOFJp5BUALSfXjcMsTJwwDQYJKoZIhvcNAQEL
BQAwLTErMCkGA1UEAxMiZXhhbXBsZS5jb20gSW50ZXJtZWRpYXRlIEF1dGhvcml0
eTAeFw0yMjAxMTMwNzQ3NTVaFw0yMjAyMTIwNzQ4MjVaMBsxGTAXBgNVBAMTEHRl
c3QuZXhhbXBsZS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDS
H7LL/cfhkO4jjvZO3b8sCxCgWmuBZseIRD1msR8lvkYH04g9CUQd9/upYjHiw45/
34Fn4Rv1eGor5TotN8/HnOCm0S5f10k+hpb46jfvSpaHoVK4iYM+dOou8My9XBy5
WdabVUcEaJZrnzFWQ8i0/oQLx/YbNtoapAeAMi8WP4+vP7Mq5ZvKBmExG2doxQyx
VG8f1LEBy6MwiGMZJQ9YPyY6F41JtqBA//7YU/9DVI8ejd6wUAooltGYVK2qCZrr
uzSJVtagSx3zdsWGXdQ6rG0TLB7y4vnXZge0yLCzHXQD2xFbfNNbGCeOHZDvN/2c
a96glhJwYt6hRnqRCVTfAgMBAAGjgY8wgYwwDgYDVR0PAQH/BAQDAgOoMB0GA1Ud
JQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQU+iz2QunNHevdRadS
BQUVrOMavJ8wHwYDVR0jBBgwFoAUhoogfLP8Kx2eTyaeWSe9sA13xM0wGwYDVR0R
BBQwEoIQdGVzdC5leGFtcGxlLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAGFlGno0W
nchaUatSh7Z5P89C2IguIJ23KLRwHaS6WjZHZaQxTNMbQq+KOKFATuuxn00O3XF1
gtr1SyR9CTR9q8gm5bNPygbb/hShdTwJZltLZsMXcqS0bw8ZZ4AxZTpwXHVOB0xc
MO082KRKARL7tLTR1DkoJIyzj1xpN3DiuUrRtWMa1gJdQeT8k/swp00KcFaWJsOI
GduolLes1yaI20Emz49PN8NSquxL06KNLPDXHyPRJRqnzhqAYbV/Hf8YITx3Crl5
nkGrZPp2s5n4J4G//+33zAMCnB26dMBlpY7XrmYv3ohEvNzVJJ3x8rTCyyPrtTBu
AaiyRTz1w8mrxA==
-----END CERTIFICATE-----
expiration          1644652105
issuing_ca          -----BEGIN CERTIFICATE-----
MIIDpjCCAo6gAwIBAgIUK3ks87DCjns2g7WAKLmo1zdM1CQwDQYJKoZIhvcNAQEL
BQAwFjEUMBIGA1UEAxMLZXhhbXBsZS5jb20wHhcNMjIwMTEzMDc0NzE1WhcNMjcw
MTEyMDc0NzQ1WjAtMSswKQYDVQQDEyJleGFtcGxlLmNvbSBJbnRlcm1lZGlhdGUg
QXV0aG9yaXR5MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsRfeyN+F
lAe+fLhWVxDfsb23oCprvb5hgAtLeVGlMeSDJ2UeXjwrxYZzR8H9/bhIPfi+QMz5
bsFcpQL/MgLkyGNcv37+hTk/CFDzNRtnXQSIULe7h3pcZ6zg1NsBDhWPXiHdAy6H
4aL6yhRyRr4DHe0Xea429Xq2VD7nZw2WNKnrTCCaOIigezZqYfGTSTvuHdvDSvLD
jlN4JfJUv07wzk2D462QOtiWDPY+kxs0v8zo/5JA4vPocihuUIajLfTdJeLFzz/9
q4d7K1gSMe0zfPrFhLA/ahUMpwqp3krB4xGk6YnKOuvVUDAqNFknGadYVwU0cvne
EQdmg/mVDM/L0wIDAQABo4HUMIHRMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E
BTADAQH/MB0GA1UdDgQWBBSGiiB8s/wrHZ5PJp5ZJ72wDXfEzTAfBgNVHSMEGDAW
gBQdLCzsiHUdOri9k5nJxZaevsXbdzA7BggrBgEFBQcBAQQvMC0wKwYIKwYBBQUH
MAKGH2h0dHA6Ly8xMjcuMC4wLjE6ODIwMC92MS9wa2kvY2EwMQYDVR0fBCowKDAm
oCSgIoYgaHR0cDovLzEyNy4wLjAuMTo4MjAwL3YxL3BraS9jcmwwDQYJKoZIhvcN
AQELBQADggEBAK6WMAxehaBiZQ+Nv6TNqDraSX+uxwOH4fGRgH49mgZrANrhu6Wt
rV0Iki8Gsg/VH5O5MLolLwLbedXQQ5PeynoC6Kr7ruFwTXAYHcXTDG4PRuVoOcKF
yhGgfHTg9CuPK9req0u1r5haP0zWq0HfoWbpinMWm40dmx78GE2BswLzuA54pjwT
DXzHiM4mdTJco1leMR0QeCGogc4hyYU5KaUJRBRdLEh1v/Hy1Kn5tXaxUxrQrSpx
e313iVQKmCUHSS2QAyVBY7SQ1DxcZWwGzIBmrnZ4SLxaoMf0c3CVR3DUfavp31v7
xstLR68wnDtj7BgcrCI9ipkH/2xUCZXvIzQ=
-----END CERTIFICATE-----
private_key         -----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA0h+yy/3H4ZDuI472Tt2/LAsQoFprgWbHiEQ9ZrEfJb5GB9OI
PQlEHff7qWIx4sOOf9+BZ+Eb9XhqK+U6LTfPx5zgptEuX9dJPoaW+Oo370qWh6FS
uImDPnTqLvDMvVwcuVnWm1VHBGiWa58xVkPItP6EC8f2GzbaGqQHgDIvFj+Prz+z
KuWbygZhMRtnaMUMsVRvH9SxAcujMIhjGSUPWD8mOheNSbagQP/+2FP/Q1SPHo3e
sFAKKJbRmFStqgma67s0iVbWoEsd83bFhl3UOqxtEywe8uL512YHtMiwsx10A9sR
W3zTWxgnjh2Q7zf9nGveoJYScGLeoUZ6kQlU3wIDAQABAoIBAGo8qpqmuhuaujxy
OKhRRynOgl6AuOAZBgModV7paJUdqaylT7mEaNM6IRxX0F8VfoO6jmPmFcu7UPrW
j95y3HPRJmhRVMASSi75v3UkfplWvdrfPsKcjCXU2N5Y0zteSguQl4e7ywc4lezd
9QTnip+wGzUVqaHCzu0vR7eSJ+p312Af2PhWdPsRwVExds2qpOiuwHAsr8Ioy2fU
jWkAWACFuLEQtxjDuwUZ4K8c/XXxqZpSem3oNRaxUTU8ipp7DokK1DtsgOlnD3vU
Lh0hLTLb9nVy1VnoDkJIDvggod1Rd7IMcyZgmYtgtMP81+RSy7YrBPPYkAKv/AnO
paGltkECgYEA2AekBPFWKUbDL/92xCsXPypmZTIJBH9XBVfWlcEQP6u+gFJXYwwv
QVgtDK8rT/Dri+csZ50OP+3vyliSRqtt/0Q188bVOohL2eihTG7F5UFxBAydpZBW
W/ybBmCF6EK8NaXhceowe1XNkNd11yJRK6CpkcQAZbLhdltMHtl87f8CgYEA+QBQ
5Heqhj2lUw6nZVbml40JNcg1MhfBcHOZvdlmhWPBxr4G+uBe0ZxKlZCqj8rxLbWT
ei2XA+GNriAvbrR05KAZzE3meDz9QgbVnM6St0baiCh4plwePfz8VWTb/S8/U5Ya
Zw9Nz/dt1V+cFWvfpH5CqpIsAAQ8ng4w8FjPWSECgYEAoWWb+gFgcPgoLbz7u3XC
KXQBCkvGhvCoUBqe/EVAYYFrky1xklNxHq3FAnwArPn+0QhGmaayFbsrco6Xwmqj
hJougNGlTtSzbrspfxQnj69Dw1W1lhNvIcxo+eu1P6BUQvSKqXPhAtRI/5SpurAt
2p1u8rNv5IsvfSCaj1XHy3sCgYEAhoORL5sl19c9lJz5+Vj0wTJDo3ZAposGyQTq
LRFgvPajHAZUJvtGvd28vQel3IA5wgOxY/N0/Xe/3i0s8pUyAMAsr531v0bTWfPv
OgKuZ6wzKhMS+mwROlOMzWTrIt9/SlxwbvRpiMuV3gsEet4HtwkuYo8MjgW76Xap
IW4YtYECgYBFA4+48XjJz5ewgRxsAERb4V/K8XL894YdIUtVRyY+kihxmUQwE+f7
q1GIuLF7CjMwfUNYr5BAgW3ReWOcoCnSeoXqpCB1D85+fcvaTvhgLwPXjYoUWPOS
YAEBgb3YZ/9G3NZ9sgk2Xpm4XvqZylvpcRkzcSktQinVm2tyGstJ2g==
-----END RSA PRIVATE KEY-----
private_key_type    rsa
serial_number       19:13:8e:b5:2a:43:38:52:69:e4:15:00:2d:27:d7:8d:c3:2c:4c:9c
tim@tim:~$ ls
CA_cert.crt  intermediate.cert.pem  pki_intermediate.csr  private.key  test.example.com.crt  test.example.com.crt2


```

5. Установите корневой сертификат созданного центра сертификации в доверенные в хостовой системе.
![Screenshot](kr2.jpg)

6. Установите nginx.
```shell
tim@tim:/home/keys$ sudo systemctl status nginx
● nginx.service - A high performance web server and a reverse proxy server
     Loaded: loaded (/lib/systemd/system/nginx.service; enabled; vendor preset: enabled)
     Active: active (running) since Tue 2022-01-11 13:16:08 UTC; 31s ago
       Docs: man:nginx(8)
    Process: 13374 ExecStartPre=/usr/sbin/nginx -t -q -g daemon on; master_process on; (code=exited, status=0/SUCCESS)
    Process: 13386 ExecStart=/usr/sbin/nginx -g daemon on; master_process on; (code=exited, status=0/SUCCESS)
   Main PID: 13387 (nginx)
      Tasks: 3 (limit: 4578)
     Memory: 3.8M
     CGroup: /system.slice/nginx.service
             ├─13387 nginx: master process /usr/sbin/nginx -g daemon on; master_process on;
             ├─13388 nginx: worker process
             └─13389 nginx: worker process

Jan 11 13:16:08 tim systemd[1]: Starting A high performance web server and a reverse proxy server...
Jan 11 13:16:08 tim systemd[1]: Started A high performance web server and a reverse proxy server.

```
7. По инструкции ([ссылка](https://nginx.org/en/docs/http/configuring_https_servers.html)) настройте nginx на https, используя ранее подготовленный сертификат:
  - можно использовать стандартную стартовую страницу nginx для демонстрации работы сервера;
  - можно использовать и другой html файл, сделанный вами;
```shell
tim@tim:/home/keys$ cat /etc/nginx/nginx.conf
user www-data;
worker_processes auto;
pid /run/nginx.pid;
include /etc/nginx/modules-enabled/*.conf;

events {
        worker_connections 768;
        # multi_accept on;
}

http {
    server {
    listen              443 ssl;
    server_name         test.example.com;
    ssl_certificate     /home/keys/test.example.com.crt;
    ssl_certificate_key /home/keys/private.key;
    ssl_protocols       TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers         HIGH:!aNULL:!MD5;

```
8. Откройте в браузере на хосте https адрес страницы, которую обслуживает сервер nginx.
![Screenshot](kr3.jpg)
9. Создайте скрипт, который будет генерировать новый сертификат в vault:
  - генерируем новый сертификат так, чтобы не переписывать конфиг nginx;
  - перезапускаем nginx для применения нового сертификата.
```shell
#!/usr/bin/env bash

export VAULT_ADDR=http://127.0.0.1:8200
export VAULT_TOKEN=root

vault write -format=json pki_int/issue/example-dot-com common_name="test.example.com" ttl="740h" > /home/keys/test.example.com.crt.json

cat /home/keys/test.example.com.crt.json | jq -r '.data.certificate' >  /etc/nginx/ssl/test.example.com.crt
cat /home/keys/test.example.com.crt.json | jq -r '.data.ca_chain[0]' >> /etc/nginx/ssl/test.example.com.crt
cat /home/keys/test.example.com.crt.json | jq -r '.data.private_key' >  /etc/nginx/ssl/test.example.com.key
systemctl reload nginx
```
10. Поместите скрипт в crontab, чтобы сертификат обновлялся какого-то числа каждого месяца в удобное для вас время.
```shell
tim@tim:/home/keys$ cat /etc/crontab
# /etc/crontab: system-wide crontab
# Unlike any other crontab you don't have to run the `crontab'
# command to install the new version when you edit this file
# and files in /etc/cron.d. These files also have username fields,
# that none of the other crontabs do.

SHELL=/bin/sh
PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin

# Example of job definition:
# .---------------- minute (0 - 59)
# |  .------------- hour (0 - 23)
# |  |  .---------- day of month (1 - 31)
# |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
# |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
# |  |  |  |  |
# *  *  *  *  * user-name command to be executed
17 *    * * *   root    cd / && run-parts --report /etc/cron.hourly
25 6    * * *   root    test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.daily )
47 6    * * 7   root    test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.weekly )
52 6    1 * *   root    test -x /usr/sbin/anacron || ( cd / && run-parts --report /etc/cron.monthly )
#
59 23   30 * *  root    /home/keys/cert.sh

```
## Результат

Результатом курсовой работы должны быть снимки экрана или текст:

- Процесс установки и настройки ufw
- Процесс установки и выпуска сертификата с помощью hashicorp vault
- Процесс установки и настройки сервера nginx
- Страница сервера nginx в браузере хоста не содержит предупреждений 
- Скрипт генерации нового сертификата работает (сертификат сервера ngnix должен быть "зеленым")
- Crontab работает (выберите число и время так, чтобы показать что crontab запускается и делает что надо)

## Как сдавать курсовую работу

Курсовую работу выполните в файле readme.md в github репозитории. В личном кабинете отправьте на проверку ссылку на .md-файл в вашем репозитории.

Также вы можете выполнить задание в [Google Docs](https://docs.google.com/document/u/0/?tgif=d) и отправить в личном кабинете на проверку ссылку на ваш документ.
Если необходимо прикрепить дополнительные ссылки, просто добавьте их в свой Google Docs.

Перед тем как выслать ссылку, убедитесь, что ее содержимое не является приватным (открыто на комментирование всем, у кого есть ссылка), иначе преподаватель не сможет проверить работу. 
Ссылка на инструкцию [Как предоставить доступ к файлам и папкам на Google Диске](https://support.google.com/docs/answer/2494822?hl=ru&co=GENIE.Platform%3DDesktop).
