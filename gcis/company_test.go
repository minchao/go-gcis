package gcis

import (
	"context"
	"reflect"
	"testing"
)

var (
	companyBasicInformationJSON = []byte(`[
  {
    "Business_Accounting_NO": "20828393",
    "Company_Status_Desc": "核准設立",
    "Company_Name": "宏碁股份有限公司",
    "Capital_Stock_Amount": 35000000000,
    "Paid_In_Capital_Amount": 30765028280,
    "Responsible_Name": "陳O聖",
    "Company_Location": "臺北市松山區民福里復興北路369號7樓之5",
    "Register_Organization_Desc": "經濟部商業司",
    "Company_Setup_Date": "0680718",
    "Change_Of_Approval_Data": "1060905",
    "Revoke_App_Date": "",
    "Case_Status": "",
    "Case_Status_Desc": "",
    "Sus_App_Date": "",
    "Sus_Beg_Date": "",
    "Sus_End_Date": ""
  }
]`)

	companyBasicInformation = &CompanyBasicInformationOutput{
		BusinessAccountingNO:     "20828393",
		CompanyStatusDesc:        "核准設立",
		CompanyName:              "宏碁股份有限公司",
		CapitalStockAmount:       35000000000,
		PaidInCapitalAmount:      30765028280,
		ResponsibleName:          "陳O聖",
		CompanyLocation:          "臺北市松山區民福里復興北路369號7樓之5",
		RegisterOrganizationDesc: "經濟部商業司",
		CompanySetupDate:         "0680718",
		ChangeOfApprovalData:     "1060905",
		RevokeAppDate:            "",
		CaseStatus:               "",
		CaseStatusDesc:           "",
		SusAppDate:               "",
		SusBegDate:               "",
		SusEndDate:               "",
	}
)

func TestCompanyService_GetBasicInformation(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/5F64D864-61CB-4D0D-8AD9-492047CC1EA6", companyBasicInformationJSON)

	got, _, err := client.Company.GetBasicInformation(context.Background(), &CompanyBasicInformationInput{"20828393"})
	if err != nil {
		t.Errorf("Company.GetBasicInformation returned error: %v", err)
	}
	if want := companyBasicInformation; !reflect.DeepEqual(got, want) {
		t.Errorf("Company.GetBasicInformation = %+v, want %+v", got, want)
	}
}

func TestCompanyService_GetBasicInformation_notFound(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/5F64D864-61CB-4D0D-8AD9-492047CC1EA6", nil)

	got, _, err := client.Company.GetBasicInformation(context.Background(), &CompanyBasicInformationInput{"20828393"})
	if err != nil {
		t.Errorf("Company.GetBasicInformation returned error: %v", err)
	}
	if got != nil {
		t.Errorf("Company.GetBasicInformation = %+v, want nil", got)
	}
}

var (
	companyBasicInformationAndBusinessJSON = []byte(`[
  {
    "Business_Accounting_NO": "20828393",
    "Company_Name": "宏碁股份有限公司",
    "Company_Status": "01",
    "Company_Status_Desc": "核准設立",
    "Company_Setup_Date": "0680718",
    "Cmp_Business": [
      {
        "Business_Seq_NO": "0001",
        "Business_Item": "F113050",
        "Business_Item_Desc": "電腦及事務性機器設備批發業"
      }
    ]
  }
]`)

	companyBasicInformationAndBusiness = &BasicInformationAndBusinessOutput{
		BusinessAccountingNO: "20828393",
		CompanyName:          "宏碁股份有限公司",
		CompanyStatus:        "01",
		CompanyStatusDesc:    "核准設立",
		CompanySetupDate:     "0680718",
		CmpBusiness: []CmpBusiness{
			{
				BusinessSeqNO:    "0001",
				BusinessItem:     "F113050",
				BusinessItemDesc: "電腦及事務性機器設備批發業",
			},
		},
	}
)

func TestCompanyService_GetBasicInformationAndBusiness(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/236EE382-4942-41A9-BD03-CA0709025E7C", companyBasicInformationAndBusinessJSON)

	got, _, err := client.Company.GetBasicInformationAndBusiness(context.Background(), &CompanyBasicInformationInput{"20828393"})
	if err != nil {
		t.Errorf("Company.GetBasicInformationAndBusiness returned error: %v", err)
	}
	if want := companyBasicInformationAndBusiness; !reflect.DeepEqual(got, want) {
		t.Errorf("Company.GetBasicInformationAndBusiness = %+v, want %+v", got, want)
	}
}

