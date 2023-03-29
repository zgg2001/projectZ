package data

import (
	"log"
)

func ParkingInit() error {

	log.Println("Parking data load ...")

	// read and load parking
	parkingRet, err := ReadParkingTbl()
	if err != nil {
		return err
	}
	for _, tempParking := range parkingRet {
		err = RedisAddParking(tempParking)
		if err != nil {
			return err
		}
	}

	return nil
}
