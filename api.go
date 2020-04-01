package awxcommunicator

import (
	"strconv"
)

//GetInventories - get inventories
func (c *AWXCommunicator) GetInventories() string {
	return (c.DoRequest("/inventories"))
}

//GetInventoriesHosts - get inventories
func (c *AWXCommunicator) GetInventoriesHosts(ID int) string {
	return (c.DoRequest("/inventories/" + strconv.Itoa(ID) + "/hosts"))
}

//GetAllHosts - Create request to AWX
func (c *AWXCommunicator) GetAllHosts() string {

	return (c.DoRequest("/hosts"))
}

//GetHosts - Create request to AWX
func (c *AWXCommunicator) GetHosts(ID int) string {

	return (c.DoRequest("/hosts/" + strconv.Itoa(ID)))
}

//GetAnsibleFacts - Create request to AWX
func (c *AWXCommunicator) GetAnsibleFacts(ID int) string {

	return (c.DoRequest("/hosts/" + strconv.Itoa(ID) + "/ansible_facts"))
}

//GetAllGroups - Create request to AWX
func (c *AWXCommunicator) GetAllGroups() string {

	return (c.DoRequest("/groups"))
}

//GetGroups - Create request to AWX
func (c *AWXCommunicator) GetGroups(ID int) string {

	return (c.DoRequest("/groups/" + strconv.Itoa(ID)))
}
