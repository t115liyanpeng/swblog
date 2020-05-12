{{define "wbconfig"}}
<div id="baseconfig">
    <form class="layui-form" method="POST">
        <div id="svrinfo" class="conf">
         <fieldset class="layui-elem-field">
             <legend>服务器基础信息</legend>
             <div class="layui-field-box">
             <div class="layui-form-item">
                 <label class="layui-form-label">Web名称</label>
                 <div class="layui-input-block">
                   <input type="text" name="webname" value="{{.Server.WebName}}" required  lay-verify="required" placeholder="请输入Web名称" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">Web端口</label>
                 <div class="layui-input-block">
                   <input type="text" name="webport" value="{{.Server.Port}}" required  lay-verify="number" placeholder="请输入端口号" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">首页条数</label>
                 <div class="layui-input-block">
                   <input type="text" name="indexpagesize" value="{{.Server.IndexPageSize}}" required  lay-verify="number" placeholder="请输入首页文章显示条数" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">归档条数</label>
                 <div class="layui-input-block">
                   <input type="text" name="artpagesize" value="{{.Server.FilePageSize}}" required  lay-verify="number" placeholder="请输入文章归档显示条数" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
             <
           </fieldset>
        </div>
        <div id="database" class="conf">
        <fieldset class="layui-elem-field">
         <legend>数据库配置信息</legend>
         <div class="layui-field-box">
             <div class="layui-form-item">
                 <label class="layui-form-label">数据库地址</label>
                 <div class="layui-input-block">
                   <input type="text" name="dbaddress" value="{{.Database.IPAddress}}" required  lay-verify="required" placeholder="请输入数据库地址" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
             <div class="layui-form-item">
                 <label class="layui-form-label">数据库名称</label>
                 <div class="layui-input-block">
                   <input type="text" name="dbname" value="{{.Database.DbName}}" required  lay-verify="required" placeholder="请输入数据库名称" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">数据库端口</label>
                 <div class="layui-input-block">
                   <input type="text" name="dbport" value="{{.Database.Port}}" required  lay-verify="number" placeholder="请输入端口号" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">用户名</label>
                 <div class="layui-input-block">
                   <input type="text" name="dbuser" value="{{.Database.User}}" required  lay-verify="required" placeholder="请输入数据库用户名" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">密码</label>
                 <div class="layui-input-block">
                   <input type="password" name="dbpassword" value="{{.Database.Password}}" required  lay-verify="required" placeholder="请输入数据库密码" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
         </div>
       </fieldset>
        </div>
        <div id="configsave">
            <button id="configsave-btn" type="submit" lay-submit lay-filter="savecfg" class="layui-btn layui-btn-normal">保存</button>
        </div>
    </form>
</div>
<script>
    layui.use('form',function(){
                var form=layui.form;
                form.on('submit(savecfg)',function(data){
                    
                    var jsondata={
                                     "servercfg":{
                                         "webname":data.field.webname,
                                         "port":data.field.webport,
                                         "indexpagesize":parseInt(data.field.indexpagesize),
                                         "artfilepagesize":parseInt(data.field.artpagesize),
                                     },
                                     "databasecfg":{
                                         "address":data.field.dbaddress,
                                         "port":parseInt(data.field.dbport),
                                         "user":data.field.dbuser,
                                         "password":data.field.dbpassword,
                                         "dbname":data.field.dbname
                                     }
                                 }
                    var loading = layer.msg('加载中...', {icon: 16, shade: 0.3, time:0});
                    //alert(JSON.stringify(jsondata));
                    $.ajax({
                        type:'post',
                        url:'/backstage/setconfig',
                        async:true,
                        timeout:5000,
                        dataType:'json',
                        contentType:'application/json',
                        data:JSON.stringify(jsondata),
                        success:function(data){
                            console.log(data);
                            layer.close(loading);
                            if(data.code=="1"){
                                 layer.open({
                                title: '提示',
                                content: "修改成功！是否重启服务器？",
                                yes:function(){
                                    alert('test');
                                }
                            });
                            }else{
                                layer.open({
                                title: '服务器无响应',
                                content: data.error
                            });
                            }
                        },
                        error:function(){
                            layer.close(loading);
                            layer.open({
                                title: '服务器无响应',
                                content: "服务器跑到月球啦！！"
                            });
                        }
                    });
                    return false;
                });
            });
</script>
{{end}}