func TestCompanyService_GetBasicInformationAndBusiness_notFound(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/236EE382-4942-41A9-BD03-CA0709025E7C", nil)

	got, _, err := client.Company.GetBasicInformationAndBusiness(context.Background(), &CompanyBasicInformationInput{"20828393"})
	if err != nil {
		t.Errorf("Company.GetBasicInformationAndBusiness returned error: %v", err)
	}
	if got != nil {
		t.Errorf("Company.GetBasicInformationAndBusiness = %+v, want nil", got)
	}
}

var (
	companyByKeywordJSON = []byte(`[
  {
    "Business_Accounting_NO": "22099131",
    "Company_Name": "台灣積體電路製造股份有限公司",
    "Company_Status": "01",
    "Company_Status_Desc": "核准設立",
    "Capital_Stock_Amount": 270500000000,
    "Paid_In_Capital_Amount": 259303804580,
    "Responsible_Name": "LOu Te-YOn Mark(劉德音)",
    "Register_Organization": "05",
    "Register_Organization_Desc": "科技部新竹科學工業園區管理局",
    "Company_Location": "新竹市力行六路8號",
    "Company_Setup_Date": "0760221",
    "Change_Of_Approval_Data": "1071128"
  }
]`)

	companyByKeyword = []CompanyByKeywordOutput{
		{
			BusinessAccountingNO:     "22099131",
			CompanyName:              "台灣積體電路製造股份有限公司",
			CompanyStatus:            "01",
			CompanyStatusDesc:        "核准設立",
			CapitalStockAmount:       270500000000,
			PaidInCapitalAmount:      259303804580,
			ResponsibleName:          "LOu Te-YOn Mark(劉德音)",
			RegisterOrganization:     "05",
			RegisterOrganizationDesc: "科技部新竹科學工業園區管理局",
			CompanyLocation:          "新竹市力行六路8號",
			CompanySetupDate:         "0760221",
			ChangeOfApprovalData:     "1071128",
		},
	}
)

func TestCompanyService_SearchByKeyword(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/6BBA2268-1367-4B42-9CCA-BC17499EBE8C", companyByKeywordJSON)

	got, _, err := client.Company.SearchByKeyword(context.Background(),
		&CompanyByKeywordInput{
			CompanyName:   "台灣積體電路製造股份有限公司",
			CompanyStatus: "01",
		})
	if err != nil {
		t.Errorf("Company.SearchByKeyword returned error: %v", err)
	}
	if want := companyByKeyword; !reflect.DeepEqual(got, want) {
		t.Errorf("Company.SearchByKeyword = %+v, want %+v", got, want)
	}
}

func TestCompanyService_SearchByKeyword_notFound(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/6BBA2268-1367-4B42-9CCA-BC17499EBE8C", nil)

	got, _, err := client.Company.SearchByKeyword(context.Background(),
		&CompanyByKeywordInput{
			CompanyName:   "台灣積體電路製造股份有限公司",
			CompanyStatus: "01",
		})
	if err != nil {
		t.Errorf("Company.SearchByKeyword returned error: %v", err)
	}
	if want := []CompanyByKeywordOutput{}; !reflect.DeepEqual(got, want) {
		t.Errorf("Company.SearchByKeyword = %+v, want %+v", got, want)
	}
}

var (
	companyByResponsibleNameJSON = []byte(`[
  {
    "Business_Accounting_NO": "22099131",
    "Company_Name": "台灣積體電路製造股份有限公司"
  }
]`)

	companyByResponsibleName = []CompanyByResponsibleNameOutput{
		{
			BusinessAccountingNO: "22099131",
			CompanyName:          "台灣積體電路製造股份有限公司",
		},
	}
)

func TestCompanyService_SearchByResponsibleName(t *testing.T) {
	setup()
	defer teardown()

	handle(t, "/od/data/api/4B61A0F1-458C-43F9-93F3-9FD6DA5E1B08", companyByResponsibleNameJSON)

	got, _, err := client.Company.SearchByResponsibleName(context.Background(),
		&CompanyByResponsibleNameInput{
			ResponsibleName: "劉德音",
		})
	if err != nil {
		t.Errorf("Company.SearchByResponsibleName returned error: %v", err)
	}
	if want := companyByResponsibleName; !reflect.DeepEqual(got, want) {
		t.Errorf("Company.SearchByResponsibleName = %+v, want %+v", got, want)
	}
}
