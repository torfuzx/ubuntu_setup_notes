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
    echo "export PATH=$PATH:/usr/local/go/bin/:$GOPATH/bin" >> ~/.bashrc
    # flush the change
    source ~/.bashrc
    ```

Go Tools
--------


1. [goimports](https://github.com/bradfitz/goimports)
    
    ```
    go get code.google.com/p/go.tools/cmd/goimports
    ```

2. [gocode](github.com/nsf/gocode)

    ```
    go get -u github.com/nsf/gocode
    ```
    
3. [golint](github.com/golang/lint)

    ```
    go get github.com/golang/lint/golint
    ```
4. [gooracle](https://docs.google.com/document/d/1SLk36YRjjMgKqe490mSRzOPYEDe0Y_WQNRv-EiFYUyw/view?pli=1)

    ```
    go get code.google.com/p/go.tools/cmd/oracle
    ```
