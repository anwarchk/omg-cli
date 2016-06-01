package vsphere_cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nats struct {

	/*Address - Descr: Address of the nats server Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Password - Descr: Password to connect to nats with Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Port - Descr: Port that the nats server listens on Default: 4222
*/
	Port interface{} `yaml:"port,omitempty"`

	/*User - Descr: Username to connect to nats with Default: nats
*/
	User interface{} `yaml:"user,omitempty"`

}