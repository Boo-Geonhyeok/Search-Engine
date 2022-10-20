package models

type BookList struct {
	Data struct {
		TotalCount  int `json:"totalCount"`
		TabContents []struct {
			SaleCmdtID          string        `json:"saleCmdtId"`
			CmdtName            string        `json:"cmdtName"`
			Cmdtcode            string        `json:"cmdtcode"`
			SaleCmdtGrpDvsnCode string        `json:"saleCmdtGrpDvsnCode"`
			SaleCmdtDvsnCode    string        `json:"saleCmdtDvsnCode"`
			SaleCmdtClstCode    string        `json:"saleCmdtClstCode"`
			PbcmName            string        `json:"pbcmName"`
			RlseDate            string        `json:"rlseDate"`
			InbukCntt           string        `json:"inbukCntt"`
			Price               float32       `json:"price"`
			Sapr                float32       `json:"sapr"`
			DscnRate            float32       `json:"dscnRate"`
			UpntAcmlAmnt        float32       `json:"upntAcmlAmnt"`
			AcmlRate            float32       `json:"acmlRate"`
			SaleLmttAge         float32       `json:"saleLmttAge"`
			Preview             string        `json:"preview"`
			RevwCont            float32       `json:"revwCont"`
			RevwRvgr            float32       `json:"revwRvgr"`
			FeelTag             interface{}   `json:"feelTag"`
			DlvrRqrmDyCont      float32       `json:"dlvrRqrmDyCont"`
			DlvrDesc            interface{}   `json:"dlvrDesc"`
			ChrcName            string        `json:"chrcName"`
			Wish                string        `json:"wish"`
			Best                interface{}   `json:"best"`
			Review              []interface{} `json:"review"`
			Casting             []interface{} `json:"casting"`
			DlvrCode            interface{}   `json:"dlvrCode"`
			SaleCdtnCode        string        `json:"saleCdtnCode"`
			CmdtCdtnCode        string        `json:"cmdtCdtnCode"`
			BkbnShpCode         interface{}   `json:"bkbnShpCode"`
			WhlRevwCont         float32       `json:"whlRevwCont"`
			RevwRvgrAvg         float32       `json:"revwRvgrAvg"`
			BestEmtnKywrName    interface{}   `json:"bestEmtnKywrName"`
			SaleCmdtClstName    interface{}   `json:"saleCmdtClstName"`
			ClturPlce           interface{}   `json:"clturPlce"`
			Period              interface{}   `json:"period"`
			ProductInfo         struct {
				SaleCmdtid      interface{} `json:"saleCmdtid"`
				Like            bool        `json:"like"`
				Basket          bool        `json:"basket"`
				Buy             bool        `json:"buy"`
				Direct          bool        `json:"direct"`
				ViewDetails     bool        `json:"viewDetails"`
				Sticky          interface{} `json:"sticky"`
				ReStockOnOff    bool        `json:"reStockOnOff"`
				ReleaseOnOff    bool        `json:"releaseOnOff"`
				ShippingCode    string      `json:"shippingCode"`
				ShippingText    string      `json:"shippingText"`
				ShippingKind    interface{} `json:"shippingKind"`
				TodayBook       bool        `json:"todayBook"`
				TodayBookLabel  interface{} `json:"todayBookLabel"`
				MdChoice        bool        `json:"mdChoice"`
				SpecialOrder    bool        `json:"specialOrder"`
				OnlyKyobo       bool        `json:"onlyKyobo"`
				LimitSale       bool        `json:"limitSale"`
				Gifts           bool        `json:"gifts"`
				Event           bool        `json:"event"`
				IncomeDeduction bool        `json:"incomeDeduction"`
				FixPrice        bool        `json:"fixPrice"`
				Bind            bool        `json:"bind"`
				CutPrice        bool        `json:"cutPrice"`
			} `json:"productInfo"`
		} `json:"tabContents"`
	} `json:"data"`
	StatusCode    float32     `json:"statusCode"`
	ResultCode    string      `json:"resultCode"`
	ResultMessage string      `json:"resultMessage"`
	DetailMessage interface{} `json:"detailMessage"`
}

type WordFrequency struct {
	CmdtName   string
	SaleCmdtID string
	Frequency  map[string]int
}
