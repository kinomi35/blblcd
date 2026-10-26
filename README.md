    

---

<p align="center">
<img src="./example/blblcd.png" alt="blblcd" style="width:3em;margin: 0px auto" /></p>

blblcd（Bilibili Comment Dowloader），一款基于 [bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect) 的B站视频评论下载工具。

> ⚠️ 本仓库为二次开发 Fork 版本，在上游原版基础上新增两个功能：**关键词搜索站内视频批量爬取评论、搜索结果过滤（视频时长 / 上传时间筛选）**。

🗂️ 上游原版仓库： [Codeberg](https://codeberg.org/wsh233/blblcd) | [GitHub](https://github.com/WShihan/blblcd)

🎉 特点

- 使用方便，仅包含一个可执行文件，无需安装其他依赖。
- 下载单/多个视频评论，可按热评/时间顺序。
- 下载最新一期视频
- 下载指定UP主视频的评论，可按投稿时间/收藏/播放顺序。
- ✅ **新增：关键词搜索站内视频，批量爬取搜索结果视频评论**
- ✅ **新增：搜索结果过滤，支持按视频时长、上传时间范围筛选目标视频**
- 下载子评论，可下载隐藏评论和“楼中楼”评论。
- 下载评论中的图片，可将图片下载到本地同时将链接写入csv。
- 地图统计，支持评论统计输出为地图展示。
- 可自定义请求延时、重试次数、协程并发数
- 跨平台，支持Mac，Linux， Windows。

🔔 请留意

- 这是一个命令行程序，没有图形化界面。
- 若爬取结果数量不对只有十几条评论和网页上看到的数量不一致，那么很大概率是cookie失效导致。
- 爬取结果的csv文件编码为utf-8，使用诸如Microsoft Office之类的办公软件打开会显示乱码，建议用记事本或者代码编辑器打开。
- 上游原版暂不接受新增功能开发请求，本fork版本新增的搜索、筛选功能相关问题可在本仓库提交issue。
- 如果你发现已有功能出现bug，欢迎提交issue，我将尽力去解决。

⚠️ 请注意

> 最新上游原版代码托管在[Codeberg](https://codeberg.org/wsh233/blblcd)，上游GitHub仓库仅作为备份，同步更新可能存在延迟。由于GitHub访问原因，0.1.8之后的版本，其编译后的二进制产物将只发布在 Codeberg [Releases](https://codeberg.org/wsh233/blblcd/releases) 。

爬取评论信息如下：

```text
    Uname         名称
    Sex           性别
    Content       评论内容
    Rpid          评论id
    Oid           评论区id
    Bvid          视频bv
    Mid           发送者id
    Parent        父级评论
    Fansgrade     是否粉丝标签
    Ctime         评论时间戳
    Like          喜欢数
    Following     是否关注
    Current_level 当前等级
    Location      位置
