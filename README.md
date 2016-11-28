# b2biRestApiClient
A go client to download a Codelist from b2bi Rest API and store as xls file

it depends on you installing

 `go get github.com/go-resty/resty`  
 `go get github.com/tealeg/xlsx`  

```
 Please provide the following Commandline Parameters when you start the tool
  -codelistname string
          specify codelistname
  -codelistversion string
        specify codelistversion
  -host string
        specify host
  -password string
        specify password for login
  -port string
        specify port
  -username string
        specify username for login
```

An example would look like this:

go run client.go -password password -username admin -host 192.168.2.3 -port 5074 -codelistname myCodelist -codelistversion 1


