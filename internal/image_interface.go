package internal

import "image/color"

type ImageInterface interface {
	Fill(color color.Color)
}
