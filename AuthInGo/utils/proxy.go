package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseUrl string, pathPrefix string) http.HandlerFunc {
	target, err := url.Parse(targetBaseUrl)

	if err != nil {
		fmt.Println("Error parsing target URL:", err)
		return nil
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			// 1. Properly set the scheme, host, and base path of the target service
			pr.SetURL(target)

			// 2. Perform path rewriting on the outbound request
			originalPath := pr.Out.URL.Path
			fmt.Println("Original Path :", originalPath)

			strippedPath := strings.TrimPrefix(originalPath, pathPrefix)
			fmt.Println("stripped Path :", strippedPath)

			pr.Out.URL.Path = target.Path + strippedPath

			fmt.Println("Final Path :", pr.Out.URL.Path)


			// 3. Keep the target's Host header (replicates NewSingleHostReverseProxy behavior)
			pr.Out.Host = target.Host
		},
	}

	return proxy.ServeHTTP
}