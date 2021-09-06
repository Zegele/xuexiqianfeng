管理外部包：
go 允许import不同代码库的代码。
对于import要导入的外部包，可以使用go get命令取下来放到GOPATH对应的目录中去。

go get github.com/go-sql-driver/mysql
// 就会在GOPATH的src目录下，出现一个github文件夹，继续有go-sql-driver文件夹，继续有mysql
// 如果在本ide（编辑软件）中发现自动导入的包不对，就用alt+enter，手动选择要选中的包。

编译包文件：
可以通过go install 编译包文件。
一个非main包在编译后会生成一个.a 的文件（在临时目录下生成，除非使用go install安装到$GOROOT或$GOPATH下，否则看不到.a），用于后续可执行程序链接使用。
比如Go标准库中的包对应的源码部分路径在：$GOROOT/src，而标准库中包编译后的.a文件路径在$GOROOT/pkg/darwin_amd64下。


