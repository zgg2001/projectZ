package data

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

func InitRedis() error {
	err := connectRedis()
	if err != nil {
		return err
	}
	return nil
}

func connectRedis() error {
	log.Println("Connect Redis ...")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		DB:       0,
		PoolSize: 20,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func RedisAddParking(p *parkingRow) error {
	strId := strconv.Itoa(int(p.Id))
	key := ParingInfoPrefix + strId
	fields := map[string]interface{}{
		"id":          p.Id,
		"count":       p.Count,
		"temperature": 0,
		"humidity":    0,
		"weather":     0,
		"address":     p.address,
	}
	err := RedisClient.HMSet(key, fields).Err()
	if err != nil {
		return err
	}
	err = redisAddParkingSpace(strId, p.Count)
	if err != nil {
		return err
	}
	return nil
}

func redisAddParkingSpace(id string, count int32) error {
	keyData := ParingSpaceDataPrefix + id
	keyLicense := ParingSpaceLicensePrefix + id
	fields := map[string]interface{}{}
	for i := int32(0); i < count; i++ {
		fields[strconv.Itoa(int(i))] = ""
	}
	err := RedisClient.HMSet(keyData, fields).Err()
	if err != nil {
		return err
	}
	err = RedisClient.HMSet(keyLicense, fields).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisAddUser(u *userRow) error {
	strId := strconv.Itoa(int(u.Id))
	// user info
	keyInfo := UserInfoPrefix + strId
	fields := map[string]interface{}{
		"id":           u.Id,
		"balance":      u.Balance,
		"username":     u.Username,
		"creationTime": u.CreationTime,
		"lastModified": u.LastModified,
	}
	err := RedisClient.HMSet(keyInfo, fields).Err()
	if err != nil {
		return err
	}
	// user login info
	value := strId + "@" + u.Password
	err = RedisClient.HSet(UserLoginMapKey, u.Username, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisAddLicense(l *licenseRow) error {
	// license info
	keyInfo := LicenseInfoPrefix + l.License
	fields := map[string]interface{}{
		"license":     l.License,
		"pid":         0,
		"psid":        0,
		"checkInTime": l.CheckInTime,
		"entryTime":   0,
	}
	err := RedisClient.HMSet(keyInfo, fields).Err()
	if err != nil {
		return err
	}
	// user license set
	keySet := LicenseSetByUIDPrefix + strconv.Itoa(int(l.Id))
	err = RedisClient.SAdd(keySet, l.License).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisAddRecord(r *recordRow) error {
	// license info
	keyInfo := LicenseInfoPrefix + r.License
	fields := map[string]interface{}{
		"pid":       r.PId,
		"psid":      r.SId,
		"entryTime": r.EntryTime,
	}
	err := RedisClient.HMSet(keyInfo, fields).Err()
	if err != nil {
		return err
	}
	return nil
}
