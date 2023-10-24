# Traffic_killer
流量杀手，快速消耗服务器流量
# 

# 自行编译
## 给服务器安装go环境
### 1. 下载go文件
```
wget https://golang.google.cn/dl/go1.21.3.linux-amd64.tar.gz
```
### 2. 解压go文件到指定位置
删除 /usr/local/go 文件夹（如果存在）来删除以前安装的任何 Go，然后将刚刚下载的存档解压到 /usr/local，在 /usr/local/go 中创建一个新的 Go 树：
```
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
```
（您可能需要以 root 身份或通过sudo运行该命令）。
### 3. 配置环境变量
```
export PATH=$PATH:/usr/local/go/bin
```
### 4. 测试go环境
```
go version
```
## 克隆仓库
