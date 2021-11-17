1. https://cloud.mail.ru/public/fPCj/gVtQstDa3

2. ```PS E:\vagrant> vagrant version``` 
```Installed Version: 2.2.19

3. Windows terminal installed

4. PS E:\vagrant> vagrant up  
Bringing machine 'default' up with 'vmware_desktop' provider...  
==> default: Verifying vmnet devices are healthy...  
==> default: Preparing network adapters...   
==> default: Starting the VMware VM...  
==> default: Waiting for the VM to receive an address...  
==> default: Forwarding ports...  
    default: -- 22 => 2222  
==> default: Waiting for machine to boot. This may take a few minutes...  
    default: SSH address: 127.0.0.1:2222  
    default: SSH username: vagrant  
    default: Warning: Connection reset. Retrying...  
    default: Warning: Connection aborted. Retrying...  
    default:  
    default: Vagrant insecure key detected. Vagrant will automatically replace  
    default: this with a newly generated keypair for better security.  
    default:  
    default: Inserting generated public key within guest...  
    default: Removing insecure key from the guest if it's present...  
==> default: Machine booted and ready!  
==> default: Configuring network adapters within the VM...  
==> default: Waiting for HGFS to become available...  
==> default: Enabling and configuring shared folders...  
    default: -- E:/vagrant: /vagrant  
PS E:\vagrant> vagrant status  
Current machine states:  
default                   running (vmware_desktop)  

5.  System information as of Fri 12 Nov 2021 03:06:07 PM UTC

  System load:  0.01               Processes:             148
  Usage of /:   11.1% of 30.88GB   Users logged in:       0
  Memory usage: 22%                IPv4 address for eth0: 192.168.138.128
  Swap usage:   0%

6. cat .\Vagrantfile
vi .\Vagrantfile
config.vm.provider "virtualbox" do |v|
  v.memory = 2048
  v.cpus = 2
end
vagrant reload

7.PS E:\vagrant> vagrant ssh  
Welcome to Ubuntu 20.04.3 LTS (GNU/Linux 5.4.0-89-generic x86_64)  

 * Documentation:  https://help.ubuntu.com  
 * Management:     https://landscape.canonical.com  
 * Support:        https://ubuntu.com/advantage  

  System information as of Mon 15 Nov 2021 09:18:29 AM UTC  

  System load:  0.08               Processes:             146   
  Usage of /:   11.1% of 30.88GB   Users logged in:       0   
  Memory usage: 23%                IPv4 address for eth0: 192.168.138.128  
  Swap usage:   0%  


This system is built by the Bento project by Chef Software  
More information can be found at https://github.com/chef/bento  
Last login: Fri Nov 12 15:06:08 2021 from 192.168.138.2  
vagrant@vagrant:~$  

8. Раздел HISTORY параметр HISTSIZE строка 2784
Опция HISTCONTROL контролирует каким образом список команд сохраняется в истории.  
ignorespace — не сохранять строки начинающиеся с символа <пробел>  
ignoredups — не сохранять строки, совпадающие с последней выполненной командой  
ignoreboth — использовать обе опции ‘ignorespace’ и ‘ignoredups’  

9. { list; } строка 242  
 This is known as a group command.  The return status is the exit status of list.  
 { and  } are reserved words and must occur where a reserved word is permitted to be recognized.      
Since they do not cause a word break, they must be separated from list by whitespace or another shell metacharacter.  

10. vagrant@vagrant:~$ touch {1..100000}   
vagrant@vagrant:~$ touch {1..300000}  
-bash: /usr/bin/touch: Argument list too long  
11. /\[\[ поиск по двум [[.  \ экранирует символ  
конструкция [[ -d /tmp ]] возвращает 0 или 1 в зависимости от выражения внутри ( В данном случае возвращает Истину,    
т. к. выражение внутри проверяет существует директория  /tmp ) Сама конструкция является улучшеной конструкией тест [] 

12. vagrant@vagrant:~$ mkdir /tmp/new_path_directory/   
vagrant@vagrant:~$ cp /bin/bash /tmp/new_path_directory/  
vagrant@vagrant:~$ PATH=/tmp/new_path_directory/:$PATH  
vagrant@vagrant:~$ type -a bash
bash is /tmp/new_path_directory/bash   
bash is /usr/bin/bash  
bash is /bin/bash  

13. at и batch читают команды из стандартного ввода или заданного файла, 
которые будут выполнены в определённое время, используя /bin/sh.  

at
запускает команды в заданное время.  

batch
запускает команды, когда уровни загрузки системы позволяют это делать;   
в других, когда средняя загрузка системы, читаемая из /proc/loadavg, 
опускается ниже 0.5, или величины, заданной при вызове atrun.  

14. vagrant halt  или vagrant suspend
