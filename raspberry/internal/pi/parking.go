package pi

type status = int

var (
	EmptyParkingSpace    int = 0
	NonEmptyParkingSpace int = 1
)

type ParkingOperation interface {
	Reset(id int)
	GetStatus() status
	DriveInto(license string)
	DriveOut()
}

type Parking struct {
	Id      int
	IsUsing bool
	License string
}

func (p *Parking) GetStatus() status {
	if p.IsUsing {
		return NonEmptyParkingSpace
	}
	return EmptyParkingSpace
}

func (p *Parking) Reset(id int) {
	p.Id = id
	p.IsUsing = false
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
