**本项目已经不再更新, 欢迎感兴趣的人来接管这个项目 [#36](https://github.com/sabersalv/freedom-routes/issues/36)**

支持平台: Linux, Mac OS X, Windows, OpenWRT, DD-WRT, ASUSWRT, OpenWRT(ChinaDNS)

生成一个可以运行的脚本, 当VPN运行的时候, 自动添加国内的IP地址到系统`路由表`, 用`直接连接`方式访问国内的网站, 用`VPN`方式访问国外的网站, 从而提升网页浏览速度. (例如: 使用前ping baidu.com是300ms延迟, 使用后可以减少到30ms)

## 更新

- 2015-12-28: 更新了新的下载地址, 请大家注意. (老地址已停止更新并且不能保障永久有效).
- 2015-12-24: Windows下导入时间加快, 现在和Linux下一样, 都是秒时间导入. [#9](https://github.com/sabersalv/freedom-routes/issues/9)

## 项目一览
命令行: [Go版本 - 本项目](https://github.com/sabersalv/freedom-routes), [Python版本](https://github.com/fivesheep/chnroutes)

图形化应用程序: [Windows](https://github.com/cqjjjzr/freedom-routes#windows%E7%89%88%E5%9B%BE%E5%BD%A2%E5%8C%96%E5%BA%94%E7%94%A8%E7%A8%8B%E5%BA%8F)

## 对chnroutes的改进

1. Linux和Windows下导入路由的速度更快, 秒时间导入
2. 支持[模板](https://github.com/sabersalv/freedom-routes/tree/master/routes/templates), 可以自定义脚本

# 网络版本

### 下载

每24小时更新一次.

- **Linux, OpenWRT, DD-WRT, ASUSWRT**: [linux.tar.gz](https://raw.githubusercontent.com/sabersalv/freedom-routes/dist/linux.tar.gz)
- **Mac OS X**: [mac.zip](https://raw.githubusercontent.com/sabersalv/freedom-routes/dist/mac.zip)
- **Windows**: [windows.zip](https://raw.githubusercontent.com/sabersalv/freedom-routes/dist/windows.zip)
- **Android**: [android.zip](https://raw.githubusercontent.com/sabersalv/freedom-routes/dist/android.zip)
- **RouterOS**: [routeros.zip](https://raw.githubusercontent.com/sabersalv/freedom-routes/dist/routeros.zip)

### 使用方法

这些ip地址库并不是固定不变的, 尽管变化不大, 但还是建议每隔两三个月更新一次.

(手动运行)

```
启用
# ./routes-up.sh

停用
# ./routes-down.sh
```

Windows需要右键routes-up.bat -> 已管理员身份运行.

(OpenVPN)

```
# cp them to /etc/openvpn
# edit /etc/openvpn/hello.conf

  script-security 2
  up ./routes-up.sh
  down ./routes-down.sh
```

(PPTP)

```
cp routes-up.sh /etc/ppp/ip-up
cp routes-down.sh /etc/ppp/ip-down.d/ip-down
```

### 自动更新

(Linux)

```
# curl https://raw.githubusercontent.com/sabersalv/freedom-routes/master/misc/linux/freedom-routes.service -o /etc/systemd/system/freedom-routes.service
# curl https://raw.githubusercontent.com/sabersalv/freedom-routes/master/misc/linux/freedom-routes.timer -o /etc/systemd/system/freedom-routes.timer
# systemctl start freedom-routes.timer
# systemctl enable freedom-routes.timer
```

(Windows)

用计划任务

# 本地版本

查看[本地版本使用说明](https://github.com/sabersalv/freedom-routes/blob/master/docs/local.md)

# 开发

查看[开发文档](https://github.com/sabersalv/freedom-routes/blob/master/docs/Development.md)

## 任何人都可以帮助这个项目

- 请保持低调, 不要推广本项目, 自己用的好就可以了, 不要点击右上角的star按钮
- 提交Bug/建议
- 帮助作者提高文档
- 欢迎其他开发者Fork和发送Pull-Request
- 欢迎其他开发者申请Collaborator权限

感谢所有[贡献者](https://github.com/sabersalv/freedom-routes/contributors). </br>
感谢原来的[chnroutes](https://github.com/fivesheep/chnroutes)作者.
