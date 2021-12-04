package pembayaran

import (
	"strconv"
	"minipro/user"
	midtrans "github.com/veritrans/go-midtrans"
)

type service struct{

}

type Service interface{
	GetPembayaranUrl(transaksi Transaksi, user user.User) (string, error)
}

func NewService() *service{
	return &service{}
}

func (s *service) GetPembayaranUrl(transaksi Transaksi, user user.User) (string, error){
	midclient := midtrans.NewClient()
    midclient.ServerKey = ""
    midclient.ClientKey = ""
    midclient.APIEnvType = midtrans.Sandbox

    
    snapGateway := midtrans.SnapGateway{
        Client: midclient,
    }

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName : user.Nama,
		},
		TransactionDetails: midtrans.TransactionDetails{
            OrderID: strconv.Itoa(transaksi.ID),
            GrossAmt: int64(transaksi.JumlahUang),
        },
	}

	snapTokenResp, err:= snapGateway.GetToken(snapReq)
	if err != nil{
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}