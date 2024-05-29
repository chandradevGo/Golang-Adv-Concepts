# Interfaces in Go
### What is an Interface?
An interface in Go is a type that specifies a set of method signatures (without implementing them). It's a way to define what methods a type must have, without saying how those methods should work. If a type has all the methods that an interface requires, it implicitly implements that interface.

### Key Points:
#### Method Set: 
An interface is defined by a set of methods. Any type that has those methods satisfies the interface.
#### Implicit Implementation: 
There's no need to explicitly declare that a type implements an interface. If the methods match, the type satisfies the interface.
#### Polymorphism: 
Interfaces allow you to write functions that can operate on different types as long as those types implement the interface. This is a form of polymorphism.


### Below examples provides an overview of interfaces in Go, using a simple example with two custom types.

### 1. Custom Types
Two structs, person and secretAgent, are defined. The secretAgent struct embeds the person struct, demonstrating inheritance-like behavior.

Two different structs are defined: `person` and `secretAgent`.

```go
type person struct {
    Name string
}

type secretAgent struct {
    person
    ltk bool
}
```


## 2. Methods:

Both structs implement a `speak()` method. This method outputs a string indicating the type of the instance and its name.

## 3. Polymorphism with Interfaces:

An interface `human` is defined, requiring the `speak()` method.
Any type that implements the `speak()` method automatically satisfies the `human` interface.

## 4. Function:

The `saysomething` function takes a `human` interface and calls the `speak()` method on it, demonstrating polymorphism.

## 5. Main Function:

Instances of `person` and `secretAgent` are created.
The `saysomething` function is called with both instances, showcasing that both types satisfy the `human` interface.

You can run this code in a Go environment to see polymorphism in action. The output will show the respective `speak()` method implementation for `person` and `secretAgent`.
