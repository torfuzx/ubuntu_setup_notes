Installing go 
-------------

1. Find the disired releae at the [golang release page](https://golang.org/dl/) and download it

    ```
    cd ~/Downloads
    wget https://storage.googleapis.com/golang/go1.3.3.linux-amd64.tar.gz
    ```

2. Extract it and move it to the right place

    ```
    tar xzvf go1.3.3.linux-amd64.tar.gz
    mv go /usr/local/go
    ```
3. Set the path 

    ```
    echo "export GOPATH='~/gopath'" >> ~/.bashrc
    echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
    # flush the change
    source ~/.bashrc
    ```

Go Tools
--------


### [goimports](https://github.com/bradfitz/goimports)

    go get code.google.com/p/go.tools/cmd/goimports

### [gocode](github.com/nsf/gocode)

    go get -u github.com/nsf/gocode
  
###[golint](github.com/golang/lint)

    go get github.com/golang/lint/golint
