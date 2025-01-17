## Interface Segregation Principle (ISP)
* The Interface Segregation Principle states that client should not be forced to depend on interfaces they do not use.
* This principle encourages creating smaller, more focused interfaces rather than large ones.

### Example
* Suppose we have a Document interface with methods for reading, writing and printing.

```
type Document interface {
    Read() string
    Write(content string)
    Print() string
}

type TextDocument struct {
    content string
}

func(d *TextDocument) Read() string {
    return d.content
}

func (d *TextDocument) Write(content string) {
    d.content = content
}

func (d *Document) Print() {
    return "Printing content" + d.content
}
```

* Suppose we have a read only document in that case that won't need to implement other than read method of the interface.

```
type ReadOnlyDocument struct {
    content string
}

func(d *ReadOnlyDocument) Read() string {
    return d.content
}

func(d *ReadOnlyDocument) Write(content string) {
    // Not Applicable
}

func(d *ReadOnlyDocument) Print() {
    // Not Applicable
}
```

* To follow the ISP we have to divide the interface into multiple interfaces as some methods of interface are not applicable.

```
type Reader interface {
    Read() string
}

type Writer interface {
    Write(content string)
}

type Printer interface {
    Print() string
}

type TextDocument struct {
    content string
    // implements Reader, Writer & Printer interface
}

type ReadOnlyDocument struct {
    content string
    // implements Reader interface
}
```
