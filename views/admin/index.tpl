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
                    <a class="button button-little bg-main" href="http://www.pintuer.com" target="_blank">前台首页</a>
                    <a class="button button-little bg-yellow" href="login.html">注销登录</a>
                </span>
					<ul class="nav nav-inline admin-nav">
						<li class="active">
							<a href="index.html" class="icon-home"> 开始</a>
							<ul>
								<li><a href="/admin/tube">管理员</a></li>
								
							</ul>
						</li>
						<li>
							<a href="/admin/column/index" class="icon-th-list">栏目管理</a>
							<ul>
								<li><a href="/admin/column/index">栏目列表</a></li>
								<li><a href="/admin/column/add">栏目添加</a></li>
							</ul>
						</li>
						<li>
							<a href="system.html" class="icon-cog">图讯</a>
							<ul>
								<li><a href="#">全局设置</a></li>
								<li class="active"><a href="#">系统设置</a></li>
								<li><a href="#">会员设置</a></li>
								<li><a href="#">积分设置</a></li>
							</ul>
						</li>
						
						
					</ul>
				</div>
				<div class="admin-bread">
					<span>您好，admin，欢迎您的光临。</span>
					<ul class="bread">
						<li><a href="index.html" class="icon-home"> 开始</a></li>
						<li>后台首页</li>
					</ul>
				</div>
			</div>
		</div>

		<div class="admin">
		
			<div class="panel">
						<div class="panel-head"><strong>系统信息</strong></div>
						<table class="table">
							<tr>
								<th colspan="2">服务器信息</th>
								<th colspan="2">系统信息</th>
							</tr>
							<tr>
								<td width="110" align="right">操作系统：</td>
								<td>Windows 2008</td>
								<td width="90" align="right">系统开发：</td>
								<td><a href="http://www.pintuer.com" target="_blank">拼图前端框架</a></td>
							</tr>
							<tr>
								<td align="right">Web服务器：</td>
								<td>Apache</td>
								<td align="right">主页：</td>
								<td><a href="http://www.pintuer.com" target="_blank">http://www.pintuer.com</a></td>
							</tr>
							<tr>
								<td align="right">程序语言：</td>
								<td>PHP</td>
								<td align="right">演示：</td>
								<td><a href="http://www.pintuer.com/demo/" target="_blank">demo/</a></td>
							</tr>
							<tr>
								<td align="right">数据库：</td>
								<td>MySQL</td>
								<td align="right">群交流：</td>
								<td><a href="http://shang.qq.com/wpa/qunwpa?idkey=a08e4d729d15d32cec493212f7672a6a312c7e59884a959c47ae7a846c3fadc1" target="_blank">201916085</a> (点击加入)</td>
							</tr>
						</table>
					</div>
			
		</div>		
	</body>

</html>