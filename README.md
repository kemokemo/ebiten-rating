# ebiten-rating

This package show the rating value with your specified image.

![sample image](media/sample.png)

## Usage

```go
// Initialize with your image
g.rating = rating.NewRating(StarImage, 15, 20, 10)

// Set the current rating value
g.rating.SetValue(5.7)

// Draw
func (g *Game) Draw(screen *ebiten.Image) {
	g.rating.Draw(screen)
}
```

## License

Apache-2.0 License

## Auther

kemokemo
