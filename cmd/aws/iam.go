package aws

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
)

//go:embed cli_iam_map.json
var CliIamMap []byte

type IamInfo struct {
	Action           string `json:"action"`
	ResourceMappings struct {
	} `json:"resource_mappings"`
	ResourcearnMappings struct {
		AggregationAuthorization string `json:"AggregationAuthorization"`
		ConfigRule               string `json:"ConfigRule"`
		ConfigurationAggregator  string `json:"ConfigurationAggregator"`
		ConformancePack          string `json:"ConformancePack"`
	} `json:"resourcearn_mappings"`
}

func IamActionsFromCliArgs(svcName, cmd string) ([]string, error) {
	//nonFlagArgs := lo.Filter(args, func(item string, index int) bool {
	//	return !strings.HasPrefix(item, "--")
	//})

	var cliIamMap map[string]map[string][]IamInfo
	if err := json.Unmarshal(CliIamMap, &cliIamMap); err != nil {
		return nil, fmt.Errorf("unmarshalling cli iam map: %w", err)
	}

	svc, ok := cliIamMap[svcName]
	if !ok {
		return nil, fmt.Errorf("unknown service: %s", svc)
	}

	iam, ok := svc[cmd]
	if !ok {
		return nil, fmt.Errorf("unknown action: %s", iam)
	}

	actions := lo.Map(iam, func(item IamInfo, index int) string {
		return item.Action
	})

	return actions, nil
}
