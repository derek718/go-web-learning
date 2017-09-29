{{template "admin/header.tpl"}}
	<body>
		{{template "admin/meau.tpl"}}
		<div class="admin">
			<form method="post" enctype="multipart/form-data">
				<div class="form-group">
					<div class="label">
					<label for="clumn_name">栏目名称</label>
					</div> 
					<div class="field">
					<input type="text"  value="{{.info.ClumnName}}" class="input" id="clumn_name" name="clumn_name" size="30" data-validate="required:必填" placeholder="栏目名称" />
					</div> 
				</div> 

				<div class="form-group">
					<div class="label">
						<label for="parentid">栏目级别</label>
					</div>
					<div class="field">

					{{if .info.Parentid}}
						<select class="input" name="parentid">
							<option value="0">一级栏目</option>
							{{range $.list}}
							<option {{if eq .Id $.info.Parentid}}
									selected="selected" 
									{{end}} value="{{.Id}}">{{.ClumnName}}</option>
							{{end}}
						</select>
					{{else}}
						<select class="input" name="parentid">
							<option value="0">一级栏目</option>
							{{range $.list}}
							<option  value="{{.Id}}">{{.ClumnName}}</option>
							{{end}}
						</select>
					{{end}}
					
						
					</div>					
				</div>

				<div class="form-group">
					<div class="label">
						<label for="desc">栏目简介</label>
					</div>
					<div class="field">
						<textarea name="desc" rows="5" class="input" placeholder="请输入简介" data-validate="required:必填">{{.info.Desc}}
						</textarea>
					</div>					
				</div>
				<div class="form-group">
					<div class="label">
						<label for="desc">栏目图标</label>
					</div>
					<div class="field">
						{{if .info.Thumbs}}
							<span>图片名称：{{.info.Thumbs}}</span><br/>
						{{end}}
						<a class="button input-file bg-blue" href="javascript:void(0);">+ 浏览文件
						<input size="100" type="file" name="uploadfile" id="uploadfile"/>
						</a>
					</div>					
				</div>

				<div class="form-group">
					<div class="label">
						<label for="desc">标签</label>
					</div>
					<div class="field">
						<div class="button-group border-main checkbox">
							
							{{range .tlist}}
								<label class="button tags">
									<input name="tags"  value="{{.Id}}" type="checkbox">{{.TagName}}
								</label>
							{{end}}
							
						</div>
					</div>					
				</div>
				
				<input type="hidden" name="_id" value="{{.info.Id}}"/>
				<div class="form-button"> <button class="button" type="submit">提交</button> </div>
			</form>
			<br />
					</div>
<script>
	var tagids = "{{.info.TagId}}"

	if (tagids != ""){
		_ids = tagids.split(",")
		$(".tags").each(function(){
			tagid = $(this).find("input").val()
		if (_ids.indexOf(tagid) >=0){
			$(this).addClass("active")
		}
	})
	}
	
</script>
	</body>

</html>