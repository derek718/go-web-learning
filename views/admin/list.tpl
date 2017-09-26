<!DOCTYPE html>
<html lang="zh-cn">

	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no" />
		<meta name="renderer" content="webkit">
		<title>集讯后台管理-后台管理</title>
		<link rel="stylesheet" href="/static/css/admin/pintuer.css">
		<link rel="stylesheet" href="/static/css/admin/admin.css">
		<script src="/static/js/jquery.js"></script>
		<script src="/static/js/admin/pintuer.js"></script>
		<script src="/static/js/admin/admin.js"></script>
		<link type="image/x-icon" href="http://www.pintuer.com/favicon.ico" rel="shortcut icon" />
		<link href="http://www.pintuer.com/favicon.ico" rel="bookmark icon" />
	</head>

	<body>
		<div class="lefter">
			<div class="logo">
				<a href="http://www.pintuer.com" target="_blank"><img src="images/logo.png" alt="后台管理系统" /></a>
			</div>
		</div>
		<div class="righter nav-navicon" id="admin-nav">
			<div class="mainer">
				<div class="admin-navbar">
					<span class="float-right">
                    <a class="button button-little bg-main" href="#">前台首页</a>
                    <a class="button button-little bg-yellow" href="login.html">注销登录</a>
                </span>
					<ul class="nav nav-inline admin-nav">
						<li>
							<a href="/admin" class="icon-home"> 开始</a>
							<ul>
								<li  class="active"><a href="/admin/tube">管理员</a></li>
								
							</ul>
						</li>
						<li>
							<a href="/admin/column/index" class="icon-th-list">栏目管理</a>
							<ul>
								<li><a href="/admin/column/index">栏目列表</a></li>
								<li class="active"><a href="/admin/column/add">栏目添加</a></li>
							</ul>
						</li>
						<li  class="active">
							<a href="/admin/p/i" class="icon-cog">图讯</a>
							<ul>
								<li class="active"><a href="/admin/p/i" >图片列表</a></li>
								<li ><a href="#">系统设置</a></li>
								<li><a href="#">会员设置</a></li>
								<li><a href="#">积分设置</a></li>
							</ul>
						</li>
					</ul>
				</div>
				<div class="admin-bread">
					<span>您好，{{.username}}，欢迎您的光临。</span>
					<ul class="bread">
						<li><a href="/admin" class="icon-home"> 开始</a></li>
						<li>管理员列表</li>
					</ul>
				</div>
			</div>
		</div>

		<div class="admin">
			<form method="post">
				<div class="panel admin-panel">
					<div class="panel-head"><strong>图片列表</strong></div>
					<div class="padding border-bottom">
						<input type="button" class="button button-small checkall" name="checkall" checkfor="id" value="全选" />
						<input type="button" class="button button-small border-green" value="添加文章" />
						<input type="button" class="button button-small border-yellow" value="批量删除" />
						<input type="button" class="button button-small border-blue" value="回收站" />
					</div>
					<table class="table table-hover">
						<tr>
							<th width="5">选择</th>
							<th width="10">ID</th>
							<th width="120">图片名称</th>
							<th width="100">操作</th>
						</tr>
						{{range .list}}
						<tr>
							<td>
								<input type="checkbox" name="id" value="1" />
							</td>
							<td>{{.Id}}</td>
							<td>{{.Title}}</td>
							<td>
							<!--<a class="button border-blue button-little" href="#">修改</a>
							<a class="button border-yellow button-little" href="#" onclick="{if(confirm('确认删除?')){return true;}return false;}">删除</a>-->
							<a class="button border-green button-little" href="/admin/p/e/{{.Id}}">详情</a>
							</td>
						</tr>
						{{end}}
						
					</table>
					<div class="panel-foot text-center">
						{{str2html .pager}}
					</div>
				</div>
			</form>
		</div>

	</body>

</html>