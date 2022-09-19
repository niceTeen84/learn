#!/bin/bash

# curl 通过代理访问
curl -x 'http://<your proxy server address>' -v https://www.google.com

# 仅返回 Headers
curl -I https://www.bing.com

# 生成 cookie 到文件
curl -c cookie.txt https://httpbin.org/cookies/set/cookiename/cookievalue