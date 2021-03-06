package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Entities struct {

	/*Description - Descr: Text associated with the VMs Default: vcd-cf
*/
	Description interface{} `yaml:"description,omitempty"`

	/*VappCatalog - Descr: The name of the calalog for vapp template Default: <nil>
*/
	VappCatalog interface{} `yaml:"vapp_catalog,omitempty"`

	/*VmMetadataKey - Descr: The key name of VM metadata Default: vcd-cf
*/
	VmMetadataKey interface{} `yaml:"vm_metadata_key,omitempty"`

	/*Organization - Descr: The organization name Default: <nil>
*/
	Organization interface{} `yaml:"organization,omitempty"`

	/*VirtualDatacenter - Descr: The virtual data center name in vCloud Director Default: <nil>
*/
	VirtualDatacenter interface{} `yaml:"virtual_datacenter,omitempty"`

	/*MediaCatalog - Descr: The name of the calalog for media files Default: <nil>
*/
	MediaCatalog interface{} `yaml:"media_catalog,omitempty"`

}