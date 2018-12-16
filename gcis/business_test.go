package gcis

import (
	"context"
	"reflect"
	"testing"
)

var (
	businessBasicInformationJSON = []byte(`[
  {
    "President_No": "26459190",
    "Business_Name": "鼎勝冷榨油行",
    "Business_Current_Status": "01",
    "Business_Current_Status_Desc": "核准設立",
    "Business_Register_Funds": 968000,
    "Responsible_Name": "朱O勝",
    "Business_Organization_Type": "06",
    "Business_Organization_Type_Desc": "獨資",
    "Agency": "376610000A",
    "Agency_Desc": "臺南市政府",
    "Business_Address": "臺南市安平區華平里怡平路485號1樓",
    "Business_Setup_Approve_Date": "1001012",
    "Business_Last_Change_Date": "1010507"
  }
]`)

	businessBasicInformation = &BusinessBasicInformationOutput{
		PresidentNo:                  "26459190",
		BusinessName:                 "鼎勝冷榨油行",
		BusinessCurrentStatus:        "01",
		BusinessCurrentStatusDesc:    "核准設立",
		BusinessRegisterFunds:        968000,
		ResponsibleName:              "朱O勝",
		BusinessOrganizationType:     "06",
		BusinessOrganizationTypeDesc: "獨資",
		Agency:                       "376610000A",
		AgencyDesc:                   "臺南市政府",
		BusinessAddress:              "臺南市安平區華平里怡平路485號1樓",
		BusinessSetupApproveDate:     "1001012",
		BusinessLastChangeDate:       "1010507",
	}
)

func TestBusinessService_GetBasicInformation(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/7E6AFA72-AD6A-46D3-8681-ED77951D912D", businessBasicInformationJSON)

	got, _, err := client.Bussiness.GetBasicInformation(context.Background(),
		&BusinessBasicInformationInput{
			PresidentNo: "26459190",
			Agency:      "376610000A",
		})
	if err != nil {
		t.Errorf("Bussiness.GetBasicInformation returned error: %v", err)
	}
	if want := businessBasicInformation; !reflect.DeepEqual(got, want) {
		t.Errorf("Bussiness.GetBasicInformation = %+v, want %+v", got, want)
	}
}

func TestBusinessService_GetBasicInformation_notFound(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/7E6AFA72-AD6A-46D3-8681-ED77951D912D", nil)

	got, _, err := client.Bussiness.GetBasicInformation(context.Background(), &BusinessBasicInformationInput{})
	if err != nil {
		t.Errorf("Bussiness.GetBasicInformation returned error: %v", err)
	}
	if got != nil {
		t.Errorf("Bussiness.GetBasicInformation = %+v, want nil", got)
	}
}

var (
	businessBasicInformationAndBusinessJSON = []byte(`[
  {
    "President_No": "26459190",
    "Business_Name": "鼎勝冷榨油行",
    "Business_Current_Status": "01",
    "Business_Current_Status_Desc": "核准設立",
    "Agency": "376610000A",
    "Agency_Desc": "臺南市政府",
    "Business_Setup_Approve_Date": "1001012",
    "Business_Item_Old": [
      {
        "Business_Seq_No": "1",
        "Business_Item": "C105010",
        "Business_Item_Desc": "食用油脂製造業"
      }
    ]
  }
]`)

	businessBasicInformationAndBusiness = &BusinessBasicInformationAndBusinessOutput{
		PresidentNo:               "26459190",
		BusinessName:              "鼎勝冷榨油行",
		BusinessCurrentStatus:     "01",
		BusinessCurrentStatusDesc: "核准設立",
		Agency:                    "376610000A",
		AgencyDesc:                "臺南市政府",
		BusinessSetupApproveDate:  "1001012",
		BusinessItemOld: []CmpBusiness{
			{
				BusinessSeqNO:    "1",
				BusinessItem:     "C105010",
				BusinessItemDesc: "食用油脂製造業",
			},
		},
	}
)

func TestBusinessService_GetBasicInformationAndBusiness(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/F570BC9A-DA4C-4813-8087-FB9CE95F9D38", businessBasicInformationAndBusinessJSON)

	got, _, err := client.Bussiness.GetBasicInformationAndBusiness(context.Background(),
		&BusinessBasicInformationInput{
			PresidentNo: "26459190",
			Agency:      "376610000A",
		})
	if err != nil {
		t.Errorf("Bussiness.GetBasicInformationAndBusiness returned error: %v", err)
	}
	if want := businessBasicInformationAndBusiness; !reflect.DeepEqual(got, want) {
		t.Errorf("Bussiness.GetBasicInformationAndBusiness = %+v, want %+v", got, want)
	}
}

func TestBusinessService_GetBasicInformationAndBusiness_notFound(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/F570BC9A-DA4C-4813-8087-FB9CE95F9D38", nil)

	got, _, err := client.Bussiness.GetBasicInformationAndBusiness(context.Background(), &BusinessBasicInformationInput{})
	if err != nil {
		t.Errorf("Bussiness.GetBasicInformationAndBusiness returned error: %v", err)
	}
	if got != nil {
		t.Errorf("Bussiness.GetBasicInformationAndBusiness = %+v, want nil", got)
	}
}
