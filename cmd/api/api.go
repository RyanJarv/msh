package api

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func New(args []string) (*Api, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("usage: %s <service> <action>", args)
	}

	opts, err := GetActionOpts(args[0], args[1])
	if err != nil {
		return nil, fmt.Errorf("getActionOpts: %w", err)
	}

	return &Api{
		ActionOpts: *opts,
	}, nil
}

type Api struct {
	rule       awsevents.Rule
	ActionOpts ActionOpts
}

func (s Api) GetName() string { return "sfn" }

func (s Api) Compile(stack awscdk.Stack, next interface{}) ([]interface{}, error) {
	var chain sfn.IChainable
	if next != nil {
		var ok bool
		chain, ok = next.(sfn.IChainable)
		if !ok {
			return nil, fmt.Errorf("next step must be chainable")
		}
	} else {
		chain = sfn.NewSucceed(stack, jsii.String("succeed"), &sfn.SucceedProps{})
	}

	block := sfn.NewPass(stack, jsii.String("pass"), &sfn.PassProps{})

	call := Paginate(stack, block, chain, s.ActionOpts)

	return []interface{}{call}, nil
}

func GetActionOpts(svc string, name string) (args *ActionOpts, err error) {
	resourceConfig, paginatorConfig, err := GetConfigs(svc)
	if err != nil {
		return nil, err
	}

	action, ok := resourceConfig.Service.HasMany[name]
	if !ok {
		return nil, fmt.Errorf("got %s, expected one of: %s", name, lo.Keys(resourceConfig.Service.HasMany))
	}

	paginator, ok := paginatorConfig.Pagination[action.Request.Operation]
	if !ok {
		return nil, fmt.Errorf("no paginator config found for %s", action.Request.Operation)
	}

	return &ActionOpts{
		Service:   svc,
		Action:    action.Request.Operation,
		Path:      action.Resource.Path,
		Paginator: paginator,
	}, nil
}

func GetConfigs(svc string) (*ResourceConfig, *PaginatorConfig, error) {
	resourceJson, err := getJsonConfig(boto3, filepath.Join("boto3", "boto3", "data", svc), "resources-1.json")
	if err != nil {
		return nil, nil, fmt.Errorf("resource json: %w", err)
	}

	var resourceConfig ResourceConfig
	if err := json.Unmarshal(resourceJson, &resourceConfig); err != nil {
		return nil, nil, err
	}

	content, err := getJsonConfig(botocore, filepath.Join("botocore", "botocore", "data", svc), "paginators-1.json")
	if err != nil {
		return nil, nil, fmt.Errorf("paginator json: %w", err)
	}

	var paginatorConfig PaginatorConfig
	if err := json.Unmarshal(content, &paginatorConfig); err != nil {
		return nil, nil, err
	}

	return &resourceConfig, &paginatorConfig, nil
}

func getJsonConfig(fs embed.FS, svcDir string, name string) ([]byte, error) {
	dir, err := fs.ReadDir(svcDir)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	latest, err := LatestVersion(dir)
	if err != nil {
		return nil, fmt.Errorf("newest version: %w", err)
	}

	content, err := fs.ReadFile(filepath.Join(svcDir, latest.Name(), name))
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	return content, nil
}

//func validateConfig(config ResourceConfig, content []byte) error {
//	var tmp map[string]interface{}
//	err := json.Unmarshal(content, &tmp)
//	if err != nil {
//		return fmt.Errorf("unmarshal: %w", err)
//	}
//
//	if diff := cmp.Diff(tmp, config); diff != "" {
//		return fmt.Errorf("config does not match original: %s", diff)
//	}
//
//	return nil
//}

func LatestVersion(dir []fs.DirEntry) (fs.DirEntry, error) {
	versions := map[time.Time]fs.DirEntry{}

	for _, ver := range dir {
		parse, err := time.Parse("2006-01-02", ver.Name())
		if err != nil {
			return nil, fmt.Errorf("parse time: %w", err)
		}

		versions[parse] = ver
	}

	ordered := lo.Keys(versions)
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].After(ordered[j])
	})

	return versions[ordered[0]], nil
}

type ActionOpts struct {
	Paginator
	Service string
	Action  string
	Path    string
}

func Paginate(stack awscdk.Stack, block sfn.Pass, next sfn.IChainable, input ActionOpts) sfn.Chain {
	// Stepfunctions doesn't accept [] in json paths, so we replace it with [*].
	path := strings.Replace(input.Path, "[]", "[*]", -1)
	path = fmt.Sprintf("$.%s", path)

	iter := sfn.NewMap(stack, jsii.String("map"), &sfn.MapProps{
		InputPath:  sfn.JsonPath_StringAt(jsii.String(path)),
		ResultPath: sfn.JsonPath_DISCARD(),
	})

	call := tasks.NewCallAwsService(stack, jsii.String("call"), &tasks.CallAwsServiceProps{
		Service:      jsii.String(input.Service),
		Action:       jsii.String(utils.UnTitle(input.Action)),
		IamResources: jsii.Strings("*"),
		Parameters:   &map[string]interface{}{
			//input.LimitKey: jsii.Number(5),
		},
	}).Next(iter)

	outputToken := fmt.Sprintf("$.%s", input.OutputToken)
	inputToken := fmt.Sprintf("%s.$", input.InputToken)

	paginator := sfn.Condition_IsPresent(jsii.String(outputToken))

	//override, ok := BrokenPaginators[input.Action]
	//if ok {
	//	paginator = sfn.Condition_And(paginator, sfn.Condition_IsPresent(jsii.String(override)))
	//}

	more := tasks.NewCallAwsService(stack, jsii.String("more"), &tasks.CallAwsServiceProps{
		Service:      jsii.String("ec2"),
		Action:       jsii.String("describeInstances"),
		IamResources: jsii.Strings("*"),
		Parameters: &map[string]interface{}{
			inputToken: jsii.String(outputToken),
			//input.LimitKey: jsii.Number(5),
		},
	}).Next(iter)

	iter.Iterator(block).
		Next(sfn.NewChoice(stack, jsii.String("choice"), &sfn.ChoiceProps{}).
			When(paginator, more, &sfn.ChoiceTransitionOptions{}).
			Otherwise(next))

	return call
}
