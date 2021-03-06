{{define "wuserconfig"}}
<div id="baseconfig">
    <form class="layui-form">
        <div id="svrinfo" class="conf">
          <!--<fieldset>
            <legend>用户头像</legend>
            <div class="layui-field-box">
              
            </div>
          </fieldset>-->
         <fieldset class="layui-elem-field">
             <legend>管理用户信息</legend>
             <div class="layui-form-item">
                <label class="layui-form-label">头像</label>
                <div id="tx" class="layui-input-block">
                  <img id="tximg" src="{{.TxURL}}" /> 
                </div>
                 
             </div>
             <span id="modetx" class="layui-input-block">修改头像</span>
             

             <div class="layui-field-box">
                <div class="layui-form-item">
                  <label class="layui-form-label">用户名</label>
                  <div class="layui-input-block">
                    <input type="text" name="name" value="{{.Name}}" required  lay-verify="required" placeholder="请输入用户名" autocomplete="off" class="layui-input"/>
                  </div>
                </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">登录名</label>
                 <div class="layui-input-block">
                   <input type="text" name="loginname" value="{{.LoginName}}" required  lay-verify="required" placeholder="请输入登录名" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">登录密码</label>
                 <div class="layui-input-block">
                   <input type="password" name="password" value="{{.PassWord}}" required  lay-verify="required" placeholder="请输入密码" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">个性签名</label>
                 <div class="layui-input-block">
                   <input type="text" name="brief" value="{{.Brief}}" required  lay-verify="required" placeholder="请输入个性签名" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
               <div class="layui-form-item">
                 <label class="layui-form-label">Email</label>
                 <div class="layui-input-block">
                   <input type="text" name="email" value="{{.Email}}" required  lay-verify="required" placeholder="请输入Email" autocomplete="off" class="layui-input"/>
                 </div>
               </div>
           </fieldset>
        </div>
        <div id="configsave">
            <button id="usersave-btn" type="submit" lay-submit lay-filter="saveuser" class="layui-btn layui-btn-normal">保存</button>
        </div>
    </form>
</div>
<script>
    layui.config({
        base: '/static/cropper/' //layui自定义layui组件目录
    }).use(['form','croppers'],function(){
                var form=layui.form;
                form.on('submit(saveuser)',function(data){
                     var password=hex_md5(data.field.password);
                    var jsondata={
                                    "id":{{.ID}},
                                    "name":data.field.name,
                                    "loginname":data.field.loginname,
                                    "password":password,
                                    "brief":data.field.brief,
                                    "email":data.field.email
                     }
                    var loading = layer.msg('加载中...', {icon: 16, shade: 0.3, time:0});
                    $.ajax({
                        type:'post',
                        url:'/backstage/setuser',
                        async:true,
                        timeout:5000,
                        dataType:'json',
                        contentType:'application/json',
                        data:JSON.stringify(jsondata),
                        success:function(data){
                            layer.close(loading);
                            if(data.code=="1"){
                                 layer.open({
                                    title: '提示',
                                    content: "修改成功！",
                                  });
                                console.log(jsondata.name);
                                $('#author').text(jsondata.name);

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
                var modurl="/backstage/modusertx?userid="+{{.ID}};
                var $ = layui.jquery
                  ,cropper = layui.croppers
                  ,layer= layui.layer;
                cropper.render({
                      elem: '#modetx'
                      ,saveW:150     //保存宽度
                      ,saveH:150
                      ,mark:1/1    //选取比例
                      ,area:'900px'  //弹窗宽度
                      ,url:modurl //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
                      ,done: function(url){ //上传完毕回调
                          //$("#inputimgurl").val(url);
                          //console.log(url);
                          var txturl="/static/user/user_tx.PNG?time="+Math.random();
                          $("#tximg").attr('src',txturl);
                          $("#navtx").attr('src',txturl);
                          //alert(url);
                      }
                  });
            });
</script>
{{end}}