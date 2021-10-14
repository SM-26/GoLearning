in order to package (build) go GUI app without cli showing up, use this flag:
go build -ldflags -H=windowsgui .\HelloWorld.go

learning materials:
https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1

https://github.com/fyne-io/fyne