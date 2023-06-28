# Go annotations

## general commands

go mod tidy is used to remove unused dependencies from the go.mod file

## access modifiers

so, apparently Go doesn't have access modifiers because it isn't an object oriented programming language

what is used to denote an access modifier is: the method usually begins with a capital letter (see packages/helper/helper.go)
when the function is public, and with a lower case letter when the method can be accessed only at package level

i found that interesting that's why i'm writing this down

## variable and constants

you can use the short := to declare a variable, making use of the type inference provided by go
on the other hand, you cannot do this when declaring a constant
when doing so, you must declare it explicitly, eg.:

const Constant string = "i'm a constant"