# stock-reminder

#### 运行方法

```shell script
docker build -t stock_reminder .
docker run -itd --rm -v /home/opc/stock-reminder/conf/config.toml:/var/data/conf/config.toml  --name stock-reminder stock_reminder
# -v 将配置文件config.toml挂载到容器内, 本人config.toml配置文件是在/home/opc/stock-reminder/conf/config.toml中，对应修改即可
```

#### 关闭

```shell script
docker stop stock-reminder
```

#### 删除所有none镜像

```shell script
docker rmi $(docker images | awk '$1=="<none>"{print $3}')
# or
docker rmi $(docker images -f 'dangling=true' -q)
```


