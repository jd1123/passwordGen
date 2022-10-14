# Random password generator in Go

## Why?
I like random passwords for sensitive logins since password attacks are pretty easy to execute. This generates a truly pseudorandom password (based on Go's crypto/rand package). This can be used effectively with the [pass](https://wiki.archlinux.org/index.php/Pass) package in Archlinux.

Looking back on this five years later, I still find this small piece of software very useful. I hope you do too!

## Instructions
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

## Password Policies
Password policies are used to set different number of letters, numbers and characters to support whatever the password requires. You can also set the entropy of each to allow for variable length passwords. Entropy simply adds between 0 and e runes of a particular type where e is the entropy value. The standard policy is a fixed length, 14 rune password with 7 letters, 4 numbers and 3 special characters. You can change this with command line flags. Here are two examples:

The following generates a password from 7 to 14 runes, with a minimum of 5 letters and 2 numbers:
```
passwordGen -ml 5 -mn 2 -mc 0 -le 4 -ne 3 -ce 0
```

## Config

On first use, passwordGen creates a file called .pginfo in your home directory with a json object outlining the policy parameters above. A default policy is written to this file. passwordGen will look for this file upon invocation. You can change the paremeters to suit your needs.


Have fun!
