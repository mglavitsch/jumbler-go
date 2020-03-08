# jumbler-go

The golang package `jumbler-go` provides functions for jumbling letters within words and sentences.

```
go get github.com/mglavitsch/jumbler-go
```

# Try it out

```
$ git clone https://github.com/mglavitsch/jumbler-go.git
$ cd jumbler-go
$ go test
// see result of unit tests
$ cd dojumble
$ go build
$ ./dojumble -text "jumbling"
=> see result of jumbling the word "jumbling"
$ ./dojumble
=> see result of jumbling the text of the hypothesis
```

# Background

## Hypothesis

It doesn't matter in what order the letters in a word are, the only important thing is that the first and last letter be at the right place. The rest can be a total mess and you can still read it without problem. This is because the human mind does not read every letter by itself, but the word as a whole.

## Example

The jumbling of the words of the hypothesis stated above might end up in the following result:

"It doesn't matetr in what order the letrets in a word are, the only imonarptt thing is that the first and last letetr be at the right place. The rest can be a total mess and you can still read it wihuott prbelom. This is beasuce the human mind does not read every letetr by itslef, but the word as a whole."

## Reference

This statement goes back to Graham Ernest Rawlinson: "The Significance of Letter Position in Word Recognition", PhD Thesis, 1976, University of Nottingham
