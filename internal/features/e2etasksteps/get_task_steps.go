package e2etasksteps

import (
	"time"

	"github.com/aws/amazon-ecs-event-stream-handler/internal/features/wrappers"
	. "github.com/gucumber/gucumber"
)

const (
	nonExistentTaskARN = "arn:aws:ecs:us-east-1:123456789012:task/31900037-daf4-40c7-93f7-102ece023cef"
)

func init() {

	eshWrapper := wrappers.NewESHWrapper()

	When(`^I get task with the cluster name and task ARN$`, func() {
		clusterName, err := wrappers.GetClusterName()
		if err != nil {
			T.Errorf(err.Error())
		}

		time.Sleep(15 * time.Second)
		if len(ecsTaskList) != 1 {
			T.Errorf("Error memorizing task started using ECS client")
		}
		taskARN := *ecsTaskList[0].TaskArn
		eshTask, err := eshWrapper.GetTask(clusterName, taskARN)
		if err != nil {
			T.Errorf(err.Error())
		}
		eshTaskList = append(eshTaskList, *eshTask)
	})

	Then(`^I get a task that matches the task started$`, func() {
		if len(ecsTaskList) != 1 || len(eshTaskList) != 1 {
			T.Errorf("Error memorizing results to validate them")
		}
		ecsTask := ecsTaskList[0]
		eshTask := eshTaskList[0]
		err := ValidateTasksMatch(ecsTask, eshTask)
		if err != nil {
			T.Errorf(err.Error())
		}
	})

	When(`^I try to get task with a non-existent ARN$`, func() {
		exceptionList = nil
		exception, err := eshWrapper.TryGetTask(nonExistentTaskARN)
		if err != nil {
			T.Errorf(err.Error())
		}
		exceptionList = append(exceptionList, exception)
	})

}
