package gcis

import (
	"context"
	"fmt"
)

type BusinessService service

type BusinessBasicInformationInput struct {
	PresidentNo string
	Agency      string
}

type BusinessBasicInformationOutput struct {
	PresidentNo                  string `json:"President_No"`
	BusinessName                 string `json:"Business_Name"`
	BusinessCurrentStatus        string `json:"Business_Current_Status"`
	BusinessCurrentStatusDesc    string `json:"Business_Current_Status_Desc"`
	BusinessRegisterFunds        int64  `json:"Business_Register_Funds"`
	ResponsibleName              string `json:"responsible_name"`
	BusinessOrganizationType     string `json:"Business_Organization_Type"`
	BusinessOrganizationTypeDesc string `json:"Business_Organization_Type_Desc"`
	Agency                       string `json:"Agency"`
	AgencyDesc                   string `json:"Agency_Desc"`
	BusinessAddress              string `json:"Business_Address"`
	BusinessSetupApproveDate     string `json:"Business_Setup_Approve_Date"`
	BusinessLastChangeDate       string `json:"Business_Last_Change_Date"`
}

// GetBasicInformation fetches the basic information of company by president no and register agency.
func (s *BusinessService) GetBasicInformation(ctx context.Context, input *BusinessBasicInformationInput) (*BusinessBasicInformationOutput, *Response, error) {
	u := fmt.Sprintf("od/data/api/7E6AFA72-AD6A-46D3-8681-ED77951D912D?$format=json&$filter=President_No eq %s and Agency eq %s", input.PresidentNo, input.Agency)
	outputs := make([]BusinessBasicInformationOutput, 1)

	resp, err := s.client.get(ctx, u, &outputs)
	if err != nil {
		return nil, resp, err
	}
	if len(outputs) == 0 {
		return nil, resp, nil
	}
	return &outputs[0], resp, nil
}

type BusinessBasicInformationAndBusinessOutput struct {
	PresidentNo               string        `json:"President_No"`
	BusinessName              string        `json:"Business_Name"`
	BusinessCurrentStatus     string        `json:"Business_Current_Status"`
	BusinessCurrentStatusDesc string        `json:"Business_Current_Status_Desc"`
	Agency                    string        `json:"Agency"`
	AgencyDesc                string        `json:"Agency_Desc"`
	BusinessSetupApproveDate  string        `json:"Business_Setup_Approve_Date"`
	BusinessItemOld           []CmpBusiness `json:"Business_Item_Old"`
}

// GetBasicInformationAndBusiness fetches the basic information and business of company by president no and register agency.
func (s *BusinessService) GetBasicInformationAndBusiness(ctx context.Context, input *BusinessBasicInformationInput) (*BusinessBasicInformationAndBusinessOutput, *Response, error) {
	u := fmt.Sprintf("od/data/api/F570BC9A-DA4C-4813-8087-FB9CE95F9D38?$format=json&$filter=President_No eq %s and Agency eq %s", input.PresidentNo, input.Agency)
	outputs := make([]BusinessBasicInformationAndBusinessOutput, 1)

	resp, err := s.client.get(ctx, u, &outputs)
	if err != nil {
		return nil, resp, err
	}
	if len(outputs) == 0 {
		return nil, resp, nil
	}
	return &outputs[0], resp, nil
}
