package pi

type ParkingOperation interface {
	Reset(id int)
	DriveInto(license string)
	DriveOut()
}

type Parking struct {
	Id      int
	IsUsing bool
	License string
}

func (p *Parking) Reset(id int) {
	p.Id = id
	p.IsUsing = true
	p.License = ""
}

func (p *Parking) DriveInto(license string) {
	p.IsUsing = true
	p.License = license
}

func (p *Parking) DriveOut() {
	p.IsUsing = false
	p.License = ""
}
