-- file's path: $GOPATH/src/github.com/Alienero/IamServer/test/lua/

-- module 
callback = {}

-- var
callback.mapping = {[""]="master",["master"]="master"}

-- callback functions

function callback.addr_mapping(public)
	return callback.mapping[public]
end

function callback.rtmp_access_check(remote, local_addr, appname, path)
	print("remote:",remote)
	print("local:",local_addr)
	print("appname:",appname)
	print("path:",path)

	if appname=="live" and path == "master" then
		return true
	else
		return false
	end
end

function callback.flv_access_check(remote, url, path,froms, cookies)
	print("remote",remote)
	print("url:",url)
	print("path:",path)

	if path == "/live/master" then
		return true
	else
		return false
	end
end

function callback.im_access_check(remote, url, path,froms, cookies)
	print("remote",remote)
	print("url:",url)
	print("path:",path)

	if path == "/im/master" then
		return "I am master",1,true
	else
		return "",0,false
	end
end

return callback