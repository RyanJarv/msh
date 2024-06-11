package awsapicallviacloudtrail
  // <para>Only update the policy if the revision ID matches the ID that's specified. Use this option to avoid modifying a policy that has changed since you last read it.</para>
type PutResourcePolicyRequestParameters struct {
    RegistryName string `json:"registryName"`
      // <para>The resource-based policy.</para>
    Policy string `json:"Policy"`
      // <para>The revision ID of the policy.</para>
    RevisionId string `json:"RevisionId,omitempty"`
}

func (p *PutResourcePolicyRequestParameters) SetRegistryName(registryName string) {
    p.RegistryName = registryName
}

func (p *PutResourcePolicyRequestParameters) SetPolicy(policy string) {
    p.Policy = policy
}

func (p *PutResourcePolicyRequestParameters) SetRevisionId(revisionId string) {
    p.RevisionId = revisionId
}
