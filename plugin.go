package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)
type (
	Repo struct {
		Owner string `json:"owner"`
		Name  string `json:"name"`
	}

	Build struct {
		Number  int    `json:"number"`
		Link    string `json:"link"`
	}

	Config struct {
		Token         string
		WaitPipelines  []string
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
	}
)

func (p Plugin) Exec() error {
	host := "https://cloud.drone.io"
	if hostPathArr := strings.Split(p.Build.Link, "/"); len(hostPathArr) > 2 {
		host = hostPathArr[0] + "//" + hostPathArr[2]
	}

	// create an http client with oauth authentication.
	config := new(oauth2.Config)
	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: p.Config.Token,
		},
	)
	// create the drone client with authenticator
	client := drone.NewClient(host, auther)

	if len(p.Repo.Owner) > 0 && len(p.Repo.Name) > 0 {
		for {
			var successFlg int
			var elementCnt int
			if gotBuild, err := client.Build(p.Repo.Owner, p.Repo.Name, p.Build.Number); err == nil {
				for _, element := range gotBuild.Stages {
					for _, v := range p.Config.WaitPipelines {
						if v == element.Name {
							fmt.Println(element.Name, element.Status)
							elementCnt++
							if element.Status != drone.StatusWaiting && element.Status != drone.StatusPending && element.Status != drone.StatusRunning {
								if element.Status != drone.StatusPassing {
									os.Exit(1)
								} else {
									successFlg++
								}
							}
							break;
						}
					}
				}
			}
			if successFlg == elementCnt {
				os.Exit(0)
			}
			time.Sleep(10 * time.Second)
		}
	}

	return nil
}

// Function checks if int is in slice of ints
func intInSlice(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
