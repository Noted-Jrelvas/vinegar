local URL_FORMAT = "http://localhost:%d/"

local HttpService = game:GetService("HttpService")

local client = {}

function client.backendInfo()
	local response = HttpService:RequestAsync({
		Url = string.format(URL_FORMAT, client.port),
		Method = "GET"
	})

	print(response)
end

return client