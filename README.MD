in order to package (build) go GUI app without cli showing up, use this flag:
```
go build -ldflags -H=windowsgui .\HelloWorld.go
```
or use:
```
fyne package -icon .\Resources\RSLicon.png
```

the make sure we've got all of the translation, use:
```
go generate .\translations\translations.go
```

learning materials:
https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1

https://github.com/fyne-io/fyne

https://developer.fyne.io/api/v2.1/