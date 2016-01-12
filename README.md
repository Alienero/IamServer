# IamServer
simple live stream solution written in golang

## Warnning
It is not ready for production evn use. Some features aren't implemented yet. Documentation is missing.

## How to build a live room
### Download with binarries
- [Mac os x (at least 10.8)](https://github.com/Alienero/IamServer/releases/download/v0.0.2a/v0.0.2.win.mac.zip)
- [Windows (at least windows 7 , it not support i386)](https://github.com/Alienero/IamServer/releases/download/v0.0.2a/v0.0.2.win.mac.zip)

### Let it work
```bash
cd mac-amd64-0.01/bin  &&  ./IamServer
```
So easy. It will listening at port 1935(RTMP default port) for pulishing and port 9090 for play live streaming. For single mode, puslihing FMS is `rtmp://localhost/live` and Path(Key) is `123`. Also We provide some useful flags to config server.     
`-name=[string]` can change your default live room's title.      
 Now, You should use some RTMP pulishing software. Like [OBS](https://obsproject.com/). After configuration OBS's parameters(live FMS and KEY), you can play the live streaming in your browser. The web live room provide a talk room and danmaku.

[中文内测版部署文档](https://ilulu.xyz/post/livebuild/)

##  License
 This project is under the MIT License. See the LICENSE file for the full license text.
