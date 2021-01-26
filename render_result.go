package sm

import "github.com/golang/geo/s2"

type RenderResult struct {
	Zoom    int
	Center  s2.LatLng
	XOffset int
	YOffset int
}
