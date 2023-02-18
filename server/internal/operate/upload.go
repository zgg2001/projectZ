package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) UploadParkingInfo(con context.Context, request *rpc.UploadInfoRequest) (*rpc.UploadInfoResponse, error) {

	pInfo := request.PInfo
	sInfo := request.SInfoArr

	pPtr, err := ss.pMgr.MgrGetParkingPtr(pInfo.PId)
	if err != nil {
		return &rpc.UploadInfoResponse{Result: 0}, err
	}
	pPtr.UpdateParkingData(pInfo.GetTemperature(), pInfo.GetHumidity(), pInfo.GetWeather())

	for _, i := range sInfo {
		sPtr, err := pPtr.GetParkingPtr(i.SId)
		if err != nil {
			return &rpc.UploadInfoResponse{Result: 0}, err
		}
		sPtr.UpdateParkingSpaceData(i.GetTemperature(), i.GetHumidity(), i.GetAlarm())
	}

	return &rpc.UploadInfoResponse{Result: 1}, nil
}
