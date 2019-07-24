# Docker notes

## 列出镜像

`docker image ls`列出所有的下载镜像。

### 镜像体积

`docker system df`便捷查看镜像、容器、数据卷所占用的空间。

### 虚悬镜像dangling image

新旧镜像同名，旧镜像名称被取消产生。`docker image ls -f dangling=true`用来查看虚悬镜像。虚悬镜像没有存在价值`docker image prune`删除。

### 中间层镜像

由于重复利用资源，docker使用中间层镜像。`docker image ls`只显示顶层镜像。显示所有镜像需要 `docker image ls -a`。中间层镜像不可删除。删除全部上层镜像之后，会自动删除。

### 列出部分镜像

根据仓库名列出镜像：`docker image ls ubuntu`

指定仓库名和标签：`docker image ls ubunut:18.04`

使用过滤器参数：`--filter` 或者 `-f` 例如：

​	`docker image ls -f since=mongo:3.2` 另外，还可以使用 `before`

​	还可以利用镜像构建时定义的 `LABEL` 过滤 `docker image ls -f label=com.example.version=0.1`

### 以特定格式显示

有需求时，例如将ID交给 `docker image rm` 命令作为参数删除指定镜像，需要使用 `docker image ls -q`

`--filter` 配合 `-q` 产生指定范围的ID列表，送给另一个 `docker` 命令作为参数。命令搭配常见且重要。

需要特殊格式时，可以使用**Go的模板语法**。例如：`{{.ID}}: {{.Repository}}`

还可以以表格等距显示，且有标题行，与默认相同。`docker image ls --format "table {{.ID}}\t{{.Repository}}\t{{.Tag}}"` 

## 删除本地镜像

### 使用ID、镜像名、摘要删除

脚本可以使用 `长ID` 。人工输入可以使用 `短ID` 。`docker image ls` 默认列出短ID，输入时取前三位能够区分即可。

也可以使用镜像名，即 `<仓库名>:<标签>` 进行删除。`docker image rm centos`

最精确的是用 `镜像摘要` 删除。 `docker image ls --digest` 可以查看镜像摘要。`docker image rm <digest>` 进行删除。

### Untagged和Deleted

首先untag所有指向这个镜像的标签，然后再进行delete。层数不一定与pull时相同，因为docker是多层结构。

同时，要先删除镜像上层的容器层（还有容器存储层），否则无法进行删除。

### 用docker image ls命令配合

```bash
docker image rm $(docker image ls -q redis)
```

## 理解镜像构成

`docker commit` 可以用于学习和入侵后的现场保护，但定制镜像应该使用`Dockerfile`

每次执行 `docker run` 要指定镜像作为容器运行的基础。如果Docker Hub不能满足需求，就要定制镜像。

镜像是多层存储，每一层在前一层基础上修改。容器也是多层存储，以镜像为基础层，在这个基础上加一层作为容器运行时的存储层。

```bash
docker run --name webserver -d -p 80:80 nginx
```

可以用浏览器直接访问nginx服务器

```bash
docker exec -it webserver bash

root@...:/# echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index.html
root@...:/# exit
exit
```

以交互式终端方式进入了 `webserver` 并执行了 `bash` 命令，获得了一个可操作的Shell。

`docker diff webserver` 可以看到具体的改动如何。

运行一个容器并且不适用卷，任何文件修改都会被记录在容器存储层中。`docker commit`命令可以将存储层保存下来成为镜像。

语法：

```bash
docker commmit [选项] <容器ID或容器名> [<仓库名>[:<标签>]]
```



```bash
docker commit \
	--author "abc <abc@gmail.com>" \ 
	--message "修改网页" \
	webserver \
	nginx:v2
	sha256:...
```

