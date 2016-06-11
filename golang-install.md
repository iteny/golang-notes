golang安装(centos7.0)
=====
* 官方下载点([点击进入](https://golang.org/dl/ "官方下载点")) //选择对应版本复制连接
* golang中国下载点([点击进入](http://www.golangtc.com/download "golang中国下载点"))
```centos
wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz 		//下载
```
```centos
tar -zxvf go1.6.2.linux-amd64.tar.gz 		//解压文件
```
```centos
vim /etc/profile		//配置环境变量
```
```centos
export GOROOT=/root/go
export PATH=$PATH:$GOROOT/bin
//具体看你自己go安装的位置
```
```centos
source /etc/profile		//使得环境变量立即生效
```
```centos
go version		//如果环境变量配置成功，那么会显示当前的版本号
```
