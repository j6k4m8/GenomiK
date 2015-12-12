# Installing Go on a UGRAD Account
1) In a directory of your choice:
```
wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
```
2) Extract the *.tar.gz file to a directory of your choice (assuming `$HOME/go`)
```
tar -C $HOME/go go1.5.2.linux-amd64.tar.gz
```
3) Add relavent paths to your `~/.bashrc` (or `~/.profile`).
```
export GOROOT=$HOME/go
export GOPATH=$HOME/gocode
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT/bin:$GOBIN
```
4) Reload that file.
```
source ~/.bashrc
```

# Cloning our Repo
We _highly recommend_ `go get`ing our repo rather than manually cloning it. To do this, run:
```
go get github.com/j6k4m8/cg/cgg
```
This will clone our repo into `$GOPATH/src/github.com/j6k4m8/cg`.

If you would like to manually place this files (if, for example, you want to test our submission rather than the live github verison...) then make sure that you manually copy the files into the correct path so that `go` can find the correct import paths.

# Compiling `genomik-cli`
While in `$GOPATH/src/github.com/j6k4m8/cg/cgg`, run:
```
go install genomik-cli.go
```
This will compile `genomik-cli` into a statically linked executable located in `$GOBIN` (which you should have previously added to your `$PATH`).
