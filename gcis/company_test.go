package gcis

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCompanyService_GetBasicInformation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/od/data/api/5F64D864-61CB-4D0D-8AD9-492047CC1EA6", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(companyBasicInformationJSON)
	})

	got, _, err := client.Company.GetBasicInformation(context.Background(), &BasicInformationInput{"20828393"})
	if err != nil {
		t.Errorf("Company.GetBasicInformation returned error: %v", err)
	}
	if want := companyBasicInformation; !reflect.DeepEqual(got, want) {
		t.Errorf("Company.GetBasicInformation = %+v, want %+v", got, want)
	}
}

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

	companyBasicInformation = &BasicInformationOutput{
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

func TestCompanyService_GetCompanyByKeyword(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/od/data/api/6BBA2268-1367-4B42-9CCA-BC17499EBE8C", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(companyByKeywordJSON)
	})

	got, _, err := client.Company.GetCompanyByKeyword(context.Background(),
		&CompanyByKeywordInput{
			CompanyName:   "台灣積體電路製造股份有限公司",
			CompanyStatus: "01",
		})
	if err != nil {
		t.Errorf("Company.GetCompanyByKeyword returned error: %v", err)
	}
	if want := companyByKeyword; !reflect.DeepEqual(got, want) {
		t.Errorf("Company.GetCompanyByKeyword = %+v, want %+v", got, want)
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
