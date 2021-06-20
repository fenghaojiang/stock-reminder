# stock-reminder

#### 运行方法

```shell script
docker build -t stock_reminder .
docker run -itd --rm --name stock-reminder stock_reminder
# -itd中的d是后台运行的意思，不想后台运行改成-it
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


