{{define "artfilepage"}}
    {{range.LineItems}}
        <li class="layui-timeline-item">
          <i class="layui-icon layui-timeline-axis">&#xe63f;</i>
          <div class="layui-timeline-content layui-text">
            <h3 class="layui-timeline-title">
                {{.CreateTime}}
            </h3>
            <p>
              <a href="/article/detail?artid={{.ID}}">{{.Name}}</a>
            </p>
          </div>
        </li>
    {{end}}
{{end}}