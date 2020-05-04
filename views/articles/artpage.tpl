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
            <!--<a>评论(0)</a>-->
            
            <a href="/article/detail?artid={{.ID}}">开始阅读</a>
        </div>
    </div>
    {{end}}
{{end}}