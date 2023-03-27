package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/parking_server/pkg/rpc"
)

func (ss *serverService) UploadParkingInfo(con context.Context, request *rpc.UploadInfoRequest) (*rpc.UploadInfoResponse, error) {

	pInfo := request.PInfo
	sInfo := request.SInfoArr
	pid := pInfo.GetPId()

	ss.UpdateParkingData(pid, pInfo.GetTemperature(), pInfo.GetHumidity(), pInfo.GetWeather())
	for _, i := range sInfo {
		ss.UpdateParkingSpaceData(pid, i.SId, i.GetTemperature(), i.GetHumidity(), int32(i.GetAlarm()))
	}

	return &rpc.UploadInfoResponse{Result: 1}, nil
}
