---
title: pyenv in pratical
tags:
  - pyenv
  - virtualenv
  - python
categories:
  - tutorial
date: 2020-01-08 10:15:12
---




# pyenv 工程实践指南

## 安装

`brew install pyenv`

或者

`Python 3.7` 依赖 `OpenSSL 1.0.2e` 或更高版本，手动安装OpenSSL或者直接安装3.6最新版本

```shell
git clone https://github.com/pyenv/pyenv.git ~/.pyenv
echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bash_profile
echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bash_profile
echo -e 'if command -v pyenv 1>/dev/null 2>&1; then\n  eval "$(pyenv init -)"\nfi' >> ~/.bash_profile

git clone https://github.com/yyuu/pyenv-virtualenv.git ~/.pyenv/plugins/pyenv-virtualenv
echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.bash_profile

sudo yum install openssl-devel zlib-devel python-devel -y

pyenv install -v 3.6.10
# 若下载失败则手动下载上传到cache目录再安装
mkdir ~/.pyenv/cache
cd ~/.pyenv/cache
wget https://www.python.org/ftp/python/3.6.10/Python-3.6.10.tgz -P ~/.pyenv/cache
mv ~/.pyenv/cache/Python-3.6.10.tgz ~/.pyenv/cache/Python-3.6.10.tar.gz
cp ~/.pyenv/cache/Python-3.6.10.tar.gz /tmp

pyenv install -v 3.6.10
# 或
CFLAGS=-I/usr/include/openssl LDFLAGS=-L/usr/lib64 pyenv install -v 3.6.10

# CentOS7 下载
# https://www.python.org/ftp/python/3.7.6/Python-3.7.6.tar.xz
# https://www.python.org/ftp/python/3.6.10/Python-3.6.10.tar.xz
```

<!-- more -->

### 安装cx_Oracle

需要保证 `Python` 位数与 `Oracle intstantclient` 位数一致

```shell
sudo rpm -ivh oracle-instantclient12.1-basic-12.1.0.2.0-1.x86_64.rpm
# 或
sudo yum install oracle-instantclient12.1-basic-12.1.0.2.0-1.x86_64.rpm

sudo cp /usr/lib/oracle/12.1/client64/lib/libclntsh.so.12.1 /usr/lib/oracle/12.1/client64/lib/libclntsh.so
sudo sh -c "echo /usr/lib/oracle/18.5/client64/lib > /etc/ld.so.conf.d/oracle-instantclient.conf"
sudo ldconfig
sudo export LD_LIBRARY_PATH=/usr/lib/oracle/12.1/client64/lib:$LD_LIBRARY_PATH
sudo mkdir -p /usr/lib/oracle/12.1/client64/lib/network/admin
sudo export TNS_ADMIN=/usr/lib/oracle/12.1/client64/lib/network/admin

pyenv shell 3.6.10
# Or
pyenv virtualenv 3.6.10 venv-db-py3.6
pyenv shell venv-db-py3.6

pip install -i https://pypi.doubanio.com/simple cx_oracle
```

## 命令

```shell
Usage: pyenv <command> [<args>]

Some useful pyenv commands are:
   commands    List all available pyenv commands
   activate    Activate virtual environment
   commands    List all available pyenv commands
   deactivate   Deactivate virtual environment
   exec        Run an executable with the selected Python version
   global      Set or show the global Python version
   help        Display help for a command
   hooks       List hook scripts for a given pyenv command
   init        Configure the shell environment for pyenv
   install     Install a Python version using python-build
   local       Set or show the local application-specific Python version
   prefix      Display prefix for a Python version
   rehash      Rehash pyenv shims (run this after installing executables)
   root        Display the root directory where versions and shims are kept
   shell       Set or show the shell-specific Python version
   shims       List existing pyenv shims
   uninstall   Uninstall a specific Python version
   version     Show the current Python version and its origin
   --version   Display the version of pyenv
   version-file   Detect the file that sets the current pyenv version
   version-name   Show the current Python version
   version-origin   Explain how the current Python version is set
   versions    List all Python versions available to pyenv
   virtualenv   Create a Python virtualenv using the pyenv-virtualenv plugin
   virtualenv-delete   Uninstall a specific Python virtualenv
   virtualenv-init   Configure the shell environment for pyenv-virtualenv
   virtualenv-prefix   Display real_prefix for a Python virtualenv version
   virtualenvs   List all Python virtualenvs found in `$PYENV_ROOT/versions/*'.
   whence      List all Python versions that contain the given executable
   which       Display the full path to an executable

See `pyenv help <command>' for information on a specific command.
For full documentation, see: https://github.com/pyenv/pyenv#readme
```

## 结合pipenv使用

```shell
pip install -i https://pypi.doubanio.com/simple pipenv
pipenv --python 3.6.10
pipenv install --pypi-mirror https://pypi.doubanio.com/simple requests pymydql fire
```

### pipevn命令

```shell
Usage: pipenv [OPTIONS] COMMAND [ARGS]...

Options:
  --where             Output project home information.
  --venv              Output virtualenv information.
  --py                Output Python interpreter information.
  --envs              Output Environment Variable options.
  --rm                Remove the virtualenv.
  --bare              Minimal output.
  --completion        Output completion (to be eval'd).
  --man               Display manpage.
  --support           Output diagnostic information for use in GitHub issues.
  --site-packages     Enable site-packages for the virtualenv.  [env var:
                      PIPENV_SITE_PACKAGES]
  --python TEXT       Specify which version of Python virtualenv should use.
  --three / --two     Use Python 3/2 when creating virtualenv.
  --clear             Clears caches (pipenv, pip, and pip-tools).  [env var:
                      PIPENV_CLEAR]
  -v, --verbose       Verbose mode.
  --pypi-mirror TEXT  Specify a PyPI mirror.
  --version           Show the version and exit.
  -h, --help          Show this message and exit.


Usage Examples:
   Create a new project using Python 3.7, specifically:
   $ pipenv --python 3.7

   Remove project virtualenv (inferred from current directory):
   $ pipenv --rm

   Install all dependencies for a project (including dev):
   $ pipenv install --dev

   Create a lockfile containing pre-releases:
   $ pipenv lock --pre

   Show a graph of your installed dependencies:
   $ pipenv graph

   Check your installed dependencies for security vulnerabilities:
   $ pipenv check

   Install a local setup.py into your virtual environment/Pipfile:
   $ pipenv install -e .

   Use a lower-level pip command:
   $ pipenv run pip freeze

Commands:
  check      Checks for security vulnerabilities and against PEP 508 markers
             provided in Pipfile.
  clean      Uninstalls all packages not specified in Pipfile.lock.
  graph      Displays currently-installed dependency graph information.
  install    Installs provided packages and adds them to Pipfile, or (if no
             packages are given), installs all packages from Pipfile.
  lock       Generates Pipfile.lock.
  open       View a given module in your editor.
  run        Spawns a command installed into the virtualenv.
  shell      Spawns a shell within the virtualenv.
  sync       Installs all packages specified in Pipfile.lock.
  uninstall  Un-installs a provided package and removes it from Pipfile.
  update     Runs lock, then sync.
```

