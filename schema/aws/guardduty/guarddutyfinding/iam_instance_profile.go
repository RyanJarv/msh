package guarddutyfinding

type IamInstanceProfile struct {
    Arn string `json:"arn,omitempty"`
    Id string `json:"id,omitempty"`
}

func (i *IamInstanceProfile) SetArn(arn string) {
    i.Arn = arn
}

func (i *IamInstanceProfile) SetId(id string) {
    i.Id = id
}
