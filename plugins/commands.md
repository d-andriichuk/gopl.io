# GO commands for plugins compiling

For compiling each file with concrete name.

```sh
go build -buildmode=plugin -o [pluginFolder]/plufinName.so [packageFolder]/packageName.go
```

For compiling all package into one plugin .so file. This comman set plugin name like package name.

```sh
go build -buildmode=plugin
```