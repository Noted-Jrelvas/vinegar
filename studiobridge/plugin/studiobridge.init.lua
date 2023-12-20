local Root = script.Parent

--This setting is defined by a tiny lua plugin created at runtime by Vinegar. Repeat since plugin load order isn't guaranteed.
local port: number
repeat
	port = plugin:GetSetting("VinegarBridge_ServerPort")
	task.wait()
until port ~= nil

--Ensures that no stale value remains for later.
plugin:SetSetting("VinegarBridge_ServerPort", nil)


local client = require(Root.client)
client.port = port


client.backendInfo()