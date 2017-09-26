{{template "header.tpl"}}
<style>
img{max-width:100%}
</style>
<div>
	<button class="button icon-navicon" data-target="#nav-link6">
</button>
<ul class="nav nav-menu nav-inline nav-split nav-justified nav-big nav-navicon" id="nav-link6">
	<li><a href="#">图集查看</a></li>
</ul>
</div>
<div class="container-layout">
	{{range .info}}
    <div class="x12">
        <img style="display:block; margin:0 auto;" src="http://owgrjeqpr.bkt.clouddn.com/{{.Names}}">
        </div>
    {{end}}
</div>
</body>
</html>
