# PESCMS RENT房租管理系统

PESCMS RENT(下称PR)是一款基于GPLv2协议发布的开源房租管理系统，程序基于Golang + VUE3编写。开源版包含基础的房租信息管理，数据库选用sqlite，基于Golang的语言特性，**编译后的程序，实现了开箱即用，无需安装任何环境**。

## PR专业版

如您需要更多使用功能，推荐使用我们的PR专业版或者网络版。

开源版、专业版和网络版差异说明：[https://pescms.com/download/9.html](https://pescms.com/download/9.html)

## 反馈和建议

邮箱：sale#pescms.com
软件下载地址：[https://document.pescms.com/article/13/696610900081115136.html](https://document.pescms.com/article/13/696610900081115136.html)  
官方问答中心：[https://forum.pescms.com/](https://forum.pescms.com/)  
软件文档：[https://document.pescms.com/article/13.html](https://document.pescms.com/article/13.html)  
PESCMS官方QQ 1群：451828934 <a target="_blank" href="http://shang.qq.com/wpa/qunwpa?idkey=70b9d382c5751b7b64117191a71d083fbab885f1fb7c009f0dc427851300be3a"><img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="PESCMS官方1群" title="PESCMS官方1群"></a>  
PESCMS官方QQ 2群：496804032 <a target="_blank" href="https://jq.qq.com/?_wv=1027&k=5HqmNLN"><img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="PESCMS官方2群" title="PESCMS官方2群"></a>

## 快速开始

### 运行后端

使用PR需要先安装golang，然后进入pescms-rent目录后，执行如何命令

```
go run main.go
```

执行此命令即可运行后端。

#### fresh 边修改边调试

若您需要边修改边调试，推荐使用fresh。安装方式请网上搜索。

在根目录下执行如何命令：

```
fresh -c fresh.conf
```

**上述任意方式运行后，浏览器访问http://127.0.0.1:8080/ 即可。**

### 运行前端

在pescms-rent目录下，有一个vue目录，此为项目的前端代码。在运行前端环境，您需要先安装node和npm，安装教程请自行网上搜索。

```
# 安装前端依赖
npm install

# 运行前端项目
npm run dev
```

**执行npm run命令后，浏览器访问http://localhost:5173/ 首次访问系统会引导您进入安装界面。**

## 使用协议

1. PR系统基于GPLv2开源协议发布，您二次开发需遵守此协议的约定。
2. 您在使用PR系统，您需要为您录入的数据负责，请根据您当地法律法规要求填写数据。PESCMS对此不负任何法律责任。

