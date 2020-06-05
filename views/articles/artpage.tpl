{{define "artpage"}}
    {{range.}}
    <div class="m_continer">
        <div class="w_title">
        <a href="/article/detail?artid={{.ID}}">{{.Name}}</a>
        </div>
        <div class="w_tags">                                                       
        <span>作者：{{.Author}}</span>
        &nbsp;
        &nbsp;
        &nbsp;
        &nbsp;
        <span>分类：{{.Classify}}</span>
        &nbsp;
        &nbsp;
        &nbsp;
        &nbsp;
        <span>标签：{{.Tag}}</span>
        &nbsp;
        &nbsp;
        &nbsp;
        &nbsp;
        <span>日期：{{.CreateTime}}</span>
        </div>
        <div class="w_content">
        <span>
            {{.Content}}
        </span>
        </div>
        <div class="w_footer">
            <i class="layui-icon layui-icon-read">
                <span class="likenum">{{.Click}}</span>
                <i class="artid" style="display: none;">{{.ID}}</i>
            </i>
            <!--<a>评论(0)</a>-->
            <i class="layui-icon layui-icon-heart like">
                <span class="likenum">{{.Like}}</span>
                <i class="artid" style="display: none;">{{.ID}}</i>
            </i>
            <a href="/article/detail?artid={{.ID}}">开始阅读</a>
        </div>
    </div>
    {{end}}
    <script>
        layui.use('layer',function(){
                //点赞
                $(".like").on("click", function() {
                      var like=$(this);
                      var sp= like.find('i');
                      var num=parseInt(sp.text());
                      var strurl='/article/like?id='+num;
                      var loading = layer.msg('加载中...', {icon: 16, shade: 0.3, time:0});
                    $.ajax({
                        type:'get',
                        url:strurl,
                        async:true,
                        success:function(data){
                            layer.close(loading);
                            if (data.code==1){
                                like.attr("class","layui-icon layui-icon-heart-fill like");
                                var splike= like.find('span');
                                var cnt=parseInt(splike.text())+1;
                                splike.text(cnt);
                                layer.msg(data.msg)
                                
                            }else if(data.code==0){
                                layer.open({
                                title: '提示',
                                content: data.msg,
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
                      
                    
                });
        });
    </script>
{{end}}