package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	gojenkins "gitlab.ushareit.me/sgt/scmp-jenkins"
)

func main() {
	jc, err := gojenkins.CreateJenkins(nil, "https://jenkins-dev.ushareit.me", "admin", "1152594940cd434aaafa12951370ce96").Init()
	if err != nil {
		log.Errorf("CreateJenkins client error : %s", err)
		return
	}

	job, err := jc.GetSubJob("SGT_scmp-cicd_bkNNnRGDuw", "scmp-event")
	if err != nil {
		log.Errorf("GetJob: %s", err)
		return
	}

	build, err := job.GetBuild(9)
	if err != nil {
		log.Errorf("GetBuild error : %s", err)
		return
	}
	// fmt.Println(build.BuildResponse())

	envs, err := build.GetInjectedEnvVars()
	if err != nil {
		log.Errorf("GetInjectedEnvVars error : %s", err)
		return
	}
	fmt.Println(envs)

	/*
		query := map[string]string{
			"fullStages": "true",
			"_": fmt.Sprintf("%d", time.Now().Unix()),
		}
		pr, err := job.GetPipelineRun("9", query)
		if err != nil {
			log.Errorf("GetPipelineRuns: %s", err)
			return
		}

		fmt.Println(pr.String())
	*/

}
