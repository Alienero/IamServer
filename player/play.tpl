<!doctype html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <!--播放器-->
    <meta charset="utf-8">
    <title>重邮自由软件组织</title>
    <!--播放器-->
    <meta http-equiv="X-UA-Compatible" content="edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
	<!--插件引入-->
	<link type="text/css" rel="stylesheet" href="lib/video-js.css">
	<script src="lib/video.js"></script>

	<link type="text/css" rel="stylesheet" href="lib/style.css">
	<script src="lib/CommentCoreLibrary.js"></script>

	<script src="lib/jquery.js"></script>

	<!--样式文件-->
	<link type="text/css" rel="stylesheet" href="src/css/auto_combine_e0f64_1b6606a.css">
	<style>
		.mod-room div.mod-inner{
			margin:0 auto;
			width:1100px;
		}

		#my-player{
			position: absolute;
			z-index: 2147483648;
			top: 0;
			left: 0;
		}
		#text-box{
			width: 300px;
			height: 100px;
			position: relative;
		}
		#text-box textarea{
			margin: 5px;
			width: 240px;
			height: 60px;
		}
		#btn-send{
			width: 40px;
			height: 25px;
			color: white;
			background: #1916D8;
			text-align: center;
			position: absolute;
			right: 5px;
			line-height: 25px;
			bottom: 8px;
			cursor: pointer;
		}
		.text-display-line{
			margin: 5px;
			font-size: 14px;
		}
		.text-display-line b{
			color: #1DBACA;
		}
		.text-display-line i{
			color: #827B7B;
		}
		#text-display{
			overflow: scroll;
		}
	</style>
</head>
<body class="on-special w-1420">
<!-- E 通用头部 -->
<!-- S 通用侧栏 -->
<!-- 侧栏s -->
<div class="mod-sidebar" style="position: absolute; top: 80px;">
    <div class="sidebar-hide">
        <ul class="sidebar-icon-list">

        </ul>
        <span id="sidebar-show-btn" class="arrow-btn"></span>
    </div>
    <div class="sidebar-show" style="left: -252px;">
        <div class="m">
            <h3 class="m-title">关于重邮自由软件组织</h3>

        </div>
        <div class="m">
            <h3 class="m-title">最近项目</h3>

        </div>
        <div class="m">
            <h3 class="m-title"> 参与人员</h3>

        </div>
        <div class="">

        </div>

        <span id="sidebar-hide-btn" class="arrow-btn"></span>
    </div>
</div>
<!-- 侧栏e -->    <!-- E 通用侧栏 -->
<div class="mod-room special-room bg-edgclearlove" id="liveRoom">

    <div class="mod-inner">
        <!-- S 直播间顶部 -->
          <div style="color:white;font-weight:900;font-size:large">
          <pre >
       Powered by                                        
                 ___ _           _ _     _                
                / __\ | _____  _(_) |__ | | ___         ,---.               ,---.     ,---.|                        
               / _\ | |/ _ \ \/ / | '_ \| |/ _ \        |__. ,---.,---.,---.`---.,---.|__. |--- . . .,---.,---.,---.
              / /   | |  __/>  <| | |_) | |  __/        |    |    |---'|---'    ||   ||    |    | | |,---||    |---'
              \/    |_|\___/_/\_\_|_.__/|_|\___|        `    `    `---'`---'`---'`---'`    `---'`-'-'`---^`    `---'
                                                Team    
                    <a href="https://github.com/FlexibleBroadband">https://github.com/FlexibleBroadband</a>
            </pre>
    </div>
        <div class="room-hd clearfix">
        <div class="room-hd clearfix">
            <div class="host-pic">
                    <span class="pic-clip">
                        <img id="avatar-img" src="src/images/logo1.png" alt="{{.HostName}}">
                    </span>
            </div>
            <div class="host-info">
                <h1 class="host-title">{{.HostName}}</h1>

                <div class="host-detail">
                    <span class="host-name"></span>
                      <span class="host-channel">
                       直播<a href="" class="host-spl clickstat" eid="click/zhibo/zbxx/yxdj" eid_desc="点击/直播间/主播信息/游戏">{{.RoomName}}</a>
                      </span>
                    <span class="host-spectator"><i class="g-icon icon-host"></i><em class="host-spl" id="live-count">{{.LiveCount}}</em>个观众</span>
                </div>
            </div>
        </div>
        <!-- E 直播间顶部 -->
        <!-- S 直播间主体 -->
        <div class="room-bd">
            <!-- S 播放器 -->
            <div class="room-player-wrap">
	            <video id="my-video" class="video-js" controls preload="auto" width="800px" height="530px"
	                   poster="" data-setup="{}" autoplay="autoplay">
		            <source src="/live" type='video/flv'>
	            </video>
	            <div id="my-player" class="abp" style="width:800px; height:500px; ">
		            <a id="my-comment-stage" class="container"></a>
	            </div>
            </div>
            <!-- E 播放器 -->
            <!-- S 右边栏 -->
            <div class="room-sidebar">
                <div class="chat-room" id="chatRoom">
                    <div class="chat-room__hd">互动聊天</div>
                    <div class="chat-room__bd" id="text-display">
                    </div>
                    <!--输入框    -->
	                <div class="chat-room__hd">参与聊天： </div>
	                <div id="text-box">
                        <input id="text-user-name" type="text" placeholder="输入用户名"/>
		                <textarea id="text-area-box" placeholder="来吐槽吧！"></textarea>
						<botton id="btn-send">发送</botton>
	                </div>
                </div>
            </div>
        </div>
        <!-- E 直播间主体 -->
    </div>
</div>
<script type="text/plain" id="textTpl">
	<p class="text-display-line">
		<b>{name}</b>
		说:
		<i>{content}</i>
	</p>
</script>
<!--基础js-->
<script src="src/js/headerFunction_a684ab5.js" data-fixed="true"></script>
<script src="src/js/sidebarFunction_5925210.js" data-fixed="true"></script>

<!--自定义js-->
<script src="src/js/function.js"></script>
</body>
</html>