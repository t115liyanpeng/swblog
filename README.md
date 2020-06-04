# swblog使用说明文档
### 一个简短的说明
- 第一次使用golang写博客系统
- 采用的是Gin Web 框架
- 数据库使用的mysql
- 因为本人不是专业的web前端开发 所以使用了一个layui的web框架，就目前看起来还挺容易上手的
- 由于很多时候就是想起来做什么就开始写代码，导致最后代码比较混乱 T^T
### 如何配置
####	1、数据库
- 安装mysql,这里我用的mysql的版本是5.7，其他的版本应该也没有问题
- 新建数据库swblogdb然后找到源码目录的db文件夹中的swblogdb.sql文件进行还原（我使用navicat备份下来的，包含部分测试数据）
####	&nbsp;&nbsp;&nbsp;1、swblog配置文件
- 配置文件放在conf目录下的conf.json内容如下
#
	{
		"servercfg": {
			"webname": "swblog", //网站名称
			"port": "8080", //服务监听的端口			
			"indexpagesize": 2, //首页的每页显示文章的条数
			"artfilepagesize": 10, //文章归档显示的每页条数
			"userid": "c999a2f041c84dc1b5970bb973c1da74" //用户id 数据库中 t_userb 表中的id值
		},
		"databasecfg": {
			"address": "192.168.99.100", //数据库IP地址
			"port": 3306, //数据库连接端口
			"user": "swblog", //数据库用户名
			"password": "swblog", //数据库密码
			"dbname": "swblogdb" //数据库名称
		}
	}
###	开始使用
####	&nbsp;&nbsp;&nbsp;1、源代码
- 通过git将源代码下载下来之后进行go build 发现并不能直接编译成功，是因为还有一些第三方的库需要下载，
比如说mysql的驱动库，需要使用go get 去将依赖的包下载下来
- 下载完成之后 直接go build 成功就可以直接使用了
####	&nbsp;&nbsp;&nbsp;2、目录使用说明
- conf目录下存放的是网站的配置文件（网站运行时必须有）
- static目录下存放的是图片js、css文件（网站运行时必须有）
- views目录下存放的是静态、动态的html文件（网站运行时必须有）
####	&nbsp;&nbsp;&nbsp;3、运行程序
- 执行编译好的程序后，通过浏览器访问http:localhost:8080
![index.png](https://i.loli.net/2020/06/04/BCQGoInWYUrh1uK.png)
- 进入管理页需要在浏览器上输入http://localhost:8080/user/login
![login.png](https://i.loli.net/2020/06/04/khOPBqGamJtEKdf.png)
- 登录后的后台显示如下，默认用户名swblog密码123456
![back.png](https://i.loli.net/2020/06/04/HFXWhqtnTd8uYJ4.png)
- 文章管理中可以对文章进行增删改，添加文章支持markdown格式，上传的图片 需要用其他的图床，这块我没有做图片的保存
![artlist.png](https://i.loli.net/2020/06/04/uwSabGvpLRECqTz.png)
![artadd.png](https://i.loli.net/2020/06/04/Hh6U5zVykduaIjY.png)

###	文档就写到这里吧，如果喜欢欢迎star，或者留言给我