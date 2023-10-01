# Flaring

> 源项目：[soulteary](https://github.com/soulteary)/[flare](https://github.com/soulteary/flare) ([AGPL-3.0 license](https://github.com/soulteary/flare/blob/main/LICENSE))

这是一个使用 GO 语言开发的书签网址导航应用程序，主要是在上游项目 [Flare](https://github.com/soulteary/flare) 的基础上修改并添加新特性。

## 新功能

在基础功能上基本和 [Flare](https://github.com/soulteary/flare) 保持一致，并额外添加以下功能：

1. 添加控制程序[日志输出级别](#日志输出级别)。
1. 生产环境下不再输出 `gin-gonic/gin` 包日志。
1. 首页的搜索框支持实时搜索书签。
1. 添加一个可选的返回顶部按钮。
1. 添加支持使用 [Simple Icons](https://simpleicons.org/)(v9.16.1) 图标，格式为 `si` 前缀 + [slug](https://github.com/simple-icons/simple-icons/blob/master/slugs.md)，如：`siGitHub`。
1. 添加可选的侧边栏功能。(_v0.4.0-20230316_)
1. 调整为可选水平(默认)或垂直排列书签。(_v0.4.0-20230326_)
1. 调整基础样式，以优化在移动端下的使用体验。
1. 书签网址支持[动态 URL 参数](#动态-url-参数)。 (_v0.4.1-2023-08-18_)

### 日志输出级别

- 名称:
  - 环境变量: `FLARE_LOG_LEVEL`
  - 启动命令: `log_level`
- 可选值: `TRACE`、`DEBUG`、`INFO`、`WARN`、`ERROR`、`PANIC`

#### 示例

配置程序日志输出级别为 `ERROR`:

- 通过环境变量来配置: 添加 `FLARE_LOG_LEVEL` 环境变量并将值设为 `ERROR`。
- 通过启动命令来配置: `flare --log_level=ERROR`。

### 动态 URL 参数

> Since v0.4.1-2023-08-18

#### 参数

假设 flare 服务的首页地址为 `https://192.168.0.1:5005/`，以下可用的各参数及其对应解析结果：

| 参数名 | 解析结果 |
| --- | --- |
| `host` | `192.168.0.1:5005` |
| `hostname` | `192.168.0.1` |
| `href` | `https://192.168.0.1:5005/` |
| `origin` | `https://192.168.0.1:5005` |
| `pathname` | `/` |
| `port` | `5005` |
| `protocol` | `https:` |

#### 示例

假设某书签网址配置为 `https://{hostname}:8888/test` 时:

- 当 flare 服务的首页地址为 `https://192.168.0.1:5005/`，其显示为 `https://192.168.0.1:8888/test`。
- 当 flare 服务的首页地址为 `https://172.17.0.1:5005/`，其显示为 `https://172.17.0.1:8888/test`。
- 当 flare 服务的首页地址为 `https://link.example.com/`，其显示为 `https://link.example.com:8888/test`。

## 其他改动

其他改动主要包含修复在 [Flare](https://github.com/soulteary/flare) 正式发行版本中存在的问题，这些问题理论上会在 [Flare](https://github.com/soulteary/flare) 的后续迭代版本中被处理和修复，所以这些更改基本只会针对特定的发行版本：

- **0.4.1:**
  - 修复加密链接可能无法解码的问题。
  - 修复无法读取 .env 文件中所配置值的问题 (_v0.4.1-20230628_)
  - 修复验证用户名或密码不正确时提示有误的问题 (_v0.4.1-20230628_)
  - 修正应用编辑下表格标题的显示 (_v0.4.1-20230628_)
  - 修复在移动端下应用编辑顶部内容溢出导致显示异常的问题 (_v0.4.1-20230628_)
  - 在线数据编辑支持拖动行来进行排序 (_v0.4.1-20230628_)
  - 修复同域下多个项目中登录状态会相互影响的问题 (_v0.4.1-20231001_)
  - 更新 [Material Design Icons](https://materialdesignicons.com/) 图标至 v7.2.96 版本。

---

<details>
<summary>旧版本中的改动</summary>

- **0.4.0:**
  - 修复应用程序在 Windows 环境下生成图标路径不正确导致图标无法显示的问题。
  - 修复界面设置中保存大小写设置的值显示异常的问题。
  - 修复在没有分类时书签显示异常的问题。
  - 修复子页面下设置按钮显示异常的问题。(_v0.4.0-20230314_)
  - 修复子页面下的按钮无法通过设置隐藏的问题。(_v0.4.0-20230314_)
  - 更新 [Material Design Icons](https://materialdesignicons.com/) 图标至 v7.2.96 版本。

</details>

## 程序截图

**桌面端默认：**

![Desktop](https://gcore.jsdelivr.net/gh/LightAPIs/PicGoImg@master/img/202303162130685.jpg)

**移动端默认：**

![Mobile](https://gcore.jsdelivr.net/gh/LightAPIs/PicGoImg@master/img/202303162131742.jpg)

## Docker 部署

Docker Hub 镜像：[giterhub/flare](https://hub.docker.com/r/giterhub/flare)，

**快速部署：**

```sh
# pull
docker pull giterhub/flare:latest

# run
docker run -d \
    --name flare \
    -p 5005:5005 \
    --mount type=bind,source=$PWD/flare/app,target=/app \
    -e FLARE_LOG_LEVEL=ERROR \
    giterhub/flare:latest
```

其他环境变量及使用方法可以参考：[docker-flare](https://github.com/soulteary/docker-flare)。
