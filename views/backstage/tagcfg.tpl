{{define "tagcfg"}}
<style>

#bkzz{
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: none;
    background-color: rgba(0, 0, 0, 0.6);
    z-index: 999;
  }
  #datainfo{
    position: relative;
    top: 25%;
    left: 35%;
    width: 500px;
    height: 300px;
    min-width: 500px;
    min-height: 300px;
    border: #d2d2d2 solid 0.5px;
    background-color: #F0F0F0;
    border-radius: 3px;
  }
  #addcontent{
    width: 460px;
    height: 240px;
    margin: auto;
    position: relative;
    top: 10px;
  }
  
  #operbtn{
    position: absolute;
    bottom: 3px;
    width: 500px;
    text-align: center;
  }
  .closebtn{
    color: red;
    font-size: 24px;
    position: relative;
    top: 3px;
    left: 470px;
  }
  .closebtn:hover{
    opacity: 0.7;
    cursor: pointer;
  }

</style>
<!--遮罩-->
<div id="bkzz">
  <div id="datainfo">
        <i class="closebtn layui-icon layui-icon-close-fill"></i>
        <div id="addcontent">
          <form class="layui-form" method="POST">
            <input id="classid" type="hidden" />
            <div class="layui-field-box">
                <div class="layui-form-item">
                  <label class="layui-form-label">所属分类</label>
                  <div class="layui-input-block">
                    <select id="classname" name="classname" lay-verify="required" lay-filter="classname">
                      <option value=""></option>
                      {{range.}}
                        <option value="{{.ID}}">{{.Name}}</option>
                      {{end}}
                    </select>
                  </div>
                </div>
                <div class="layui-form-item">
                    <label class="layui-form-label">标签名称</label>
                    <div class="layui-input-block">
                      <input type="text" name="name" id="name"  required  lay-verify="required" placeholder="标签名称" autocomplete="off" class="layui-input"/>
                    </div>
                </div>
                <div class="layui-form-item">
                  <label class="layui-form-label">标签描述</label>
                  <div class="layui-input-block">
                    <input type="text" name="brief" id="brief" required  lay-verify="required" placeholder="标签描述" autocomplete="off" class="layui-input"/>
                  </div>
                </div>
                <input id="operflg" type="hidden" value="add" />
              </div>
              <div id="operbtn">
                <div class="layui-btn-group">
                  <button id="add" class="layui-btn layui-btn-normal" type="submit" lay-submit lay-filter="addclass">确定</button>
                </div>  
              </div>
          </form>
        </div>
  </div>
</div>
<div id="bgindex">
    <div class="layui-btn-group toolsbar">
        <button id="addclassify" class="layui-btn layui-btn-normal">添加</button>
        <button id="reflushclassify" class="layui-btn layui-btn-normal">刷新</button>
    </div>
      <table class="layui-hide" id="classifytb" lay-filter="classifytb"></table> 
</div>
<script type="text/html" id="operbar">
    <a class="layui-btn layui-btn-xs layui-btn-normal" lay-event="edit">编辑</a>
    <a class="layui-btn layui-btn-danger layui-btn-xs layui-btn-normal" lay-event="del">删除</a>
  </script>
<script>
    layui.use(['table','form'], function(){
        var table = layui.table;
        //表单
        var form=layui.form;
        table.render({
          elem:'#classifytb'
          ,url:'/backstage/gettagsjson'
          ,cellMinWidth: 200
          ,toolbar:'#toolsbar'
          ,cols:[[
             {checkbox: true, fixed: false}
            ,{field:'id',title:'ID',sort:true,fixed:false}
            ,{field:'name',title:'标签名称'}
            ,{field:'classname',title:'所属分类'}
            ,{field:'brief',title:'描述'}
            ,{fixed: 'right', title:'操作', toolbar: '#operbar'}
          ]]
          ,id:'classtable'
        });
        $('#reflushclassify').click(function(){
          table.reload('classtable', 'data');
        });
        //表格操作栏
        table.on('tool(classifytb)',function(obj){
          if(obj.event === 'del'){
            layer.confirm('真的删除'+obj.data.name+'么?', function(index){
              
              var urlstr='/backstage/delclassify?id='+obj.data.id;
              $.ajax({
                url:urlstr,
                type:'get',
                async:true,
                timeout:5000,
                success:function(data){
                  if(data.code==1){
                    obj.del();
                  }else{
                      layer.open({
                      title: '删除失败',
                      content: data.msg
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
          else if(obj.event==='edit'){
            $('#operflg').val("edit");
             $('#classid').val(obj.data.id);
            $('#name').val(obj.data.name);
            $('#classname').val(obj.data.pid);
            $('#brief').val(obj.data.brief);
            $('#bkzz').css("display","block");
            form.render();
          }
        });  
        
        form.on('submit(addclass)',function(data){
          console.log(data);
          var jd={
            "id":parseInt($('#classid').val()),
            "pid":parseInt(data.field.classname),
            "name":data.field.name,
            "brief":data.field.brief
          }
         
          var loading = layer.msg('加载中...', {icon: 16, shade: 0.3, time:0});
          var flg=$('#operflg').val();
          var urlstr='/backstage/addtag';
          if(flg==="edit"){
            urlstr='/backstage/modclass';
          }
          $.ajax({
            url:urlstr,
            dataType:'json',
            contentType:'application/json',
            type:'post',
            async:true,
            timeout:5000,
            data:JSON.stringify(jd),
            success:function(data){
              layer.close(loading);
              if(data.code==1){
                //关闭当前添加界面
                $('#bkzz').css("display","none");
                
                //执行重载表格
                table.reload('classtable', 'data');
              }else{
                layer.open({
                  title: '错误',
                  content: data.msg
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

        form.render();

        $('#addclassify').click(function(){
            $('#operflg').val("add")
            $('#classid').val("");
            $('#classname').val("");
            $('#name').val("");
            $('#icon').val("");
            $('#brief').val("");
            $('#bkzz').css("display","block");
            form.render();
        });
    });
    
   
    $('#close').click(function(){
        $('#bkzz').css("display","none");
    });
    $('.closebtn').click(function(){
        $('#bkzz').css("display","none");
    });
</script>
{{end}}