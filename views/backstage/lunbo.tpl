{{define "lunbo"}}
<style>
    #lunbo-panel{
        width: auto;
        height: 100%;
        padding: 10px;
    }
    .lb-cont a{
        text-align: right;
    }
    .imgclick{
        color: #01AAED;
    }
    .imgclick:hover{
        cursor: pointer;
    }
    .lunb-img{
        height: 28px;
    }
</style>
<div id="lunbo-panel">
    <div id="lunbo-title">
        <div class="layui-upload">
            <button type="button" class="layui-btn layui-btn-normal" id="selectfile">选择文件</button>
            <button type="button" class="layui-btn" id="startupload">开始上传</button>
          </div>
    </div>
    <div id="lunbo-body">
        <table class="layui-hide" id="lunboimg" lay-filter="lunboimg"></table>
    </div>
</div>
<script type="text/html" id="operbar">
    <a class="layui-btn layui-btn-danger layui-btn-xs layui-btn-normal" lay-event="del">删除</a>
</script>
<script>
    layui.use(['upload','table','layer'],function(){

        var table = layui.table;
        var upload = layui.upload;
        var layer=layui.layer;
        table.render({
          elem:'#lunboimg'
          ,url:'/backstage/getlunbojson'
          ,cellMinWidth: 200
          ,toolbar:'#toolsbar'
          ,cols:[[
            {checkbox: true, fixed: false}
            ,{field:'id',title:'ID',sort:true,fixed:false}
            ,{field:'name',title:'图片名称'}
            ,{field:'webdir',title:'图片',templet:function(d){
              return '<div class="lb-cont"><img class="lunb-img" src="'+d.webdir+'"/>&nbsp;&nbsp;<a class="imgclick" target="_blank" href="'+d.webdir+'">查看</a></div>';
            }}
            ,{field:'md5',title:'MD5值'}
            ,{fixed: 'right', title:'操作', toolbar: '#operbar'}
          ]]
          ,id:'luboimgs'
        });

       
        //选完文件后不自动上传
        upload.render({
          elem: '#selectfile'
          ,url: '/backstage/uploadlunboimg'
          ,auto: false
          ,accept:'images'
          ,acceptMime: 'image/jpg,image/png,image/JPG'
          ,size:300
          //,multiple: true
          ,bindAction: '#startupload'
          ,done: function(res){
            console.log(res)  
            if(res.code==1){
                //刷新数据
                table.reload('luboimgs', 'data');
            }else{
                layer.open({
                    title: '错误',
                     content: res.msg
                });
            }
          }
        });
        //table操作
        table.on('tool(lunboimg)',function(obj){
            if(obj.event === 'del'){
                layer.confirm('真的删除“'+obj.data.name+'”么?', function(index){
                    var urlstr='/backstage/dellunboimg?id='+obj.data.id;
                    $.ajax({
                        url:urlstr,
                        type:'get',
                        timeout:5000,
                        success:function(data){
                            if(data.code==1){
                                obj.del();
                            }else{
                                layer.open({
                                    title: '提示',
                                    content: res.msg
                                });
                            }
                        },
                        error:function(){
                            layer.open({
                              title: '服务器无响应',
                              content: "服务器跑到月球啦！！"
                            });
                        }
                    });

                    layer.close(index);
                });
            }
        });
    });
</script>
{{end}}