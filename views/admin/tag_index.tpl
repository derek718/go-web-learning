{{template "admin/header.tpl"}}

	<body>
		{{template "admin/meau.tpl"}}

		<div class="admin">
		<div class="panel admin-panel">
			<div class="panel-head"><strong>内容列表</strong></div>
					<div class="padding border-bottom">
						<input onclick="javascript:window.location.href='/admin/tag/add'" type="button" class="button button-small border-green" value="添加标签" />
					</div>
			<div class="panel-foot text-center">
            {{range .list}}
                <button class="button">{{.TagName}}</button>
            {{end}}
			</div>
		</div>
        <div class="panel-foot text-center">
				{{str2html .pager}}
			</div>
		</div>
	</body>

</html>