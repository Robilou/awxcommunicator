package awxcommunicator

import "time"

type (

	//Role -
	Role struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		ID          int    `json:"id"`
	}

	//EditUser -
	EditUser struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	//Inventories - http://localhost/api/v2/inventories/
	Inventories struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			ID      int    `json:"id"`
			Type    string `json:"type"`
			URL     string `json:"url"`
			Related struct {
				CreatedBy              string `json:"created_by"`
				ModifiedBy             string `json:"modified_by"`
				Hosts                  string `json:"hosts"`
				Groups                 string `json:"groups"`
				RootGroups             string `json:"root_groups"`
				VariableData           string `json:"variable_data"`
				Script                 string `json:"script"`
				Tree                   string `json:"tree"`
				InventorySources       string `json:"inventory_sources"`
				UpdateInventorySources string `json:"update_inventory_sources"`
				ActivityStream         string `json:"activity_stream"`
				JobTemplates           string `json:"job_templates"`
				AdHocCommands          string `json:"ad_hoc_commands"`
				AccessList             string `json:"access_list"`
				ObjectRoles            string `json:"object_roles"`
				InstanceGroups         string `json:"instance_groups"`
				Copy                   string `json:"copy"`
				Organization           string `json:"organization"`
			} `json:"related"`
			SummaryFields struct {
				Organization struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Description string `json:"description"`
				} `json:"organization"`
				CreatedBy   []EditUser `json:"created_by"`
				ModifiedBy  []EditUser `json:"modified_by"`
				ObjectRoles struct {
					AdminRole  []Role `json:"admin_role"`
					UpdateRole []Role `json:"update_role"`
					AdhocRole  []Role `json:"adhoc_role"`
					UseRole    []Role `json:"use_role"`
					ReadRole   []Role `json:"read_role"`
				} `json:"object_roles"`
				UserCapabilities struct {
					Edit   bool `json:"edit"`
					Delete bool `json:"delete"`
					Copy   bool `json:"copy"`
					Adhoc  bool `json:"adhoc"`
				} `json:"user_capabilities"`
			} `json:"summary_fields"`
			Created                      time.Time `json:"created"`
			Modified                     time.Time `json:"modified"`
			Name                         string    `json:"name"`
			Description                  string    `json:"description"`
			Organization                 int       `json:"organization"`
			Kind                         string    `json:"kind"`
			HostFilter                   string    `json:"host_filter"`
			Variables                    string    `json:"variables"`
			HasActiveFailures            bool      `json:"has_active_failures"`
			TotalHosts                   int       `json:"total_hosts"`
			HostsWithActiveFailures      int       `json:"hosts_with_active_failures"`
			TotalGroups                  int       `json:"total_groups"`
			HasInventorySources          bool      `json:"has_inventory_sources"`
			TotalInventorySources        int       `json:"total_inventory_sources"`
			InventorySourcesWithFailures int       `json:"inventory_sources_with_failures"`
			InsightsCredential           string    `json:"insights_credential"`
			PendingDeletion              bool      `json:"pending_deletion"`
		} `json:"results"`
	}

	//Hosts - http://localhost/api/v2/hosts/
	Hosts struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			ID      int    `json:"id"`
			Type    string `json:"type"`
			URL     string `json:"url"`
			Related struct {
				CreatedBy          string `json:"created_by"`
				ModifiedBy         string `json:"modified_by"`
				VariableData       string `json:"variable_data"`
				Groups             string `json:"groups"`
				AllGroups          string `json:"all_groups"`
				JobEvents          string `json:"job_events"`
				JobHostSummaries   string `json:"job_host_summaries"`
				ActivityStream     string `json:"activity_stream"`
				InventorySources   string `json:"inventory_sources"`
				SmartInventories   string `json:"smart_inventories"`
				AdHocCommands      string `json:"ad_hoc_commands"`
				AdHocCommandEvents string `json:"ad_hoc_command_events"`
				Insights           string `json:"insights"`
				AnsibleFacts       string `json:"ansible_facts"`
				Inventory          string `json:"inventory"`
			} `json:"related"`
			SummaryFields struct {
				Inventory struct {
					ID                           int    `json:"id"`
					Name                         string `json:"name"`
					Description                  string `json:"description"`
					HasActiveFailures            bool   `json:"has_active_failures"`
					TotalHosts                   int    `json:"total_hosts"`
					HostsWithActiveFailures      int    `json:"hosts_with_active_failures"`
					TotalGroups                  int    `json:"total_groups"`
					HasInventorySources          bool   `json:"has_inventory_sources"`
					TotalInventorySources        int    `json:"total_inventory_sources"`
					InventorySourcesWithFailures int    `json:"inventory_sources_with_failures"`
					OrganizationID               int    `json:"organization_id"`
					Kind                         string `json:"kind"`
				} `json:"inventory"`
				CreatedBy        []EditUser `json:"created_by"`
				ModifiedBy       []EditUser `json:"modified_by"`
				UserCapabilities struct {
					Edit   bool `json:"edit"`
					Delete bool `json:"delete"`
				} `json:"user_capabilities"`
				Groups struct {
					Count   int    `json:"count"`
					Results string `json:"results"`
				} `json:"groups"`
				RecentJobs string `json:"recent_jobs"`
			} `json:"summary_fields"`
			Created              time.Time `json:"created"`
			Modified             time.Time `json:"modified"`
			Name                 string    `json:"name"`
			Description          string    `json:"description"`
			Inventory            int       `json:"inventory"`
			Enabled              bool      `json:"enabled"`
			InstanceID           string    `json:"instance_id"`
			Variables            string    `json:"variables"`
			HasActiveFailures    bool      `json:"has_active_failures"`
			HasInventorySources  bool      `json:"has_inventory_sources"`
			LastJob              string    `json:"last_job"`
			LastJobHostSummary   string    `json:"last_job_host_summary"`
			InsightsSystemID     string    `json:"insights_system_id"`
			AnsibleFactsModified string    `json:"ansible_facts_modified"`
		} `json:"results"`
	}

	//Host
	Host struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		URL     string `json:"url"`
		Related struct {
			NamedURL           string `json:"named_url"`
			CreatedBy          string `json:"created_by"`
			ModifiedBy         string `json:"modified_by"`
			VariableData       string `json:"variable_data"`
			Groups             string `json:"groups"`
			AllGroups          string `json:"all_groups"`
			JobEvents          string `json:"job_events"`
			JobHostSummaries   string `json:"job_host_summaries"`
			ActivityStream     string `json:"activity_stream"`
			InventorySources   string `json:"inventory_sources"`
			SmartInventories   string `json:"smart_inventories"`
			AdHocCommands      string `json:"ad_hoc_commands"`
			AdHocCommandEvents string `json:"ad_hoc_command_events"`
			Insights           string `json:"insights"`
			AnsibleFacts       string `json:"ansible_facts"`
			Inventory          string `json:"inventory"`
		} `json:"related"`
		SummaryFields struct {
			Inventory struct {
				ID                           int    `json:"id"`
				Name                         string `json:"name"`
				Description                  string `json:"description"`
				HasActiveFailures            bool   `json:"has_active_failures"`
				TotalHosts                   int    `json:"total_hosts"`
				HostsWithActiveFailures      int    `json:"hosts_with_active_failures"`
				TotalGroups                  int    `json:"total_groups"`
				HasInventorySources          bool   `json:"has_inventory_sources"`
				TotalInventorySources        int    `json:"total_inventory_sources"`
				InventorySourcesWithFailures int    `json:"inventory_sources_with_failures"`
				OrganizationID               int    `json:"organization_id"`
				Kind                         string `json:"kind"`
			} `json:"inventory"`
			CreatedBy        []EditUser `json:"created_by"`
			ModifiedBy       []EditUser `json:"modified_by"`
			UserCapabilities struct {
				Edit   bool `json:"edit"`
				Delete bool `json:"delete"`
			} `json:"user_capabilities"`
			Groups struct {
				Count   int    `json:"count"`
				Results string `json:"results"`
			} `json:"groups"`
			RecentJobs string `json:"recent_jobs"`
		} `json:"summary_fields"`
		Created              time.Time `json:"created"`
		Modified             time.Time `json:"modified"`
		Name                 string    `json:"name"`
		Description          string    `json:"description"`
		Inventory            int       `json:"inventory"`
		Enabled              bool      `json:"enabled"`
		InstanceID           string    `json:"instance_id"`
		Variables            string    `json:"variables"`
		HasActiveFailures    bool      `json:"has_active_failures"`
		HasInventorySources  bool      `json:"has_inventory_sources"`
		LastJob              string    `json:"last_job"`
		LastJobHostSummary   string    `json:"last_job_host_summary"`
		InsightsSystemID     string    `json:"insights_system_id"`
		AnsibleFactsModified string    `json:"ansible_facts_modified"`
	}
)
