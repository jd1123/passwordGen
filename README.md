# Random password generator in Go

# Instructions
Download the package and build it:
```
go get github.com/jd1123/passwordGen
cd $GOPATH/src/github.com/jd1123/passwordGen
go build
```

Move the built binary to somewhere useful:
```
cp passwordGen /usr/local/bin/
```

Generate passwords!
```
passwordGen
```

Password policies have a minimum number of different types of characters. You can edit this in:
```
min_letters := 6
min_numbers := 4
min_chars := 2
```
You can also edit the entropy of each type of rune:
```
letter_ent := 3
number_ent := 6
char_ent := 3
```

Have fun!
