# Testing nginx ingress controller reload


## Usage

Install all the applications in your cluster, and get a connection to your ingress controller

```bash
kubectl apply -f install/k8s
helm install stable/nginx-ingress --version 1.21.0 --values install/nginx/values.yaml --name nginx-ingress
(while true; do kubectl port-forward $(k get po | grep nginx-ingress-controller | head | awk '{print $1}') 8080:32080; done) &
while true; do kubectl logs -f $(k get po | grep nginx-ingress-controller | head | awk '{print $1}'); done
```

Then in a different terminal, run the test: `go run main.go`

you should observe:

```
connecting to proxy, original address: http-echo.example.com:80
ingress.extensions/http-echo annotated
Post http://http-echo.example.com/1571994168?count=9: EOF
it looks like we reproduced the problem, exiting
exit status 1
```

And the logs looks like
```
{ "time": "2019-10-25T09:02:48+00:00", "ingress_class": "legacynginx", "remote_addr": "192.168.1.3","x-forward-for": "192.168.1.3", "request_id": "6ff509487a7fcf1e708aed5587df5f95", "remote_user":"", "bytes_sent": 657, "request_time": 0.102, "status":200, "vhost": "http-echo.example.com", "request_proto": "HTTP/1.1", "path": "/1571994168","request_query": "count=0", "request_length": 209, "duration": 0.102,"method": "POST", "http_referrer": "", "http_user_agent":"Go-http-client/1.1","namespace":"default","ingress_name":"http-echo","service_name":"http-echo","geoip":{"country_code": "","country_name": "","city_name": "","region_name": "","region_code": "","location": ","}}
{ "time": "2019-10-25T09:02:48+00:00", "ingress_class": "legacynginx", "remote_addr": "192.168.1.3","x-forward-for": "192.168.1.3", "request_id": "5a0ddc61cd75fa3750700e44ed62fb59", "remote_user":"", "bytes_sent": 657, "request_time": 0.107, "status":200, "vhost": "http-echo.example.com", "request_proto": "HTTP/1.1", "path": "/1571994168","request_query": "count=1", "request_length": 209, "duration": 0.107,"method": "POST", "http_referrer": "", "http_user_agent":"Go-http-client/1.1","namespace":"default","ingress_name":"http-echo","service_name":"http-echo","geoip":{"country_code": "","country_name": "","city_name": "","region_name": "","region_code": "","location": ","}}
{ "time": "2019-10-25T09:02:49+00:00", "ingress_class": "legacynginx", "remote_addr": "192.168.1.3","x-forward-for": "192.168.1.3", "request_id": "99b20c0f344a415337e04a6f3ccedc05", "remote_user":"", "bytes_sent": 658, "request_time": 0.110, "status":200, "vhost": "http-echo.example.com", "request_proto": "HTTP/1.1", "path": "/1571994168","request_query": "count=2", "request_length": 209, "duration": 0.110,"method": "POST", "http_referrer": "", "http_user_agent":"Go-http-client/1.1","namespace":"default","ingress_name":"http-echo","service_name":"http-echo","geoip":{"country_code": "","country_name": "","city_name": "","region_name": "","region_code": "","location": ","}}
{ "time": "2019-10-25T09:02:49+00:00", "ingress_class": "legacynginx", "remote_addr": "192.168.1.3","x-forward-for": "192.168.1.3", "request_id": "5d348b436e5695b60d53fe8dffd4f582", "remote_user":"", "bytes_sent": 659, "request_time": 0.108, "status":200, "vhost": "http-echo.example.com", "request_proto": "HTTP/1.1", "path": "/1571994168","request_query": "count=3", "request_length": 209, "duration": 0.108,"method": "POST", "http_referrer": "", "http_user_agent":"Go-http-client/1.1","namespace":"default","ingress_name":"http-echo","service_name":"http-echo","geoip":{"country_code": "","country_name": "","city_name": "","region_name": "","region_code": "","location": ","}}
{ "time": "2019-10-25T09:02:49+00:00", "ingress_class": "legacynginx", "remote_addr": "192.168.1.3","x-forward-for": "192.168.1.3", "request_id": "17f6c24a3e9f1b293f549af8cf30c71d", "remote_user":"", "bytes_sent": 657, "request_time": 0.104, "status":200, "vhost": "http-echo.example.com", "request_proto": "HTTP/1.1", "path": "/1571994168","request_query": "count=4", "request_length": 209, "duration": 0.104,"method": "POST", "http_referrer": "", "http_user_agent":"Go-http-client/1.1","namespace":"default","ingress_name":"http-echo","service_name":"http-echo","geoip":{"country_code": "","country_name": "","city_name": "","region_name": "","region_code": "","location": ","}}
{ "time": "2019-10-25T09:02:49+00:00", "ingress_class": "legacynginx", "remote_addr": "192.168.1.3","x-forward-for": "192.168.1.3", "request_id": "85ce914712a6413a63da3bfc40a07e65", "remote_user":"", "bytes_sent": 657, "request_time": 0.108, "status":200, "vhost": "http-echo.example.com", "request_proto": "HTTP/1.1", "path": "/1571994168","request_query": "count=5", "request_length": 209, "duration": 0.108,"method": "POST", "http_referrer": "", "http_user_agent":"Go-http-client/1.1","namespace":"default","ingress_name":"http-echo","service_name":"http-echo","geoip":{"country_code": "","country_name": "","city_name": "","region_name": "","region_code": "","location": ","}}
I1025 09:02:49.767070       6 server.go:61] handling admission controller request /extensions/v1beta1/ingresses?timeout=30s
{ "time": "2019-10-25T09:02:49+00:00", "ingress_class": "legacynginx", "remote_addr": "192.168.1.3","x-forward-for": "192.168.1.3", "request_id": "c8288e8b9a065cc6a039a1aaf0156a44", "remote_user":"", "bytes_sent": 661, "request_time": 0.108, "status":200, "vhost": "http-echo.example.com", "request_proto": "HTTP/1.1", "path": "/1571994168","request_query": "count=6", "request_length": 209, "duration": 0.108,"method": "POST", "http_referrer": "", "http_user_agent":"Go-http-client/1.1","namespace":"default","ingress_name":"http-echo","service_name":"http-echo","geoip":{"country_code": "","country_name": "","city_name": "","region_name": "","region_code": "","location": ","}}
I1025 09:02:49.796334       6 main.go:86] successfully validated configuration, accepting ingress http-echo in namespace default
I1025 09:02:49.801898       6 event.go:255] Event(v1.ObjectReference{Kind:"Ingress", Namespace:"default", Name:"http-echo", UID:"4d509f2c-f649-11e9-b208-025000000001", APIVersion:"networking.k8s.io/v1beta1", ResourceVersion:"108901", FieldPath:""}): type: 'Normal' reason: 'UPDATE' Ingress default/http-echo
I1025 09:02:49.802249       6 controller.go:134] Configuration changes detected, backend reload required.
I1025 09:02:49.903885       6 controller.go:150] Backend successfully reloaded.
{ "time": "2019-10-25T09:02:49+00:00", "ingress_class": "legacynginx", "remote_addr": "192.168.1.3","x-forward-for": "192.168.1.3", "request_id": "81f80afe48188b60c3342d91b1a3862a", "remote_user":"", "bytes_sent": 658, "request_time": 0.144, "status":200, "vhost": "http-echo.example.com", "request_proto": "HTTP/1.1", "path": "/1571994168","request_query": "count=7", "request_length": 209, "duration": 0.144,"method": "POST", "http_referrer": "", "http_user_agent":"Go-http-client/1.1","namespace":"default","ingress_name":"http-echo","service_name":"http-echo","geoip":{"country_code": "","country_name": "","city_name": "","region_name": "","region_code": "","location": ","}}
{ "time": "2019-10-25T09:02:50+00:00", "ingress_class": "legacynginx", "remote_addr": "192.168.1.3","x-forward-for": "192.168.1.3", "request_id": "c9dfeb414aa4b13877da4b8e6255a0dc", "remote_user":"", "bytes_sent": 657, "request_time": 0.069, "status":200, "vhost": "http-echo.example.com", "request_proto": "HTTP/1.1", "path": "/1571994168","request_query": "count=8", "request_length": 209, "duration": 0.069,"method": "POST", "http_referrer": "", "http_user_agent":"Go-http-client/1.1","namespace":"default","ingress_name":"http-echo","service_name":"http-echo","geoip":{"country_code": "","country_name": "","city_name": "","region_name": "","region_code": "","location": ","}}
```

Note that request number 9 is not logged

