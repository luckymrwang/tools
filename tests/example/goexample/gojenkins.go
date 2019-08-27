package main

import (
	"fmt"

	"github.com/bndr/gojenkins"
)

func main() {
	jenkins := gojenkins.CreateJenkins(nil, "http://localhost:8080/", "sino", "200800")
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert
	_, err := jenkins.Init()

	if err != nil {
		panic("Something Went Wrong")
	}

	build, err := jenkins.GetJob("job_name")
	if err != nil {
		panic("Job Does Not Exist")
	}

	lastSuccessBuild, err := build.GetLastSuccessfulBuild()
	if err != nil {
		panic("Last SuccessBuild does not exist")
	}

	duration := lastSuccessBuild.GetDuration()
	fmt.Println(duration)

	job, err := jenkins.GetJob("jobname")

	if err != nil {
		panic("Job does not exist")
	}

	job.Rename("SomeotherJobName")

	configString := `<?xml version='1.0' encoding='UTF-8'?>
<project>
  <actions/>
  <description></description>
  <keepDependencies>false</keepDependencies>
  <properties/>
  <scm class="hudson.scm.NullSCM"/>
  <canRoam>true</canRoam>
  <disabled>false</disabled>
  <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
  <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
  <triggers class="vector"/>
  <concurrentBuild>false</concurrentBuild>
  <builders/>
  <publishers/>
  <buildWrappers/>
</project>`

	jenkins.CreateJob(configString, "someNewJobsName")
}
