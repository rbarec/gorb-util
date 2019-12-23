#  Layout golang CMD folder

## Package

- [] package logger, extended datadog logrus
- [] package str

## getting started
https://github.com/golang-standards/project-layout/blob/master/cmd/README.md

https://github.com/heptio/ark/tree/master/cmd 
velero/cmd/velero/main.go/
package main


## package logger
Wrapper de logrus v 0.0.1
 
## hacer un tag por module
git tag v1.0-alpha.1
git push origin v1.0-alpha.1

## En otros proyectos deben invocarse al paquete 
go get github.com/rbarec/gorb-util/str
>OK

go doc github.com/rbarec/gorb-util/str
>package str // import "github.com/rbarec/gorb-util/str"
>func GetInfo() string 	