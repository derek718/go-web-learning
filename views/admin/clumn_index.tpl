{{template "admin/header.tpl"}}

	<body>
		{{template "admin/meau.tpl"}}

		<div class="admin">
		
		<div class="padding border-bottom">
			<input type="button" onclick="javascript:window.location.href='/admin/column/add'" class="button button-small border-green" value="添加栏目" />
		</div>
		<div class="panel admin-panel">
			<table class="table table-hover">
				<tr>
					<th width="10%">栏目名称</th>
					<th width="5%">级别</th>
					<th width="40%">简介</th>
					<th width="15%">栏目图</th>
					<th width="5%">操作</th>
				</tr>
				{{range .list}}
				<tr>
					<td>{{.ClumnName}}</td>
					<td>
						{{if gt .Parentid 1}}
							二级栏目
						{{else}}
							一级栏目
						{{end}}
					</td>
					<td>{{.Desc}}</td>
					<td><img src="/{{.Thumbs}}"/></td>
					<td>
					<a class="button border-blue button-little" href="/admin/column/edit/{{.Id}}">修改</a>
					<a class="button border-yellow button-little" href="#" onclick="{if(confirm('确认删除?')){return true;}return false;}">删除</a>
					</td>
				</tr>
				{{end}}
			</table>
			<div class="panel-foot text-center">
				{{str2html .pager}}
			</div>
		</div>
		</div>
	</body>

</html>