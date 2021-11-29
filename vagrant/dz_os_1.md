# Домашнее задание к занятию "3.3. Операционные системы, лекция 1"

1. Какой системный вызов делает команда `cd`? В прошлом ДЗ мы выяснили, что `cd` не является самостоятельной  программой, это `shell builtin`, поэтому запустить `strace` непосредственно на `cd` не получится. Тем не менее, вы можете запустить `strace` на `/bin/bash -c 'cd /tmp'`. В этом случае вы увидите полный список системных вызовов, которые делает сам `bash` при старте. Вам нужно найти тот единственный, который относится именно к `cd`. Обратите внимание, что `strace` выдаёт результат своей работы в поток stderr, а не в stdout.

Ответ:  
```
vagrant@vagrant:/tim$ strace /bin/bash -c 'cd /tmp'
chdir("/tmp")
```



2.Попробуйте использовать команду `file` на объекты разных типов на файловой системе. Например:
    ```bash
    vagrant@netology1:~$ file /dev/tty
    /dev/tty: character special (5/0)
    vagrant@netology1:~$ file /dev/sda
    /dev/sda: block special (8/0)
    vagrant@netology1:~$ file /bin/bash
    /bin/bash: ELF 64-bit LSB shared object, x86-64
    ```
    Используя `strace` выясните, где находится база данных `file` на основании которой она делает свои догадки.  

Ответ:

``openat(AT_FDCWD, "/usr/share/misc/magic.mgc", O_RDONLY) = 3
fstat(3, {st_mode=S_IFREG|0644, st_size=5811536, ...}) = 0
mmap(NULL, 5811536, PROT_READ|PROT_WRITE, MAP_PRIVATE, 3, 0) = 0x7f4832775000      
close(3)                                = 0``


3. Предположим, приложение пишет лог в текстовый файл. Этот файл оказался удален (deleted в lsof), однако возможности сигналом сказать приложению переоткрыть файлы или просто перезапустить приложение – нет. Так как приложение продолжает писать в удаленный файл, место на диске постепенно заканчивается. Основываясь на знаниях о перенаправлении потоков предложите способ обнуления открытого удаленного файла (чтобы освободить место на файловой системе).

Ответ:
``ping 192.168.138.128 > log &`` команда пинг пишет в файл и переводим в фон  
``ps aux | grep ping  ``  
``vagrant    26439  0.0  0.0   7092   872 pts/0    S    12:57   0:00 ping 192.168.138.128``
``sudo lsof –p 26439``  
находим ``ping    26439 vagrant    1w   REG  253,0     6453 1966088 /home/vagrant/tim/log (deleted)``  
обнуляем ``echo ‘’ | sudo tee /proc/26439/fd/1`` 

4. Занимают ли зомби-процессы какие-то ресурсы в ОС (CPU, RAM, IO)?

Ответ:
Зомби не занимают памяти (как процессы-сироты), но блокируют записи в таблице процессов, размер которой ограничен для каждого пользователя и системы в целом.

При достижении лимита записей все процессы пользователя, от имени которого выполняется создающий зомби родительский процесс, не будут способны создавать новые дочерние процессы. Кроме этого, пользователь, от имени которого выполняется родительский процесс, не сможет зайти на консоль (локальную или удалённую) или выполнить какие-либо команды на уже открытой консоли (потому что для этого командный интерпретатор sh должен создать новый процесс), и для восстановления работоспособности (завершения виновной программы) будет необходимо вмешательство системного администратора.  

Зомби процессы можно посмотреть в htop столбец S 
Возможные значения состояния:

R — [running or runnable] запущенные или находятся в очереди на запуск  
S — [interruptible sleep] прерываемый сон  
D — [uninterruptible sleep] непрерываемый сон (в основном IO)  
Z — [zombie] процесс зомби, прекращенный, но не забранный родителем  
T — Остановленный сигналом управления заданиями  
t — Остановленный отладчиком  
X — Мёртвый (не должен показываться)  


