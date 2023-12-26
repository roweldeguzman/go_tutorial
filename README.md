## To  run in development
`nodemon`

**needs to install nodemon via `npm or yarn` as global package.

## Or you can use
`https://github.com/gravityblast/fresh` for autorefresh development

## To run migration
to  run migration: go to migrations folder and run `go run *.go`

<br />
<br />
<br />

# Add custom development location (OS: linux/ubuntu)

#### declare custom location: 
```export D=/mnt/D/projects/go```
#### set your default home
```export GOROOT=/home/{USER}/.go```
#### set default path
```export PATH=$GOROOT/bin:$PATH```
#### set custom gopath
```export GOPATH=$D```
#### add custom path for go development
```export PATH=$GOPATH/bin:$PATH```