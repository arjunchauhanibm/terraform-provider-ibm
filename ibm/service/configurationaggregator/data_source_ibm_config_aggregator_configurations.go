// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

package configurationaggregator

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
)

func DataSourceIbmConfigAggregatorConfigurations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmConfigAggregatorConfigurationsRead,

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The account id of the resource.",
			},
			"config_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of resource configuration that are to be retrieved.",
			},
			"service_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the IBM Cloud service for which resources are to be retrieved.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource group id of the resources.",
			},
			"resource_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the resource.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The location or region in which the resources are created.",
			},
			"resource_crn": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The crn of the resource.",
			},
			"prev": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reference to the previous page of entries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reference to the previous page of entries.",
						},
						"start": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "the start string for the query to view the page.",
						},
					},
				},
			},
			"configs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of resource configurations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"about": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The basic metadata fetched from the query API.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The account ID in which the resource exists.",
									},
									"config_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of configuration of the retrieved resource.",
									},
									"resource_crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique CRN of the IBM Cloud resource.",
									},
									"resource_group_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The account ID.",
									},
									"service_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the service to which the resources belongs.",
									},
									"resource_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "User defined name of the resource.",
									},
									"last_config_refresh_time": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date/time stamp identifying when the information was last collected. Must be in the RFC 3339 format.",
									},
									"location": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Location of the resource specified.",
									},
									"tags": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Tags associated with the resource.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tag": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the tag.",
												},
											},
										},
									},
								},
							},
						},
						"config": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration of the resource.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmConfigAggregatorConfigurationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	configurationAggregatorClient, err := meta.(conns.ClientSession).ConfigurationAggregatorV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_configurations", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listConfigsOptions := &configurationaggregatorv1.ListConfigsOptions{}

	if _, ok := d.GetOk("config_type"); ok {
		listConfigsOptions.SetConfigType(d.Get("config_type").(string))
	}
	if _, ok := d.GetOk("account_id"); ok {
		listConfigsOptions.SetConfigType(d.Get("account_id").(string))
	}
	if _, ok := d.GetOk("resource_name"); ok {
		listConfigsOptions.SetConfigType(d.Get("resource_name").(string))
	}
	if _, ok := d.GetOk("service_name"); ok {
		listConfigsOptions.SetServiceName(d.Get("service_name").(string))
	}
	if _, ok := d.GetOk("resource_group_id"); ok {
		listConfigsOptions.SetResourceGroupID(d.Get("resource_group_id").(string))
	}
	if _, ok := d.GetOk("location"); ok {
		listConfigsOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("resource_crn"); ok {
		listConfigsOptions.SetResourceCrn(d.Get("resource_crn").(string))
	}

	var pager *configurationaggregatorv1.ConfigsPager
	pager, err = configurationAggregatorClient.NewConfigsPager(listConfigsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_configurations", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ConfigsPager.GetAll() failed %s", err), "(Data) ibm_config_aggregator_configurations", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmConfigAggregatorConfigurationsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIbmConfigAggregatorConfigurationsConfigToMap(&modelItem)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_configurations", "read")
			return tfErr.GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("configs", mapSlice); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting configs %s", err), "(Data) ibm_config_aggregator_configurations", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIbmConfigAggregatorConfigurationsID returns a reasonable ID for the list.
func dataSourceIbmConfigAggregatorConfigurationsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmConfigAggregatorConfigurationsPaginatedPreviousToMap(model *configurationaggregatorv1.PaginatedPrevious) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.Start != nil {
		modelMap["start"] = *model.Start
	}
	return modelMap, nil
}

func DataSourceIbmConfigAggregatorConfigurationsConfigToMap(model *configurationaggregatorv1.Config) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	aboutMap, err := DataSourceIbmConfigAggregatorConfigurationsAboutToMap(model.About)
	if err != nil {
		return modelMap, err
	}
	modelMap["about"] = []map[string]interface{}{aboutMap}
	configMap, err := DataSourceIbmConfigAggregatorConfigurationsConfigurationToMap(model.Config)
	if err != nil {
		return modelMap, err
	}
	modelMap["config"] = []map[string]interface{}{configMap}
	return modelMap, nil
}

func DataSourceIbmConfigAggregatorConfigurationsAboutToMap(model *configurationaggregatorv1.About) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["account_id"] = *model.AccountID
	modelMap["config_type"] = *model.ConfigType
	modelMap["resource_crn"] = *model.ResourceCrn
	modelMap["resource_group_id"] = *model.ResourceGroupID
	modelMap["service_name"] = *model.ServiceName
	modelMap["resource_name"] = *model.ResourceName
	modelMap["last_config_refresh_time"] = model.LastConfigRefreshTime.String()
	modelMap["location"] = *model.Location
	if model.Tags != nil {
		tagsMap, err := DataSourceIbmConfigAggregatorConfigurationsTagsToMap(model.Tags)
		if err != nil {
			return modelMap, err
		}
		modelMap["tags"] = []map[string]interface{}{tagsMap}
	}
	return modelMap, nil
}

func DataSourceIbmConfigAggregatorConfigurationsTagsToMap(model *configurationaggregatorv1.Tags) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Tag != nil {
		modelMap["tag"] = *model.Tag
	}
	return modelMap, nil
}

func DataSourceIbmConfigAggregatorConfigurationsConfigurationToMap(model *configurationaggregatorv1.Configuration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	return modelMap, nil
}