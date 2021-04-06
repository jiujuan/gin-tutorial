## Air 介绍

[Air](https://github.com/cosmtrek/air) 能实时监听项目的代码文件，在代码发生变化之后会重新编译执行，这样就可以提高开发的效率。

- 地址：https://github.com/cosmtrek/air



具体的使用方法可以看 [README](https://github.com/cosmtrek/air/blob/master/README.md)

## 安装

有好几种安装方法：

1. 经典安装

   ```sh
   get -u github.com/cosmtrek/air
   ```

   

2. ### macOS, Linux, Windows

   ```sh
   # binary will be $(go env GOPATH)/bin/air
   curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
   
   # or install it into ./bin/
   curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
   
   air -v
   ```



3. docker

```shell
docker run -it --rm \
    -w "<PROJECT>" \
    -e "air_wd=<PROJECT>" \
    -v $(pwd):<PROJECT> \
    -p <PORT>:<APP SERVER PORT> \
    cosmtrek/air
    -c <CONF>
```

例子：

```shell
docker run -it --rm \
    -w "/go/src/github.com/cosmtrek/hub" \
    -v $(pwd):/go/src/github.com/cosmtrek/hub \
    -p 9090:9090 \
    cosmtrek/air
```

## 使用

1. 进入你的项目目录： `cd /path/to/your_project`

2. 然后使用命令：`air -c .air.conf`

   .air.conf 是一个配置文件，这个配置文件的示例可以在这里找到：[air_example.toml](https://github.com/cosmtrek/air/blob/master/air_example.toml)，它是一个 toml 格式文件。



具体操作：

1.新建文件 .air.conf

2.把 [air_example.toml](https://github.com/cosmtrek/air/blob/master/air_example.toml) 中的内容复制到 .air.conf 文件里，然后根据你的需求修改配置内容

3.如果你的文件名跟我一样是 .air.conf ，那么直接执行 air 命令就可以了，如果不是，那么运行命令：air -c your_file_name



> 说明：我是 windwos 下测试，更改文件后，每次都会弹出一个终端 cmd 出来，也是麻烦。