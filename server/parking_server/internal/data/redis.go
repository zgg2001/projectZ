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

func RedisAddParking(p *ParkingRow) error {
	strId := strconv.Itoa(int(p.Id))
	key := ParkingInfoPrefix + strId
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
	// parking login info
	value := strconv.Itoa(int(p.Count)) + "@" + p.Password
	err = RedisClient.HSet(ParkingLoginMapKey, strconv.Itoa(int(p.Id)), value).Err()
	if err != nil {
		return err
	}
	return nil
}

func redisAddParkingSpace(id string, count int32) error {
	keyData := ParkingSpaceDataPrefix + id
	keyLicense := ParkingSpaceLicensePrefix + id
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

func RedisAddUser(u *UserRow) error {
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

func RedisAddLicense(l *LicenseRow) error {
	// license info
	keyInfo := LicenseInfoPrefix + l.License
	fields := map[string]interface{}{
		"uid":         l.Id,
		"license":     l.License,
		"pid":         -1,
		"psid":        -1,
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

func RedisAddRecord(r *RecordRow) error {
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

func RedisSetBalance(uid string, balance int32) error {
	keyInfo := UserInfoPrefix + uid
	err := RedisClient.HSet(keyInfo, "balance", balance).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisSetParkingInfo(pid, temperature, humidity, weather int32) error {
	strId := strconv.Itoa(int(pid))
	key := ParkingInfoPrefix + strId
	fields := map[string]interface{}{
		"temperature": temperature,
		"humidity":    humidity,
		"weather":     weather,
	}
	err := RedisClient.HMSet(key, fields).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisSetParkingSpaceInfo(pid, sid int32, data string) error {
	strpid := strconv.Itoa(int(pid))
	strsid := strconv.Itoa(int(sid))
	key := ParkingSpaceDataPrefix + strpid
	err := RedisClient.HSet(key, strsid, data).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisGetUidByLicense(license string) (ok bool, uid string) {
	key := LicenseInfoPrefix + license
	uid, err := RedisClient.HGet(key, "uid").Result()
	if err != nil || uid == "" {
		ok = false
	}
	return
}

func RedisGetBalanceByUid(uid string) (bool, int32) {
	ok := true
	key := UserInfoPrefix + uid
	balance, err := RedisClient.HGet(key, "balance").Int()
	if err != nil || balance == 0 {
		ok = false
	}
	return ok, int32(balance)
}

func RedisCheckCarIsEntered(license string) bool {
	key := LicenseInfoPrefix + license
	pid, err1 := RedisClient.HGet(key, "pid").Int()
	sid, err2 := RedisClient.HGet(key, "psid").Int()
	entryTime, err3 := RedisClient.HGet(key, "entryTime").Int64()
	if err1 != nil || err2 != nil || err3 != nil {
		return false
	}
	if pid >= 0 && sid >= 0 && entryTime > 0 {
		return true
	}
	return false
}
