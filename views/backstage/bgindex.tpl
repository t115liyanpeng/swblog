{{define "backstagedefault"}}
 <div id="bgindex">
        <div id="dashconter">
            <div  class="dsashboard layui-bg-blue">
                <div id="dcontent">
                    <div class="dtitle">文章总数</div>
                    <div class="dcount">{{.ArtCount}}</div>
                    <div class="dbotom"><span>点击进入文章管理</span></div>
                </div>

            </div>
            <div  class="dsashboard layui-bg-green">
                <div id="dcontent">
                    <div class="dtitle">评论总数</div>
                    <div class="dcount">{{.CommentCount}}</div>
                    <div class="dbotom"><span>点击进入评论管理</span></div>
                </div>
            </div>
            <div  class="dsashboard layui-bg-orange">
                <div id="dcontent">
                    <div class="dtitle">访问总数</div>
                    <div class="dcount">{{.AllSee}}</div>
                    <div class="dbotom">今日访问数量：{{.TodaySee}}</div>
                </div>
            </div>
        </div>
        <div class="clear"></div>
</div>
{{end}}