注：powerShell中使用 **`** 进行换行操作

`docker history nginx:v2` 可以具体查看镜像内的历史记录

定制完成之后，可以运行该镜像。

```bash
docker run --name web2 -d -p 81:80 nginx:v2 #映射到81端口
```

### 慎用docker commit

由于命令执行，大量无关文件被改动或添加。如果不认真清理，镜像将异常臃肿。

`docker commit` 对镜像操作都是黑箱操作，生成的镜像也是 **黑箱镜像** ，维护异常痛苦。

### Docker 容器的终止

利用 `docker ps -a` 可以查看终止状态的容器。

`docker stop $CONTAINER_ID` 和 `docker restart $CONTAINER_ID`

## Dockerfile定制镜像

Dockerfile用于定制镜像，其中每一条**指令(instruction)**都用于构建一层。

Dockerfile内容：

```dockerfile
FROM nginx
RUN echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index.html
```

### FROM指定基础镜像

定制镜像是以一个镜像为基础，在其上进行定制。`FROM`是必备指令，且必须是第一条指令。

**Docker Hub**上有很多优质镜像：

| 服务类镜像       | nginx  | redis   | mongo  | httpd  | php    |
| ---------------- | ------ | ------- | ------ | ------ | ------ |
| 各种语言镜像     | node   | openjdk | python | ruby   | golang |
| 基础操作系统镜像 | ubuntu | debian  | centos | fedora | alpine |

特殊镜像 `scratch` 表示空白镜像。Linux下静态编译的程序不需要有操作系统运行时提供支持，一切库都在可执行文件内，因此直接 `FROM scratch` 会让镜像体积更小巧。**Golang** 就具有这样的优势。

### RUN执行命令

两种格式：

- *shell*： `RUN <command>`就像在命令行中输入命令一样。
- *exec* ： `RUN [".exe", "para1", "para2"]`类似函数调用格式。

每一个`RUN`都会建立一层，过多`RUN`会使镜像非常臃肿。

这是一个良好的示例：

```dockerfile
FROM debian:stretch

RUN buildDeps='gcc libc6-dev make wget' \
	&& apt-get update \
	&& apt-get install -y $buildDeps \
	&& wget -O redis.tar.gz "http://download.redis.io.releases/redis-5.0.3.tar.gz" \
	&& mkdir -p /usr/src/redis \
	&& tar -xzf redis.tar.gz -C /usr/src/redis --strip-components=1 \
	&& make -C /usr/src/redis \
	&& make -C /usr/src/redis install \
	&& rm -rf /var/lib/apt/lists/* \
	&& rm redis.tar.gz \
	&& rm -r /usr/src/redis \
	&& apt-get purge -y --auto-remove $buildDeps
	
# this is a comment
```

### 构建镜像

在 `Dockerfile` 目录执行：

```bash
docker build -t nginx:v3 . #.表示当前目录
```

### 镜像构建上下文(Context)

`.`并不是在指定 `Dockerfile` 的目录。这是在指定**上下文路径**。Docker实际上是在远程Docker引擎中执行各种命令。`docker build`命令构建镜像并非在本地构建，而是上传到引擎。构建时，并不是所有指令都通过 `RUN` ，还有 `COPY` 和 `ADD`。让服务器端获得本地文件就需要引入上下文概念。

构建时，用户指定构建镜像上下文路径。`docker build`获得路径之后，将路径下所有内容打包。上传Docker引擎。Docker引擎收到上下文包后，展开就可以获得所有需要文件。

```dockerfile
COPY ./package.json /app/
```

这并不是复制执行`docker build`命令所在目录下的`package.json`，也不是复制dockerfile目录下的，二十复制**上下文(Context)**目录下的 `package.json`

诸如：`COPY ../package.json /app`和 `COPY /opt/xxxx /app`无法工作就是因为超出了上下文的范围，Docker引擎无法获知。

`.`实际上时在指定上下文目录。

注意！不能因为避免 `COPY` 不工作就将Dockerfile置于硬盘根目录，这会打包上传整个硬盘！

### 其它docker build用法

#### 直接用Git repo进行构建

```bash
docker build https://github.com/.../xxx.git#11.1 这是指定默认目录/11.1/
```

#### 用给定的tar压缩包构建

```bash
docker build http://server/context.tar.gz
```

#### 从标准输入中读取Dockerfile构建

```
docker build - < Dockerfile

cat Dockerfile | docker build -
```

没有上下文，不能将本地文件 `COPY` 进镜像。

#### 从标准输入中读取上下文压缩包构建

```bash
docker build - < context.tar.gz
```

