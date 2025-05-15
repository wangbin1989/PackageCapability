# PackageCapability

## 初次安装

```sh
sudo /var/packages/PackageCapability/target/first-install.sh
```

## 系统要求

- go 1.24+

## 准备环境

### 安装 SpkSrc

```sh
apt-get install git
mkdir -p /toolkit
cd /toolkit
git clone https://github.com/SynoCommunity/spksrc.git
```

## 构建

```sh
git clone https://github.com/wangbin1989/PackageCapability.git
cd PackageCapability/
make setup
make
```