5. В iovisor BCC есть утилита `opensnoop`:
    ```bash
    root@vagrant:~# dpkg -L bpfcc-tools | grep sbin/opensnoop
    /usr/sbin/opensnoop-bpfcc
    ```
    На какие файлы вы увидели вызовы группы `open` за первую секунду работы утилиты? Воспользуйтесь пакетом `bpfcc-tools` для Ubuntu 20.04. Дополнительные [сведения по установке](https://github.com/iovisor/bcc/blob/master/INSTALL.md).  

Ответ:

```
PID    COMM               FD ERR PATH  
599    multipathd          8   0 /sys/devices/pci0000:00/0000:00:10.0/host2/target2:0:0/2:0:0:0/state  

394    systemd-journal    31   0 /proc/599/status  
394    systemd-journal    31   0 /proc/599/status  
394    systemd-journal    31   0 /proc/599/comm  
394    systemd-journal    31   0 /proc/599/cmdline  
394    systemd-journal    31   0 /proc/599/status  
394    systemd-journal    31   0 /proc/599/attr/current  
394    systemd-journal    31   0 /proc/599/sessionid  
394    systemd-journal    31   0 /proc/599/loginuid  
394    systemd-journal    31   0 /proc/599/cgroup

```

6. Какой системный вызов использует `uname -a`? Приведите цитату из man по этому системному вызову, где описывается альтернативное местоположение в `/proc`, где можно узнать версию ядра и релиз ОС.

Ответ:
```
vagrant@vagrant:~$ strace uname -a
write(1, "Linux vagrant 5.4.0-89-generic #"..., 107Linux vagrant 5.4.0-89-generic #100-Ubuntu SMP Fri Sep 24 14:50:10 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux  


```
Part of the utsname information is also accessible via /proc/sys/kernel/{ostype, hostname, osrelease, version, domainname}.


7.Чем отличается последовательность команд через `;` и через `&&` в bash? Например:
    ```bash
    root@netology1:~# test -d /tmp/some_dir; echo Hi
    Hi
    root@netology1:~# test -d /tmp/some_dir && echo Hi
    root@netology1:~#  
    ```

Ответ:  
&&логический оператор ;это простая последовательность.

В cmd1 && cmd2, cmd2 будет запущен, только если cmd1 завершается с успешным кодом возврата.  

Принимая во внимание cmd1; cmd2, что cmd2 будет работать независимо от состояния выхода cmd1 (при условии, что вы не настроили свою оболочку на выход при всех сбоях в вашем скрипте или чем-то еще).


    Есть ли смысл использовать в bash `&&`, если применить `set -e`?

Опция -e  Exit immediately if a command exits with a non-zero status.
Эта опция будет завершать команду, если статус выполнения команды не 0, использовать нет смысла &&

8. Из каких опций состоит режим bash `set -euxo pipefail` и почему его хорошо было бы использовать в сценариях?

Ответ:  
-e Немедленный выход, если команда завершается с ненулевым статусом.  
-u При подстановке обрабатывать неустановленные переменные как ошибку.  
-x Печатать команды и их аргументы по мере их выполнения.
-o имя-параметра установить переменную, pipefail

Причина использования pipefail заключается в том, что иначе команда, неожиданно завершившаяся с ошибкой и находящаяся где-нибудь в середине конвейера, обычно остаётся незамеченной. Но, если использовалась опция set -e, не приведёт к аварийному завершению скрипта. 


9.Используя `-o stat` для `ps`, определите, какой наиболее часто встречающийся статус у процессов в системе. В `man ps` ознакомьтесь (`/PROCESS STATE CODES`) что значат дополнительные к основной заглавной буквы статуса процессов. Его можно не учитывать при расчете (считать S, Ss или Ssl равнозначными).

```
vagrant@vagrant:~$ ps -o stat
STAT
Ss
R+
```


 ```PROCESS STATE CODES
       Here are the different values that the s, stat and state output specifiers (header "STAT" or "S") will display to describe the state of a
       process:

               D    uninterruptible sleep (usually IO)
               I    Idle kernel thread
               R    running or runnable (on run queue)
               S    interruptible sleep (waiting for an event to complete)
               T    stopped by job control signal
               t    stopped by debugger during the tracing
               W    paging (not valid since the 2.6.xx kernel)
               X    dead (should never be seen)
               Z    defunct ("zombie") process, terminated but not reaped by its parent

       For BSD formats and when the stat keyword is used, additional characters may be displayed:

               <    high-priority (not nice to other users)
               N    low-priority (nice to other users)
               L    has pages locked into memory (for real-time and custom IO)
               s    is a session leader
               l    is multi-threaded (using CLONE_THREAD, like NPTL pthreads do)
               +    is in the foreground process group
```
   
Наиболее часто встречаемые статусы S и R с заглавной буквы.
Маленькая буква s означает s    is a session leader.
Считать S, Ss или Ssl равнозначными