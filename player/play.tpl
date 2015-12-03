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
<div class="duya-header" id="duya-header">
    <div class="duya-header-wrap clearfix">
        <div class="duya-header-bd clearfix">
            <h1 id="duya-header-logo">
                <a href="" class="clickstat" eid="click/navi/logo" eid_desc="点击/导航/logo">
	                <img
                        src="src/images/112213.png" alt="重邮直播" title="重邮直播">
                </a>
            </h1>

            <div class="duya-header-nav">
                <span class="hy-nav-link"><a href="" class="hy-nav-title clickstat" eid="click/navi/home"
                                             eid_desc="点击/导航/首页">首页</a></span>
                <span class="hy-nav-link"><a href="" class="hy-nav-title clickstat" eid="click/navi/home"
                                             eid_desc="点击/导航/首页">直播</a></span>
                <span class="hy-nav-link"><a href="" class="hy-nav-title clickstat" eid="click/navi/home"
                                             eid_desc="点击/导航/首页">技术社区</a></span>
                <span class="hy-nav-link"><a href="" class="hy-nav-title clickstat" eid="click/navi/home"
                                             eid_desc="点击/导航/首页">代码托管</a></span>
                <span class="hy-nav-link"><a href="" class="hy-nav-title clickstat" eid="click/navi/home"
                                             eid_desc="点击/导航/首页">弹性宽带</a></span>
            </div>
            <!--<div class="duya-header-search clearfix">-->
            <!--<form method="get" id="searchForm_id" name="navSearchForm" action="">-->
            <!--<input type="text" name="hsk" value="开源、极客、直播" autocomplete="off">-->
            <!--<button type="submit" class="btn-search clickstat" eid="click/search/searchbutton"-->
            <!--eid_desc="点击/搜索/搜索按钮"></button>-->
            <!--</form>-->
            <!--</div>-->
            <div class="duya-header-control clearfix">
                <div class="hy-nav-right un-login" style="display: block;">
                    <div class="hy-nav-title hy-nav-title-login">
                        <i class="hy-nav-icon hy-nav-login-icon"></i>
                        <span class="hy-nav-text">
                            <a href="" class="clickstat title" id="nav-login" eid="click/navi/sign" eid_desc="点击/导航/登录">登录</a>
                            <span class="un-login" style="display: inline;">|</span>
                            <span class="un-login" style="display: inline;"><a href="" class="clickstat register-btn"
                                                                               id="nav-regiest" target="_blank"
                                                                               eid="click/navi/login"
                                                                               eid_desc="点击/导航/注册">注册</a></span>
                        </span>
                    </div>
                </div>
                <div class="hy-nav-right success-login nav-user">
                    <div class="hy-nav-title"><span id="login-username" class="username"></span></div>
                    <div class="nav-expand-list">
                        <i class="arrow"></i>

                        <div class="nav-host-link">

                            <a href="">个人中心</a>
                            <a href="">我的消息</a>
                            <a href="" id="nav-loggout">退出</a>
                        </div>
                    </div>
                </div>
                <div class="hy-nav-right nav-subscribe success-login">
                    <div class="hy-nav-title">
                        <i class="hy-nav-icon hy-nav-subscribe-icon"></i>
                        <a class="title">订阅</a>
                    </div>
                    <div class="nav-expand-list">
                        <i class="arrow"></i>

                        <div class="subscribe-hd clearfix" style="display:none;">
                            <h5>我订阅的有<em class="subscribe-key">0</em>个正在直播</h5>

                        </div>
                        <div class="subscribe-bd" style="display:none;">
                            <ul class="subscribe-list">
                            </ul>
                            <a class="subscribe-all" href="">全部订阅</a>
                        </div>
                        <div class="mod-list-more">
                            <div class="more-loading">
                                <i class="icon-loading"></i>
                                <em>正在加载您的订阅...</em>
                            </div>
                            <div class="more-empty">
                                <i class="icon-empty"></i>

                                <p>暂无订阅的直播。<br><em>你可以在主播的播放页进行订阅喔！</em></p>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="hy-nav-right">
                    <div class="hy-nav-title nav-history" id="nav-history">
                        <i class="hy-nav-icon hy-nav-history-icon"></i>
                        <a class="title">历史</a>
                    </div>
                    <div class="nav-expand-list nav-expand-history">
                        <i class="arrow"></i>

                        <div class="load-history">
                            <div class="mod-list-more">
                                <div class="more-loading">
                                    <i class="icon-loading"></i>
                                    <em>正在加载您的历史...</em>
                                </div>
                            </div>
                        </div>
                        <a href="" style="display:none">更多</a>
                    </div>
                </div>
                <div class="hy-nav-right no-border">
                    <div class="hy-nav-title">
                        <i class="hy-nav-icon hy-nav-zhubo-icon"></i>
                        <a href="" class="clickstat title" eid="click/navi/shower" eid_desc="点击/导航/我要做主播">成为主播</a>
                    </div>
                </div>
            </div>
            <div class="duya-header-tips">
                <em class="close-tips">×</em>

                <p class="tips-top">有新的直播内容，<span class="refresh-current">刷新看看</span></p>

                <div id="refresh-wrap">
                    <h5>你订阅的主播信息</h5>

                    <div class="refresh-tips"></div>
                    <h6><input type="checkbox" id="checkTips"><label for="checkTips">今天内不提醒</label></h6>
                </div>
            </div>
        </div>
    </div>
