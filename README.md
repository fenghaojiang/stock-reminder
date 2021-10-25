# 炒股提醒小助手  
# stock-reminder


[![Go Report Card](https://goreportcard.com/report/github.com/fenghaojiang/stock-reminder)](https://goreportcard.com/report/github.com/fenghaojiang/stock-reminder)

#### **投资有风险，代码仅供参考学习**  

股票信息请求的是  https://xueqiu.com  
代码执行后会在每天9点到15点之间每分钟去请求对应的股票信息，当低于配置文件中的预设值时会向配置文件中的邮箱发送信息邮件  



<br>

#### 运行方法(quickstart)  

修改conf文件夹下的config.toml文件为自己的配置(modify conf/config.toml with your config)   

<br>

build
```shell script
docker build -t stock_reminder .
```
run
```shell script 
docker run -itd --rm -v /home/opc/stock-reminder/conf/config.toml:/var/data/conf/config.toml  --name stock-reminder stock_reminder
```
-v 将配置文件config.toml挂载到容器内, 本人config.toml配置文件是在/home/opc/stock-reminder/conf/config.toml中，对应修改main.go中的config文件路径即可(modify config directory in main.go)  



#### 关闭
#### stop stock-reminder

```shell script
docker stop stock-reminder
```

<br>  



#### 删除所有none镜像
#### Remove All none images

```shell script
docker rmi $(docker images | awk '$1=="<none>"{print $3}')
# or
docker rmi $(docker images -f 'dangling=true' -q)
```


