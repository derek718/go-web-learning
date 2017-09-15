<!DOCTYPE html>
<html lang="zh-cn">

	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no" />
		<meta name="renderer" content="webkit">
		<title>拼图后台管理-后台管理</title>
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
						<li class="active">
							<a href="index.html" class="icon-home"> 开始</a>
							<ul>
								<li  class="active"><a href="/admin/tube">管理员</a></li>
								
							</ul>
						</li>
						<li>
							<a href="system.html" class="icon-cog"> 系统</a>
							<ul>
								<li><a href="#">全局设置</a></li>
								<li class="active"><a href="#">系统设置</a></li>
								<li><a href="#">会员设置</a></li>
								<li><a href="#">积分设置</a></li>
							</ul>
						</li>
						<li >
							<a href="content.html" class="icon-file-text"> 内容</a>
							<ul>
								<li><a href="#">添加内容</a></li>
								<li><a href="#">内容管理</a></li>
								<li><a href="#">分类设置</a></li>
								<li><a href="#">链接管理</a></li>
							</ul>
						</li>
						<li><a href="#" class="icon-shopping-cart"> 订单</a></li>
						<li><a href="#" class="icon-user"> 会员</a></li>
						<li><a href="#" class="icon-file"> 文件</a></li>
						<li><a href="#" class="icon-th-list"> 栏目</a></li>
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
	
					<table class="table table-hover">
						<tr>

							<th width="120">管理员名称</th>
							<th width="*">最后登录</th>
							<th width="100">IP</th>
							<th width="100">操作</th>
						</tr>
						{{range .listData}}
						<tr>

							<td>{{.Username}}</td>
							<td>{{date .Lastlogintime "Y-m-d H:i:s"}}</td>
							<td>{{.Lastloginip}}</td>
							<td>
							\
							</td>
						</tr>
						{{end}}
					</table>
					
				</div>
			</form>
			<br />
					</div>
	</body>

</html>