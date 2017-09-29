{{template "admin/header.tpl"}}
	<body>
		{{template "admin/meau.tpl"}}
		<div class="admin">
			<form method="post" enctype="multipart/form-data">
				<div class="form-group">
					<div class="label">
					<label for="tag_name">标签名称</label>
					</div> 
					<div class="field">
					<input type="text"  value="" class="input" id="tag_name" name="tag_name" size="30" data-validate="required:必填" placeholder="标签名称" />
					</div> 
				</div>
                <div class="form-group">
					<div class="label">
					<label for="types">标签分组</label>
					</div> 
					<div class="field">
                        <select class="input" name="types">
							<option value="1">服饰类型</option>
                            <option value="2">美女特征</option>
                            <option value="3">模特人物</option>
                            <option value="4">职业类型</option>
							<option value="5">动物类型</option>
                            <option value="6">套图</option>
                            <option value="7">搞笑</option>
                            <option value="8">素材</option>
						</select>
					</div> 
				</div> 
				<div class="form-button"> <button class="button" type="submit">提交</button> </div>
			</form>
			<br />
			</div>
	</body>

</html>