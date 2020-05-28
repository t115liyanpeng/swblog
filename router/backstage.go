package router

import (
	"fmt"
	"net/http"
	"swblog/controllers"
	"swblog/models/artclassify"
	"swblog/models/conf"
	"swblog/models/page"
	"swblog/models/user"
	"swblog/tools"

	"github.com/gin-gonic/gin"
)

var backstage *gin.RouterGroup

//RegisterBackStageRoute 注册后台的路由
func RegisterBackStageRoute(eng *gin.Engine) {
	backstage = eng.Group("backstage")
	backstage.GET("/index", backstageIndex)
	backstage.GET("/wbconfig", backstageBaseConfig)
	backstage.POST("/setconfig", setconfig)
	backstage.GET("/wuconfig", wuconfig)
	backstage.POST("/setuser", setuserconfig)
	backstage.GET("/classify", getClassify)
	backstage.GET("/getclassjson", getclassjson)
	backstage.POST("/addclass", addclassify)
	backstage.GET("/delclassify", delclass)
	backstage.POST("/modclass", modclassify)
	backstage.GET("/tagcfg", tagcfg)
	backstage.GET("/gettagsjson", gettagsjson)
	backstage.POST("/addtag", addtag)
	backstage.POST("/modtag", modtag)
	backstage.GET("/lunbo", lunbo)
	backstage.GET("/getlunbojson", getlunbojson)
	backstage.POST("uploadlunboimg", uploadlunboimg)
}

//后台默认页
func backstageIndex(ctx *gin.Context) {
	data := controllers.GetBgIndex(tools.SvrCfg.Server.UserID)
	ctx.HTML(http.StatusOK, "backstageIndex", data)
}

//服务器基本配置
func backstageBaseConfig(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "wbconfig", tools.SvrCfg)
}

//设置配置文件
func setconfig(ctx *gin.Context) {
	cfg := conf.Config{
		Server:   &conf.Servercfg{},
		Database: &conf.Databasecfg{},
	}
	err := ctx.BindJSON(&cfg)
	if err == nil {
		if controllers.SetConfig(&cfg) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "1",
			})
			return
		}
		err = fmt.Errorf("set json faild")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  "0",
		"error": err.Error,
	})

}

//用户设置
func wuconfig(ctx *gin.Context) {
	data := controllers.GetDbUserInfo(tools.SvrCfg.Server.UserID)
	ctx.HTML(http.StatusOK, "wuserconfig", data)
}

//修改用户
func setuserconfig(ctx *gin.Context) {
	dbu := &user.DbUser{}
	err := ctx.BindJSON(dbu)
	if err == nil {
		ret, msg := dbu.UpdateDbUser()
		if ret {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  msg,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  msg,
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "read json faild",
		})
	}
}

//获取分类
func getClassify(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "classifycfg", nil)
}

//获取分类列表
func getclassjson(ctx *gin.Context) {
	jsonstr := controllers.GetClassJosn(tools.SvrCfg.Server.UserID)
	if jsonstr != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  jsonstr.Code,
			"msg":   jsonstr.Msg,
			"count": jsonstr.Count,
			"data":  jsonstr.Data,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"msg":   "没有数据",
			"count": 0,
			"data":  "",
		})
	}
}

//添加分类
func addclassify(ctx *gin.Context) {
	data := artclassify.Classify{}
	err := ctx.BindJSON(&data)
	if err == nil {
		if controllers.AddClassify(tools.SvrCfg.Server.UserID, &data) {

			ctx.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  "success",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "faild",
	})
}

//删除分类
func delclass(ctx *gin.Context) {
	id := ctx.Query("id")
	if id != "" {
		r, e := controllers.DelClassify(id)
		var code int = 0
		var msg string = ""
		if r {
			code = 1
		}
		if e != nil {
			msg = e.Error()
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ID不存在哦",
		})
	}

}

//修改分类
func modclassify(ctx *gin.Context) {
	var cla artclassify.Classify
	err := ctx.BindJSON(&cla)
	if err == nil {
		ret, err := controllers.UpdateClass(&cla)
		var code int
		var msg string
		if ret {
			code = 1
			msg = ""
		} else {
			code = 0
			msg = err.Error()
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error,
		})
	}
}

//标签管理
func tagcfg(ctx *gin.Context) {
	droplist := controllers.GetClassDropList(tools.SvrCfg.Server.UserID)
	ctx.HTML(http.StatusOK, "tagcfg", droplist)
}

//获取标签列表
func gettagsjson(ctx *gin.Context) {
	jsonstr := controllers.GetTagsJSON(tools.SvrCfg.Server.UserID)
	if jsonstr != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  jsonstr.Code,
			"msg":   jsonstr.Msg,
			"count": jsonstr.Count,
			"data":  jsonstr.Data,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"msg":   "没有数据",
			"count": 0,
			"data":  "",
		})
	}
}

//添加标签
func addtag(ctx *gin.Context) {
	var tag artclassify.Tags = artclassify.Tags{}
	err := ctx.BindJSON(&tag)
	if err == nil {
		if controllers.AddTag(tools.SvrCfg.Server.UserID, &tag) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "添加成功！",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "添加失败！",
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求参数不正确！",
		})
	}

}

//修改标签
func modtag(ctx *gin.Context) {
	var tag artclassify.Tags = artclassify.Tags{}
	err := ctx.BindJSON(&tag)
	if err == nil {
		if controllers.UpdateTag(&tag) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "修改成功！",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "修改失败！",
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求参数不正确！",
		})
	}
}

//轮播管理
func lunbo(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "lunbo", nil)
}

//获取轮播数据json
func getlunbojson(ctx *gin.Context) {
	list := page.GetLunBoPic(tools.SvrCfg.Server.UserID)
	code := -1
	msg := ""
	if list != nil && len(list) > 0 {
		code = 0
	} else if list != nil && len(list) == 0 {
		code = 1
		msg = "没有数据！"
	}
	data := page.TableJSONData{
		Code:  code,
		Msg:   msg,
		Count: 1,
		Data:  list,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  data.Code,
		"msg":   data.Msg,
		"count": data.Count,
		"data":  data.Data,
	})
}

//上传轮播图片
func uploadlunboimg(ctx *gin.Context) {
	code := -1
	msg := "上传失败！"
	file, err := ctx.FormFile("file")
	if err == nil {
		savepath := fmt.Sprintf("%s/static/user/%s", tools.AppPath, file.Filename)
		err = ctx.SaveUploadedFile(file, savepath)
		if err == nil {
			err = controllers.SaveLunBoPic(savepath, file.Filename)
			if err == nil {
				code = 0
				msg = "上传成功"
			}
		} else {
			msg = err.Error()
		}

	} else {
		msg = err.Error()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
