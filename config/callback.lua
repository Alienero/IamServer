-- example callback lua file.

-- module 
local callback = {}

-- var
callback.mapping = {["123"]="master"}

-- callback functions

function callback.addr_mapping(private)
	return callback.mapping[private]
end

function callback.rtmp_access_check(remote, local_addr, appname, path)
	print("remote:",remote)
	print("local:",local_addr)
	print("appname:",appname)
	print("path:",path)

	if appname=="live" and path == "123" then
		return true
	else
		return false
	end
end

function callback.flv_access_check(remote, url, path,froms, cookies)
	print("remote",remote)
	print("url:",url)
	print("path:",path)

	if path == "/live/master.flv" then
		return true
	else
		return false
	end
end

function callback.im_access_check(remote, url, path,forms, cookies)
	print("remote",remote)
	print("url:",url)
	print("path:",path)

	-- room id.
	print(forms["room_id"])

	if path == "/im/master.room" then
		-- get user name.
		local m = require("libs.html")
		local name = m.html_escape(forms["name"])
		if name == "" then
			name = "gust"
		end
		return name,1,true
	else
		return "",0,false
	end
end

return callback