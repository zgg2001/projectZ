package data

import (
	"log"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
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

func RedisCheckUserIsExists(username string) bool {
	exist, _ := RedisClient.HExists(UserLoginMapKey, username).Result()
	return exist
}

func RedisCheckUserIsExistsByUid(uid int32) bool {
	strId := strconv.Itoa(int(uid))
	key := UserInfoPrefix + strId
	count, _ := RedisClient.Exists(key).Result()
	return count == 1
}

func RedisCheckLicenseIsExists(license string) bool {
	key := LicenseInfoPrefix + license
	count, _ := RedisClient.Exists(key).Result()
	return count == 1
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

func RedisDelLicense(uid int32, license string) error {
	// license info
	keyInfo := LicenseInfoPrefix + license
	err := RedisClient.Del(keyInfo).Err()
	if err != nil {
		return err
	}
	// user license set
	keySet := LicenseSetByUIDPrefix + strconv.Itoa(int(uid))
	err = RedisClient.SRem(keySet, license).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisUpdLicense(uid int32, oldLicense, newLicense string, checkInTime int64) error {
	// delete old license info
	keyInfo := LicenseInfoPrefix + oldLicense
	err := RedisClient.Del(keyInfo).Err()
	if err != nil {
		return err
	}
	// add new license info
	err = RedisAddLicense(&LicenseRow{
		Id:          uid,
		License:     newLicense,
		CheckInTime: checkInTime,
	})
	if err != nil {
		return err
	}
	// remove old license from user license set
	keySet := LicenseSetByUIDPrefix + strconv.Itoa(int(uid))
	err = RedisClient.SRem(keySet, oldLicense).Err()
	if err != nil {
		return err
	}
	// add new license to user license set
	err = RedisClient.SAdd(keySet, newLicense).Err()
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

func RedisGetBalanceByUid(uid string) (bool, int32) {
	ok := true
	key := UserInfoPrefix + uid
	balance, err := RedisClient.HGet(key, "balance").Int()
	if err != nil || balance == 0 {
		ok = false
	}
	return ok, int32(balance)
}

func RedisGetPasswordByUsername(username string) (bool, int32, string) {
	value, err := RedisClient.HGet(UserLoginMapKey, username).Result()
	if err != nil || len(value) == 0 {
		return false, -1, ""
	}
	parts := strings.SplitN(value, "@", 2)
	if len(parts) < 2 {
		return false, -1, ""
	}
	uid, _ := strconv.Atoi(parts[0])
	password := parts[1]
	return true, int32(uid), password
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

func RedisGetLicensesByUID(uid int32) ([]string, error) {
	keySet := LicenseSetByUIDPrefix + strconv.Itoa(int(uid))
	licenses, err := RedisClient.SMembers(keySet).Result()
	if err != nil {
		return nil, err
	}
	return licenses, nil
}

func RedisGetLicenseInfo(license string) *rpc.CarInfo {
	keyInfo := LicenseInfoPrefix + license
	result, err := RedisClient.HGetAll(keyInfo).Result()
	if err != nil || len(result) == 0 {
		return nil
	}
	sid, _ := strconv.Atoi(result["psid"])
	keyParking := ParkingInfoPrefix + result["pid"]
	pTemperature, _ := RedisClient.HGet(keyParking, "temperature").Int()
	pHumidity, _ := RedisClient.HGet(keyParking, "humidity").Int()
	pWeather, _ := RedisClient.HGet(keyParking, "weather").Int()
	pAddress, _ := RedisClient.HGet(keyParking, "address").Result()
	sdata, _ := RedisClient.HGet(ParkingSpaceDataPrefix+result["pid"], result["psid"]).Result()
	return &rpc.CarInfo{
		PTemperature: int32(pTemperature),
		PHumidity:    int32(pHumidity),
		PWeather:     int32(pWeather),
		PAddress:     pAddress,
		SId:          int32(sid),
		SData:        sdata,
	}
}

func RedisGetParkingPasswordById(id int32) (bool, int32, string) {
	value, err := RedisClient.HGet(ParkingLoginMapKey, strconv.Itoa(int(id))).Result()
	if err != nil || len(value) == 0 {
		return false, -1, ""
	}
	parts := strings.SplitN(value, "@", 2)
	if len(parts) < 2 {
		return false, -1, ""
	}
	count, _ := strconv.Atoi(parts[0])
	password := parts[1]
	return true, int32(count), password
}
