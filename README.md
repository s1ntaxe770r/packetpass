# Packetpass 

An ebpf program that allows packets only on a specifc port 


usage 

```bash
pip install bcc
```

start a small webserver 
```bash
curl -o main.go https://raw.githubusercontent.com/s1ntaxe770r/pong-server/master/main.go && go build -o pong main.go && chmod +x pong; ./pong -port ":4040"

```

## Without packetpass  

![without_pass](https://file.coffee/u/CPt8b7pEPXmY5X.png)


### With packetpass 

![with_pass](https://file.coffee/u/eUYSDxjbsWvZuC.png)


