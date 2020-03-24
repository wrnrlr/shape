# Shape

A library for drawing common shapes like: line, circle, rectangle and more.

## Line
```
shape.Line{a,b,...}.Stroke(color, width, gtx)
```

## Circle
```
shape.Circle{c,r}.Stroke(color, width, gtx)
shape.Circle{c,r}.Fill(color, gtx)
```

## Rectangle
```
shape.Rectangle{a,b}.Stroke(color, width, gtx)
shape.Rectangle{a,b}.Fill(color, gtx)
```

## Triangle
```
shape.Triangle{a,b,c}.Stroke(color, width, gtx)
shape.Triangle{a,b,c}.Fill(color, gtx)
```

## Curve
```
shape.Curve{a,b,...}.Stroke(color, width,gtx)
```

## Points
```
...
```

## Example Draw Red Line 
```go
func drawLine(gtx *layout.Context) {
    red := color.RGBA{255,0,0,255}
    s := gtx.Constraints
    w, h := float32(s.Width.Max), float32(s.Height.Max)
    shape.Line{{0,h/2},{w,h/2}}.Stroke(red, width, gtx)
}
```

See in example folder for gio application with more examples of shapes.

## Read More

* [Drawing Antialiased Lines with OpenGL](https://blog.mapbox.com/drawing-antialiased-lines-with-opengl-8766f34192dc)