# Tap

Tap is a simple CLI program to organize files

## Installing

Installing via Go

```bash
$ go get -u github.com/boomsquared/tap
$ cd $GOPATH/src/github.com/boomsquared/tap
$ go install .
```

## Usage

### Group files(mainly images) into folders base on the by flag
```
tap group [folder] --by [device,extension,fnumber,iso,lens,shutterspeed]
```

### Rename files base on the by flag (Developing)
```
tap rename [folder] --by [size] --prefix[whatever]
```


## Remarks

This is created with the intent to learn, you will encounter bugs and insects! \n
Any suggestion of how could I improve is certainly welcome.


## Contributions

Feel free :D






