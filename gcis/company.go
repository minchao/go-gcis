package gcis

import (
	"context"
	"fmt"
)

type CompanyService service

type BasicInformationInput struct {
	BusinessAccountingNO string
}

type BasicInformationOutput struct {
	BusinessAccountingNO     string `json:"Business_Accounting_NO"`
	CompanyStatusDesc        string `json:"Company_Status_Desc"`
	CompanyName              string `json:"Company_Name"`
	CapitalStockAmount       int64  `json:"Capital_Stock_Amount"`
	PaidInCapitalAmount      int64  `json:"Paid_In_Capital_Amount"`
	ResponsibleName          string `json:"Responsible_Name"`
	CompanyLocation          string `json:"Company_Location"`
	RegisterOrganizationDesc string `json:"Register_Organization_Desc"`
	CompanySetupDate         string `json:"Company_Setup_Date"`
	ChangeOfApprovalData     string `json:"Change_Of_Approval_Data"`
	RevokeAppDate            string `json:"Revoke_App_Date"`
	CaseStatus               string `json:"Case_Status"`
	CaseStatusDesc           string `json:"Case_Status_Desc"`
	SusAppDate               string `json:"Sus_App_Date"`
	SusBegDate               string `json:"Sus_Beg_Date"`
	SusEndDate               string `json:"Sus_End_Date"`
}

// GetBasicInformation fetches the basic information of company by accounting no.
func (s *CompanyService) GetBasicInformation(ctx context.Context, input *BasicInformationInput) (*BasicInformationOutput, *Response, error) {
	u := fmt.Sprintf("od/data/api/5F64D864-61CB-4D0D-8AD9-492047CC1EA6?$format=json&$filter=Business_Accounting_NO eq %s", input.BusinessAccountingNO)
	outputs := make([]BasicInformationOutput, 1)

	resp, err := s.client.get(ctx, u, &outputs)
	if err != nil {
		return nil, resp, err
	}
	if len(outputs) == 0 {
		return nil, resp, nil
	}
	return &outputs[0], resp, nil
}

type BasicInformationAndBusinessOutput struct {
	BusinessAccountingNO string        `json:"Business_Accounting_NO"`
	CompanyName          string        `json:"Company_Name"`
	CompanyStatus        string        `json:"Company_Status"`
	CompanyStatusDesc    string        `json:"Company_Status_Desc"`
	CompanySetupDate     string        `json:"Company_Setup_Date"`
	CmpBusiness          []CmpBusiness `json:"Cmp_Business"`
}

type CmpBusiness struct {
	BusinessSeqNO    string `json:"Business_Seq_NO"`
	BusinessItem     string `json:"Business_Item"`
	BusinessItemDesc string `json:"business_item_desc"`
}

// GetBasicInformationAndBusiness fetches the basic information and business of company by accounting no.
func (s *CompanyService) GetBasicInformationAndBusiness(ctx context.Context, input *BasicInformationInput) (*BasicInformationAndBusinessOutput, *Response, error) {
	u := fmt.Sprintf("od/data/api/236EE382-4942-41A9-BD03-CA0709025E7C?$format=json&$filter=Business_Accounting_NO eq %s", input.BusinessAccountingNO)
	outputs := make([]BasicInformationAndBusinessOutput, 1)

	resp, err := s.client.get(ctx, u, &outputs)
	if err != nil {
		return nil, resp, err
	}
	if len(outputs) == 0 {
		return nil, resp, nil
	}
	return &outputs[0], resp, nil
}

type CompanyByKeywordInput struct {
	CompanyName   string
	CompanyStatus string
	Skip          int
	Top           int
}

type CompanyByKeywordOutput struct {
	BusinessAccountingNO string `json:"Business_Accounting_NO"`
	CompanyName          string `json:"Company_Name"`
	// Status see https://data.gcis.nat.gov.tw/od/cmpStatusCodeData?type=xls
	CompanyStatus            string `json:"Company_Status"`
	CompanyStatusDesc        string `json:"Company_Status_Desc"`
	CapitalStockAmount       int64  `json:"Capital_Stock_Amount"`
	PaidInCapitalAmount      int64  `json:"Paid_In_Capital_Amount"`
	ResponsibleName          string `json:"Responsible_Name"`
	RegisterOrganization     string `json:"Register_Organization"`
	RegisterOrganizationDesc string `json:"Register_Organization_Desc"`
	CompanyLocation          string `json:"Company_Location"`
	CompanySetupDate         string `json:"Company_Setup_Date"`
	ChangeOfApprovalData     string `json:"Change_Of_Approval_Data"`
}

// GetCompanyByKeyword fetches the information of company by keyword.
func (s *CompanyService) GetCompanyByKeyword(ctx context.Context, input *CompanyByKeywordInput) ([]CompanyByKeywordOutput, *Response, error) {
	if input.Top == 0 {
		input.Top = 50
	}
	u := fmt.Sprintf("od/data/api/6BBA2268-1367-4B42-9CCA-BC17499EBE8C?$format=json&$filter=Company_Name like %s and Company_Status eq %s&$skip=%d&$top=%d",
		input.CompanyName,
		input.CompanyStatus,
		input.Skip,
		input.Top)
	outputs := make([]CompanyByKeywordOutput, 1)

	resp, err := s.client.get(ctx, u, &outputs)
	if err != nil {
		return nil, resp, err
	}
	return outputs, resp, nil
}

type CompanyByResponsibleNameInput struct {
	ResponsibleName string `json:"Responsible_Name"`
	Skip            int
	Top             int
}

type CompanyByResponsibleNameOutput struct {
	BusinessAccountingNO string `json:"Business_Accounting_NO"`
	CompanyName          string `json:"Company_Name"`
}

// SearchByResponsibleName searches the companies by responsible name.
func (s *CompanyService) SearchByResponsibleName(ctx context.Context, input *CompanyByResponsibleNameInput) ([]CompanyByResponsibleNameOutput, *Response, error) {
	if input.Top == 0 {
		input.Top = 50
	}
	u := fmt.Sprintf("od/data/api/4B61A0F1-458C-43F9-93F3-9FD6DA5E1B08?$format=json&$filter=Responsible_Name eq %s&$skip=%d&$top=%d",
		input.ResponsibleName,
		input.Skip,
		input.Top)
	outputs := make([]CompanyByResponsibleNameOutput, 1)

	resp, err := s.client.get(ctx, u, &outputs)
	if err != nil {
		return nil, resp, err
	}
	return outputs, resp, nil
}
