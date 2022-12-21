package fd

import (
	pipesTypes "github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"io"
)

type Sqs struct {
	io.ReadWriteCloser
	pipesTypes.PipeSourceParameters
	pipesTypes.PipeTargetParameters
}
