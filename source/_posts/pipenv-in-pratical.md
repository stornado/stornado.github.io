---
title: pipenv 使用入门及规范
date: 2019-01-23 00:52:22
tags:
  - tutorial
  - pipenv
  - pip
  - virtualenv
categories:
  - tutorial
---

## 简介

[Pipenv](https://pipenv.readthedocs.io/en/latest/)，它的项目简介为 Python Development Workflow for Humans，是 Python 著名的 requests 库作者 kennethreitz 写的一个包管理工具，它可以为我们的项目自动创建和管理虚拟环境并非常方便地管理 Python 包。

[Pipenv](https://pipenv.readthedocs.io/en/latest/)我们可以简单理解为 pip 和 virtualenv 的集合体，它可以为我们的项目自动创建和管理一个虚拟环境。virtualenv 在使用时我们需要手动创建一个虚拟环境然后激活，Pipenv 会自动创建。

总的来说，[Pipenv](https://pipenv.readthedocs.io/en/latest/)可以解决如下问题：

- 我们不需要再手动创建虚拟环境，[Pipenv](https://pipenv.readthedocs.io/en/latest/)会自动为我们创建，它会在某个特定的位置创建一个 virtualenv 环境，然后调用 pipenv shell 命令切换到虚拟环境。
- 使用 requirements.txt 可能会导致一些问题，所以 [Pipenv](https://pipenv.readthedocs.io/en/latest/)使用 Pipfile 和 Pipfile.lock 来替代之，而且 Pipfile 如果不存在的话会自动创建，而且在安装、升级、移除依赖包的时候会自动更新 Pipfile 和 Pipfile.lock 文件。
- 广泛使用 Hash 校验，保证安全性。
- 可以更清晰地查看 Python 包及其关系，调用 pipenv graph 即可呈现，结果简单明了。
- 可通过自动加载 .env 读取环境变量，简化开发流程。

## 安装

`pip install pipenv`

## 基本命令

### 1. 创建环境

`pipenv` 或者指定具体 python 版本 `pipenv --python 3.7`

### 2. 安装依赖库

```shell
pipenv intall flask                     # 安装依赖
pipenv install 'flask==1.0.2'           # 指定版本安装
pipenv install 'pylint<2.0.0' --dev     # 只在开发环境安装
pipenv install -r requirements.txt      # 使用requirements.txt安装

pipenv install                          # clone其他库，本地已有Pipfile.lock文件，直接安装Pipfile.lock指定的依赖库
```

**卸载依赖**

```shell
pipenv uninstall
pipenv uninstall flask
```

### 3. 查看依赖

```shell
pipenv graph          # 查看目前安装的库及其依赖
pipenv run pip list
```

`pipenv graph` 展示效果如下

```shell
requests==2.19.1
  - certifi [required: >=2017.4.17, installed: 2018.8.24]
  - chardet [required: >=3.0.2,<3.1.0, installed: 3.0.4]
  - idna [required: >=2.5,<2.8, installed: 2.7]
  - urllib3 [required: >=1.21.1,<1.24, installed: 1.23]
```

### 4. 生成 Pipfile.lock 文件

```shell
pipenv lock
pipenv update   # 或者
pipenv sync     # 或者
```

### 5. 生成 requirements.txt 文件

`pipenv lock -r >requirements.txt`

### 6. 其他命令

```shell
pipenv shell             # 激活当前虚拟环境，当项目根目录下有 .env 文件时，自动解析添加环境变量
pipenv --where           # 查看当前工作目录
pipenv --venv            # 查看当前虚拟环境路径
pipenv --py              # 查看当前python解释器路径
pipenv --envs            # 查看可用环境变量
```

## 使用规范

### 1. 创建虚拟环境指定 python 版本

`pipenv --python 3.7`

### 2. 指定安装源加速安装

生成虚拟环境后修改 `Pipfile` 文件，指定阿里云或豆瓣为安装源

```toml
[[source]]
url = "https://pypi.doubanio.com/simple"
verify_ssl = true
name = "douban"
```

### 3. 安装/增加依赖后，必须生成 `Pipfile.lock` 文件

`pipenv update`

### 4. _(可选) 安装/变更依赖后，生成 requirements.txt 文件_

`pipenv lock -r > requirements.txt`
