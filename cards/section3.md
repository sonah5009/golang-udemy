# 3-17. Slices and For Loops

Array: Fixed len
Slice: An array that can grow or shrink, every ele must be same type

# 3-18. OO Approach vs Go Approach

객체지향인 Python, Ruby 처럼 class 란 idea가 go 내부엔 없음

# 3-19. Custom Type Declarations

Generating deck.go file and adding print() function

# 3-20. Receiver Functions

convention: usually refers one or two letter of Reveiver

# 3-21. Creating a New Deck

\_ underscore: index of for loop case when we don't need to use index function

# 3-22. Slice Range Syntax

fruits[0:2] : 0-1
fruits[:2] : starts in the begining up to 2 not including 2
fruits[2:] : starts 2 to the end

# 3-23. Multiple Return Values

type이 서로 다른 2개 값을 return 하는 예제 작성

# 3-quiz6

    package main

    import "fmt"

    func main() {
    c := color("Red")

    fmt.Println(c.describe("is an awesome color"))
    }

    type color string

    func (c color) describe(description string) (string) {
    return string(c) + " " + description
    }

string(c) 때문에 error 난다고 착각함

# 3-24. Byte Slices

<del>golang.org</del>  
https://go.dev/  
<del>io -> ioutil: Package ioutil implements some I/O utility finctions</del> decrepted

`func WriteFile(filename string, data []byte, perm os.FileMode) error`  
data []byte: pass argument of byte slice
perm: permission

asciitable.com

# 3-25. Deck to String

write the code

# 3-26. Joining a Slice of Strings

`func Join(a []string, sep string) string`

`func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}`
What I guess:
Although the type of `deck` is as []string, the need for `[]string(d)` is Go's strict type system enforcing explicit type conversion.
Ofc `d` works but the explicit conversion `[]stirng(d)` is still commonly used.

# 3-27. Saving Data to the Hard Drive

<del>ioutil.WriteFile(filename, []byte(d.toString()), 0666)</del>

    import os
    os.WriteFile(filename, []byte(d.toString()), 0666)

# 3-28. Reading From the Hard Drive

add: code in deck.go file

# 3-29. Error Handling

...

# 3-30. Shuffling a Deck

`func (d deck) shuffle()`
In the lecture, doesn't show randomization every single time.
But now it's randomization every singe time.

# 3-31. Random Number Generation

`func NewSource(seed int64) Source`
since Go1.20, Intn() and other functions automatically set the Seed, so the implementation of rand.newSource() and rand.New() mentioned in the current section 31 is unnecessary.

https://tip.golang.org/doc/go1.20#math/rand

# 3-32. Creating a go.mod File

`go mod init cards`

# 3-33. Testing With Go

Add: `deck_test.go`

# 3-34. Errorf call has arguments but no formatting directives

%v: whatever value

# 3-35. Writing Useful Tests

# 3-36. Asserting Elements in a Slice

if statement, check the first and last of cards
