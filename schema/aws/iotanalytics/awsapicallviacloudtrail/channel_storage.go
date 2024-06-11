package awsapicallviacloudtrail

type ChannelStorage struct {
    CustomerManagedS3 CustomerManagedS3 `json:"customerManagedS3,omitempty"`
}

func (c *ChannelStorage) SetCustomerManagedS3(customerManagedS3 CustomerManagedS3) {
    c.CustomerManagedS3 = customerManagedS3
}
