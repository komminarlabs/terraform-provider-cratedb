package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/komminarlabs/cratedb"
)

// OrganizationModel maps CrateDB organization schema data.
type OrganizationModel struct {
	Dc                   types.Object `tfsdk:"dc"`
	Email                types.String `tfsdk:"email"`
	Id                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	NotificationsEnabled types.Bool   `tfsdk:"notifications_enabled"`
	PlanType             types.Int32  `tfsdk:"plan_type"`
	ProjectCount         types.Int32  `tfsdk:"project_count"`
	RoleFQN              types.String `tfsdk:"role_fqn"`
}

// DatabasePartitionTemplateModel maps CrateDB organization DC schema data.
type OrganizationDCModel struct {
	Created  types.String `tfsdk:"created"`
	Modified types.String `tfsdk:"modified"`
}

func (o OrganizationDCModel) GetAttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"created":  types.StringType,
		"modified": types.StringType,
	}
}

func getOrganizationModel(ctx context.Context, organization cratedb.Organization) (*OrganizationModel, error) {
	dcValue := OrganizationDCModel{
		Created:  types.StringValue(organization.Dc.Created.String()),
		Modified: types.StringValue(organization.Dc.Modified.String()),
	}

	dcObjectValue, diags := types.ObjectValueFrom(ctx, dcValue.GetAttrType(), dcValue)
	if diags.HasError() {
		return nil, fmt.Errorf("error getting organization DC value")
	}

	organizationModel := OrganizationModel{
		Dc:                   dcObjectValue,
		Email:                types.StringValue(string(*organization.Email)),
		Id:                   types.StringValue(*organization.Id),
		Name:                 types.StringValue(organization.Name),
		NotificationsEnabled: types.BoolValue(*organization.NotificationsEnabled),
		PlanType:             types.Int32Value(int32(*organization.PlanType)),
		ProjectCount:         types.Int32Value(int32(*organization.ProjectCount)),
		RoleFQN:              types.StringValue(string(*organization.RoleFqn)),
	}
	return &organizationModel, nil
}