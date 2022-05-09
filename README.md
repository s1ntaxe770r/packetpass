# Packetpass 

An ebpf program that allows packets only on a specifc port 


usage 

```bash
pip install bcc
```

start a small webserver 
```bash
export URL=https://raw.githubusercontent.com/s1ntaxe770r/pong-server/master/main.go
curl -o pong $URL\ 
chmod +x pong && ./pong -port ":4040"
```

## Without packetpass  

![without_pass](https://file.coffee/u/CPt8b7pEPXmY5X.png)


### With packetpass 

![with_pass](https://file.coffee/u/eUYSDxjbsWvZuC.png)


