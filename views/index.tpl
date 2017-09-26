{{template "header.tpl"}}
<div>
	<button class="button icon-navicon" data-target="#nav-link6">
</button>
<ul class="nav nav-menu nav-inline nav-split nav-justified nav-big nav-navicon" id="nav-link6">
	<li><a href="#">首页</a></li>
</ul>
</div>
<div class="container-layout">
	
  {{range $i,$v:=.list}}
  <div class="line">
		
        {{range $ii,$vv:=$v}}
        <div class="x2 xl6 xs4 xm4 xb4">
          <img width="200" height="200" src="http://owgrjeqpr.bkt.clouddn.com/{{.Thumbs}}?imageView2/2/w/200/h/200/interlace/1/q/100"/>
          <br/>
          <a href="/i/{{.Id}}">{{.Title}}</a>
        </div>
        {{end}}
    </div>
    <div style="margin-top:10px"></div>
    {{end}}  
</div>

<div class="container-layout">
  <div class="line">
      <div class="x12">
        <div class="panel-foot text-center">
						{{str2html .pager}}
					</div>
      </div>
  </div>
</div>
</body>
</html>
