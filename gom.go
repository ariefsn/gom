package gom

// Gom struct
type Gom struct {
	Mongo Mongo
}

// NewGom = Create new
func NewGom() *Gom {
	return new(Gom)
}

// Init = Init
func (d *Gom) Init(config MongoConfig) {
	d.Mongo.SetConfig(config)
	d.Mongo.SetContext(30)
	d.Mongo.SetClient()
}

// Data = Get data with gom
func (d *Gom) Data() *Set {
	s := NewSet(d)

	return s
}
