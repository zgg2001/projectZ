package data

import (
	"fmt"
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
		fmt.Println(tempParking)
		err = RedisAddParking(tempParking)
		if err != nil {
			return err
		}
	}

	return nil
}
