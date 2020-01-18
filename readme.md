# Shape

A library for drawing common shapes like: line, circle, rectangle and more.

## Line
```
shape.Line{a,b,...}.Stroke(width, gtx)
```

## Circle
```
shape.Circle{c,r}.Stroke(width, gtx)
shape.Circle{c,r}.Fill(gtx)
```

## Rectangle
```
shape.Rectangle{a,b}.Stroke(width, gtx)
shape.Rectangle{c,r}.Fill(gtx)
```

## Triangle
```
shape.Triangle{c,r}.Stroke(width, gtx)
shape.Triangle{c,r}.Fill(gtx)
```

## Curve
```
shape.Curve{a,b,...}.Stroke(width,gtx)
```

## Example Draw Red Line 
```go
func drawLine(gtx *layout.Context) {
    var stack op.StackOp
    stack.Push(gtx.Ops)
    line := shape.Line{{0, merginTop}, {lineLen, merginTop}}
    box := line.Stroke(width, shape.Solid, gtx)
    paint.ColorOp{color.RGBA{255,0,0,255}}.Add(gtx.Ops)
    paint.PaintOp{box}.Add(gtx.Ops)
    stack.Pop()
}
```

See in example folder for gio application with more examples of shapes.