# http://localhost:8080/index

# Change the server port by defining it on deployment machine as OS env property:  
    On linux:
    export GO_SERVICE_PORT=9090 
    
    On windows:
    set GO_SERVICE_PORT=9090
    
    
# go build || go run

# set the GOPATH on windows
    in cmd run:
    go env -w GOPATH=D:\goLang\workspace
    
    https://github.com/golang/go/wiki/GOPATH
    
        
# testing:
     go get github.com/stretchr/testify/assert

# run the test from terminal in the root dir:      
     go test ./api -v


# Concurrency patterns:
   https://www.youtube.com/watch?v=f6kdp27TYZs
   
# MongoDB:
    $ go get gopkg.in/mgo.v2
    $ go get github.com/gorilla/mux

    https://gopkg.in/mgo.v2
    http://niemeyer.net/mgo
   
# Module versions management / Dependency management:
 # Version naming conventions:
 Everything starts with v1.0.0
    - Major version 1
    - Minor version 0 
    - Patch version 0
 
 Minor version is not backward compatible   
 Minor version is for implementing features and is backward compatible 
 Patch for fixing bugs (security issues ) / issues and is backward compatible