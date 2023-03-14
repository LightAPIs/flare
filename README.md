# Flaring

> 源项目：[soulteary](https://github.com/soulteary)/[flare](https://github.com/soulteary/flare) ([AGPL-3.0 license](https://github.com/soulteary/flare/blob/main/LICENSE))

这是一个使用 GO 语言开发的书签网址导航应用程序，主要是在上游项目 [Flare](https://github.com/soulteary/flare) 的基础上修改并添加新特性。

## 新功能

在基础功能上基本和 [Flare](https://github.com/soulteary/flare) 保持一致，并额外添加以下功能：

1. 添加控制程序日志输出级别的环境变量(`FLARE_LOG_LEVEL`)/启动命令(`log_level`)，可选值：`TRACE`、`DEBUG`、`INFO`、`WARN`、`ERROR`、`PANIC`。示例：
    - 通过环境变量配置程序日志输出级别：添加 `FLARE_LOG_LEVEL` 环境变量并将值设为 `ERROR`。
    - 通过启动命令配置程序日志输出级别：`flare --log_level=ERROR`。
1. 生产环境下不再输出 `gin-gonic/gin` 包日志。
1. 首页的搜索框支持实时搜索书签。
1. 添加一个可选的返回顶部按钮。
1. 添加支持可选水平排列书签的功能。
1. 添加支持使用 [Simple Icons](https://simpleicons.org/) 图标，格式为 `si` 前缀 + [slug](https://github.com/simple-icons/simple-icons/blob/master/slugs.md)，如：`siGitHub`。(*注：在 [Flare](https://github.com/soulteary/flare) v0.4.0 后便会添加此功能。*)
1. 调整基础样式，以优化在移动端下的使用体验。

## 其他改动

其他改动主要包含修复在 [Flare](https://github.com/soulteary/flare) 正式发行版本中存在的问题，这些问题理论上会在 [Flare](https://github.com/soulteary/flare) 的后续迭代版本中被处理和修复，所以这些更改基本只会针对特定的发行版本：

- **0.4.0:**
    - 修复应用程序在 Windows 环境下生成图标路径不正确导致图标无法显示的问题。
    - 修复界面设置中保存大小写设置的值显示异常的问题。
    - 修复在没有分类时书签显示异常的问题。
    - 修复子页面下设置按钮显示异常的问题。(*v0.4.0-20230314*)
    - 修复子页面下的按钮无法通过设置隐藏的问题。(*v0.4.0-20230314*)
    - 更新 [Material Design Icons](https://materialdesignicons.com/) 图标至 v7.1.96 版本。

## 程序截图

**默认情况：**

![Flaring](https://gcore.jsdelivr.net/gh/LightAPIs/PicGoImg@master/img/202303121516709.jpg)

**水平排列书签：**

![HorizontalBookmarks](https://gcore.jsdelivr.net/gh/LightAPIs/PicGoImg@master/img/202303121518536.jpg)

## Docker 部署

Docker Hub 镜像：[giterhub/flare](https://hub.docker.com/r/giterhub/flare)，部署可以参考：[docker-flare](https://github.com/soulteary/docker-flare)。
