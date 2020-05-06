{{define "artclassifypage"}}
    {{range.}}
    <li class="layui-timeline-item">
        <i class="layui-icon layui-timeline-axis">&#xe63f;</i>
        <div class="layui-timeline-content layui-text">
          <h3 class="layui-timeline-title">
              {{if eq .Pid 0}}
              <a href="/article/artclassifylst?index=1&classify={{.ID}}&title={{.Name}}">{{.Name}}</a>
              {{else}}
              <a href="/article/artclassifylst?index=1&tag={{.ID}}&title={{.Name}}">{{.Name}}</a>
              {{end}}
          </h3>
        </div>
    </li>
    {{end}}
{{end}}