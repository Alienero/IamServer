/**
 * Created by Mr_Harry on 15/12/1.
 */

var user_name="";
var room_id="master"

var CM = new CommentManager(document.getElementById('my-comment-stage'));
// 创建一个Socket实例
var socket = new WebSocket('ws://'+window.location.host+'/im/'+room_id+".room");

var xmlHttpRequest;

//初始化操作
var C = {};
C.$text_area_box = $('#text-area-box');
C.$btn_send = $('#btn-send');
C.$text_display = $('#text-display');

var textTplStr = $('#textTpl').text();
var heightInit = document.getElementById('text-display').scrollHeight;


CM.init();
// 调整弹幕速度
CM.options.global.scale=2.2*CM.options.global.scale; 
// 先启用弹幕播放（之后可以停止）
CM.start();

// 打开Socket
socket.onopen = function(event) {
	console.log('socket open');

	document.onkeydown=function(e){
		console.log("press enter");
		if(e.keyCode==13 && e.ctrlKey){
			var text = C.$text_area_box.val();
			C.$text_area_box.val(text);
		}else if (e.keyCode==13){
			e.preventDefault();
			send();
		}
	}

	C.$btn_send.on('click', function(e) {
		e.preventDefault();
		send();
	});

	// 监听消息
	socket.onmessage = function(event) {
		console.log('Client received a message',event);
		var msg=JSON.parse(event.data);
		console.log(msg.user);
		showMess(msg.user,msg.playload);
	};

	// 监听Socket的关闭
	socket.onclose = function(event) {
		console.log('socket closed',event);
		alert("talk room connetion closed, please flush the page.")
	};
};

var isFullScreen =false
videojs('my-video').on('fullscreenchange', function() {
	if(!isFullScreen){
		isFullScreen = true;
		jQuery('#my-player').css({
			"width":"100%",
		 	"height":"90%",
		});
		CM.setBounds();
	}else{
		isFullScreen = false;
		jQuery('#my-player').css({
			"width":"800px",
		 	"height":"500px",
		});
		CM.setBounds("800px", "500px");
	}
  	// alert(videojs('my-video').height());
});

var i=setInterval(function(){
	var url = "/count?room_id="+room_id;       
	//1.创建XMLHttpRequest组建    
	xmlHttpRequest = new XMLHttpRequest();      
        
    	//2.设置回调函数    
    	xmlHttpRequest.onreadystatechange = setLiveCount;    
        
    	//3.初始化XMLHttpRequest组建    
   	 xmlHttpRequest.open("GET",url,true);    
        
    	//4.发送请求    
    	xmlHttpRequest.send(null);      
    	}, 5*1000); // update live count.

function send(){
	if(socket==undefined){

	}
	if(user_name==""){
		getUserName();
		if(user_name==""){
			alert("请输入用户名：")
			return
		}
	}
	if(socket.readyState==socket.CLOSED){
		alert("talk room connetion closed, please flush the page.")
		return 
	}
	var text = C.$text_area_box.val();
	var msg = {"user":user_name,"time":0,"type":0,"color":"","playload":""};
	msg.playload=text;
	socket.send(JSON.stringify(msg));
	C.$text_area_box.val("");
}

/**
 * 渲染数据到页面
 * @param text
 */
function showMess (name,text){
	//发送弹幕
	var danmaku = {
		"mode": 1,
		"text": text,
		"stime": 0,
		"size": 25,
		"color": 0xffffff
	};
	CM.send(danmaku);

	//添加到聊天框
	var htmlStr = textTplStr.replace(/\{\w+}/g, function(tpl) {
		return {
			'{name}': name,
			'{content}': text
		}[tpl];
	});

	C.$text_display.append(htmlStr);

	//滚动条保持在最底部
	var heightCur = document.getElementById('text-display').scrollHeight;
	if( heightCur > heightInit){
		var scrollTop = heightCur - heightInit;
		C.$text_display.scrollTop(scrollTop);
	}
}

function getUserName(){
	user_name=$("#text-user-name").val();
}

function setLiveCount(){
	if(xmlHttpRequest.readyState == 4 && xmlHttpRequest.status == 200){    
		var b = xmlHttpRequest.responseText;    
		// 反序列化
		var resp=JSON.parse(b);
		$("#live-count").text(resp.liveCount);
	}    
}
