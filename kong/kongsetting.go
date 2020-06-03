package kong

var upstreamJson = []byte(`
{
		"tags": [],
		"name": "",
		"hash_on": "none",
		"healthchecks": {
			"active": {
				"unhealthy": {
					"http_statuses": [ 429, 500, 501, 502, 503, 504, 505 ],
					"tcp_failures": 1,
					"timeouts": 1,
					"http_failures": 1,
					"interval": 5
				},
				"type": "http",
				"http_path": "/kong/healthchecks",
				"timeout": 1,
				"healthy": {
					"successes": 1,
					"interval": 5,
					"http_statuses": [ 200, 302 ]
				},
				"https_verify_certificate": true,
				"concurrency": 1
			},
			"passive": {
				"unhealthy": {
					"http_statuses": [ 429, 500, 501, 502, 503, 504, 505 ]
				},
				"healthy": {
					"http_statuses": [ 200, 302 ]
				},
				"type": "http"
			}
		},
		"hash_on_cookie_path": "/",
		"hash_fallback": "none",
		"slots": 10000
	}`)
