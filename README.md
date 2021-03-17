# SNIProxyGo

Simple SNI proxy written in go.  
Source from <https://github.com/TachibanaSuzume/SNIProxyGo>.

## Compile

```sh
go get gopkg.in/yaml.v2
go build
```

## Usage:

* Open port 443
* Edit rules in config.yaml
* Start

## Usage:
Usage of ./SNIProxyGo:  
```bash
-D          Enable debug  
-F <string> log to file  
-c <string> config file (default "config.yaml")  
```

## Credit

| Author | Link | License |
| ------ | ---- | ------- |
| TachibanaSuzume | <https://github.com/TachibanaSuzume/SNIProxyGo> | GPL-3.0 |
| fangdingjun | <https://github.com/fangdingjun/sniproxy> | GPL-3.0 |
| ziozzang | <https://github.com/ziozzang/SimpleSNIProxy>  | BSD-2-Clause |
