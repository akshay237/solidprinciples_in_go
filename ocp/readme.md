## Open Closed Principle
* The open closed principle states that software entities(classes, modules, functions) should be open for extension but closed for modification.
*  This principle encourages developers to write code that id flrxible and can be extended without much modifications.

### Example
* We have a simple function to calculate the area of rectangle.
```
type Rectangle struct {
    width float64
    height float64
}

func Area(r *Rectangle) float64 {
    return r.width * r.height
}
```

* Now if we need to calculate the area of circle then we have to made modifications.
```
type Circle struct {
    radius float64
}

func Area(shape interface{}) float64 {
    switch s := shape.(type) {
    case *Rectangle:
        return s.width * s.height
    case *Circle:
        return math.Pi * math.Pow(s.Radius, 2)
    default:
        return 0
    }
}
```

* In order to use open/closed principle we need to define a interface and implement it for each shape.
```
type Shape interface {
    Area() float64
}

func (r *Rectange) Area() float64 {
    return r.width * r.height
}

func (c *Circle) Area() float64 {
    return  math.Pi * math.Pow(c.Radius, 2)
}
```

* Now our code is extension for more shapes as well they need to implement the shape interface.

