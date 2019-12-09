# go-zabbix-supervisord
Easy way to find and track supervisor processes by Zabbix.
## Requirements
* Zabbix server >= 4.4.3
* Zabbix agent >= 4.4.3
* go >= 1.13.4
* supervisord >= 3.3.4
* superuser rights on the host
## Usage
```bash
~# go-zabbix-supervisord -h
Usage of go-zabbix-supervisord:
  -sock string
    	The full path to the socket
    	 (default "/tmp/supervisor.sock")

~# go-zabbix-supervisord [-sock=/var/run/supervisor.sock] discovery
{
    "data":[
        {"{#PROCESS.NAME}":"001","{#PROCESS.GROUP}":"XXX","{#PROCESS.PID}":52390},
        {"{#PROCESS.NAME}":"321","{#PROCESS.GROUP}":"XX1","{#PROCESS.PID}":52398},
        {"{#PROCESS.NAME}":"123","{#PROCESS.GROUP}":"XX2","{#PROCESS.PID}":52411},
    ]
}

~# go-zabbix-supervisord [-sock=/var/run/supervisor.sock] state.name XX2:123
RUNNING
```
## Installation
1. Before you should check go version on your host. Just type command:  
    ```bash
    ~# go version
    go version go1.13.5 linux/amd64
    ```  
    
    > NOTE: If golang doesnt installed on the host, you can go to official golang  
    site and install it.

2. You need to check two environment variables that they are set correctly.  
    ```bash
    ~# echo $GOROOT; echo $GOBIN
    /usr/local/go
    /usr/local/go/bin
    ```
   
 3. **supervisord** must be running. Type command:
    ```bash
    ~# service supervisor status
    ● supervisor.service - Supervisor process control system for UNIX
       Loaded: loaded (/lib/systemd/system/supervisor.service; enabled; vendor preset: enabled)
       Active: active (running) since Wed 2019-12-04 14:05:31 MSK; 5 days ago
       ...
    ```
 4. Now you can install **go-zabbix-supervisord**.
    ```bash
    ~# go get -u github.com/cardinalit/go-zabbix-supervisord
    ```
    After installing the dependencies, the executable will be installed in `$GOBIN` path.  
    The executable will be called **go-zabbix-supervisord**. You can check it:
    ```bash
    ~# go-zabbix-supervisord help [-h | --help]
    Usage of go-zabbix-supervisord:
      -sock string
        	The full path to the socket
        	 (default "/tmp/supervisor.sock")
    ```
 5. You should create (or copy from my repository) `sudoers.d` exception file for  
 _zabbix_ user, containing the following text:
     ```text
    zabbix ALL=(ALL) NOPASSWD:/usr/local/go/bin/go-zabbix-supervisord
    Defaults:zabbix !requiretty
    ``` 
    For copy, just type: 
    ```bash
    ~# cp ~/go/src/github.com/cardinalit/go-zabbix-supervisord/zabbix \
       /etc/sudoers.d/
    ```
 6. Import `zbx_export_template.xml` to you Zabbix server.
 7. Create (or copy from my repository) file `userparameter_supervisord.conf`  
 to Zabbix agent path on the host. For copy, just type: 
     ```bash
     ~# cp ~/go/src/github.com/cardinalit/go-zabbix-supervisord/userparameter_supervisord.conf \
        /etc/zabbix/zabbix_agentd.d/
     ```
 8. Restart your Zabbix agent on the host. 