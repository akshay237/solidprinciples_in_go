## Liskov Substitution Principle (LSP)
* The LSP states that objects of a derived class should be able to replace objects of base class without affecting the correctness of program.
* In golang the principle applies to interfaces and their implementations, ensuring that the code remains consistent and reliable.

### Example
* Consider an example with a simple Bird interface

```
type Bird interface {
    Fly() string
}

type Pigeon struct{}

func (p *Pigeon) Fly() string {
    return "Pigeon is flying".
}

type Penguin struct{}

func (p *Penguin) Fly() string {
    return "Penguin is flying"
}
```

* In this case both Pigeon and Penguin implement the Bird interface.
* However, penguins can not fly so the penguin implementation violates the Liskov Substitution Principle, we have to fix this.

```
type Bird interface {
    MakeSound() string
}

type FlyingBird interface {
    Bird
    Fly() string
}

type Pigeon struct{}

func (p *Pigeon) MakeSound() string {
    return "coo coo"
}

func (p *Pigeon) Fly() string {
    return "Pigeon is flying"
}

type Penguin struct{}

func (p *Penguin) MakeSound() string {
    return "sqwak"
}
```