go.srs
======

[SRS(simple-rtmp-server)](https://github.com/winlinvip/simple-rtmp-server) written by google go language.

### ��Ʒ��λ

GO.SRS��SRS�Ķ�λ��̫һ�������Ի�����<br/>
SRS��Ҫ��RTMPԴվ�����HLS��ת�룬�����̣�ͨ��ת������ϵͳ��������<br/>
GO.SRS��Ҫ����������������Դվ�ͱ�Ե��֧��RTMP/RTMPT/RTMPE/RTMPS/HLS/DASH/RTMFP/RTSP�ȣ�֧�ֶ���̡�<br/>
����ͼ��ʾ��
<pre>
+---------------------------+    +------------------------------+
|     GO.SRS(��������)      +-->-+  SRSת��/Chnvideo��ת�뼯Ⱥ  |
|   (IPv4/IPv6/TCP/UDP)     |    +------------------------------+
| (Դվ/��Ե/������/�����) |                                    
| (RTMP/RTMPE/RTMPT/RTMPS)  |    +------------------------------+
|   (HTTP/HLS/HDS/DASH)     |-->-+  Chnvideo��¼/ʱ��/������    |
|     (RTSP,RTMFP)          |    +------------------------------+
+---------------------------+                                    
IPv4/IPv6: ͬʱ֧��IPv4/IPv6��
TCP��֧�ֻ���TCP��Э�飬Ʃ��RTMP��HTTPϵ�С�
UDP��֧�ֻ���UDP��Э�飬Ʃ��RTMFP�ȡ�
Դվ�ͱ�Ե��֧�ּ�Ⱥ��Ʃ��RTMPϵ�е�Դվ�ͱ�Ե��HTTPֻ��Ҫ֧��Դվ����Ե��NGINX�ȳ��췽������
RTMPϵ�У����������Ļ���Э�顣
RTSPϵ�У�֧��RTSP��Э�飬֧��һЩ����ͷ��
RTMFP��Adobe��FlashP2P������
SRSת�룺SRS������ffmpegת�룬Ϊת��Ŀ�Դ������
Chnvideo����ҵ����
</pre>

### GO��������

�ο�[http://golang.org/doc/install](http://golang.org/doc/install)<br/>
<strong>Step 1:</strong>����GO<br/>
https://code.google.com/p/go/downloads/list<br/>
<strong>Step 2:</strong>��ѹGO<br/>
<pre>
tar xf go1.2.linux-amd64.tar.gz
sudo ln -sf `pwd`/go /usr/local/go
</pre>
<strong>Step 3:</strong>���û�������<br/>
<pre>
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
</pre>
ע�⣺���л������������ÿ��Ա༭/etc/profile

### ���뷽��(Build)

<strong>Step 1:</strong> set GOPATH if not set<br/>
<pre>
export GOPATH=~/mygo
</pre>
<strong>Step 2:</strong> get and build srs<br/>
<pre>
go get github.com/winlinvip/go.srs/go_srs
</pre>
<strong>Step 3:</strong> start SRS <br/>
<pre>
$GOPATH/bin/go_srs
</pre>
ע�⣺���������go_srs��������GO�������������Զ�������

Beijing, 2014<br/>
Winlin

