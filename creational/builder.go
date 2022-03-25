package creational

import "fmt"

// Sub Classes

type camera struct {
	lens lens
	body body
}

type lens struct {
	lensType    string
	focalLength string
	aperture    float32
}

type body struct {
	shutter  string
	sensor   string
	material string
}

func (c camera) Briefing() {
	fmt.Println(c.Assemble())
}

func (c camera) Assemble() string {
	return "Assmbled Camera:\n" + fmt.Sprintf("Lens:\n\t%s\n\t%s\n\t%.1f\n", c.lens.lensType, c.lens.focalLength, c.lens.aperture) + fmt.Sprintf("Body:\n\t%s\n\t%s\n\t%s\n", c.body.shutter, c.body.sensor, c.body.material)
}

type Camera interface {
	Briefing()
	Assemble() string
}

// Builders

type builder struct {
	cam camera
}

type Builder interface {
	AddLens(lensType, focalLength string, aperture float32)
	AddBody(shutter, sensor, material string)
	GetProduct() Camera
}

func (b *builder) AddLens(lensType, focalLength string, aperture float32) {
	b.cam.lens = lens{lensType: lensType, focalLength: focalLength, aperture: aperture}
}

func (b *builder) AddBody(shutter, sensor, material string) {
	b.cam.body = body{shutter: shutter, sensor: sensor, material: material}
}

func (b *builder) GetProduct() Camera {
	return b.cam
}

func GetBuilder() Builder {
	cam := camera{}
	return &builder{cam: cam}
}
