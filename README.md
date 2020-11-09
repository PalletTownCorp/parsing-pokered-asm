## Parsing-pokered-asm
This cli read a chosen asm fil in pokered project to search specific string or whole without take care of blank line or asm object

1. Git clone project in in $GOPATH/src/github.com folder
```
git clone git@github.com:PalletTownCorp/parsing-pokered-asm.git
```
2. Build cli
```
go build
```
3. Install go binaries
```
go install
```

The cli have two arguments
```
parsing-pokered-asm <filepath> <pattern>
```

The results look like :

```
linenumber: the line which contains pattern
```
