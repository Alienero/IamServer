# IamServer
simple live stream solution written in golang

## Warnning
It is not ready for production evn use. Some features aren't implemented yet. Documentation is missing.

## How to build a living room
### Download with binarries
- (Mac os x(at least 10.8))[https://github.com/Alienero/IamServer/releases/download/v0.0.1/mac-amd64-0.01.zip]
- (Windows at least windows 7 not support i386)[https://github.com/Alienero/IamServer/releases/download/v0.0.1/windws-amd64-0.01.zip]

### Let it work
```bash
cd mac-amd64-0.01/bin  &&  ./IamServer
```
So easy. It will listening at port 1935(RTMP default port) for pulishing and port 80 for play live streaming. For single mode, puslihing FMS is `rtmp://localhost/live` and Path(Key) is `123`. Also We provide some useful flags to config server.     
`-name=[string]` can change your default living room's title.      
 Now, You should use some RTMP pulishing software. Like [OBS](https://obsproject.com/). After configuration OBS's parameters(living FMS and KEY), you can play the live streaming in your browser. The web living room provide a talk room and danmaku.
 
##  License
 This project is under the MIT License. See the LICENSE file for the full license text.