</div>
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
          <div style="color:white;font-weight:900;font-size:large"><pre >
                Powered by
                           ___ _           _ _     _
                          / __\ | _____  _(_) |__ | | ___
                         / _\ | |/ _ \ \/ / | '_ \| |/ _ \
                        / /   | |  __/>  <| | |_) | |  __/
                        \/    |_|\___/_/\_\_|_.__/|_|\___|
                                                          Team
                         <a href="https://github.com/FlexibleBroadband">https://github.com/FlexibleBroadband</a></pre></div>
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
                    <span class="host-level" style="background-position:0 -380px"></span>
                    <span class="host-name"></span>
                      <span class="host-channel">
                       直播<a href="" class="host-spl clickstat" eid="click/zhibo/zbxx/yxdj" eid_desc="点击/直播间/主播信息/游戏">{{.RoomName}}</a>
                      </span>
                    <span class="host-spectator"><i class="g-icon icon-host"></i><em class="host-spl" id="live-count">{{.LiveCount}}</em>个观众</span>
                </div>
            </div>
            <div class="host-control">
                <div class="mobile-entrance">
                    <em class="mobile-entrance-icon"></em>

                    <p>扫一扫<em class="entrance-arrow"></em></p>

                    <div class="entrance-expand" id="pop-box-scan"></div>
                </div>
                <div class="share-entrance">
                    <em class="share-entrance-icon"></em>

                    <p>分享至<em class="entrance-arrow"></em></p>

                    <div class="entrance-expand pop-box-social">
                        <div class="entrance-expand-bor"></div>
                        <em class="arrow"></em>

                        <p class="share-entrance-p">喊好友一起看大神敲代码吧</p>

                        <div class="social-media">
                            <a title="分享到新浪微博" class="wb-icon social-media-em hiido_stat" target="_blank" href=""></a>
                            <a title="分享到腾讯微博" class="qqwb-icon social-media-em hiido_stat" target="_blank" href=""
                               hiido_code=""></a>
                            <a title="分享到QQ空间" class="qzone-icon social-media-em hiido_stat" target="_blank"
                               href=""></a>
                        </div>
                        <input type="text" readonly="readonly" value="" id="flash-link" style="display:none">

                        <div class="share-copy-btn" data-clipboard-target="flash-link">复制播放器地址</div>
                    </div>
                </div>
                <div class="btn-subscribe">
                    <div class="subscribe-control" id="yyliveRk_game_newsBut"><em></em>订阅</div>
                    <em class="arrow"></em>

                    <div id="activityCount">0</div>
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