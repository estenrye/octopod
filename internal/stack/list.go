package stack

import (
	"context"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/estenrye/octopod/internal/types"
	"strings"
)

func ListStatusByName(name string) (types.ServiceSummaryList, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	f := filters.NewArgs()
	f.Add("name", name)

	services, err := cli.ServiceList(
		ctx,
		dockertypes.ServiceListOptions{
			Filters: f,
		},
	)

	if err != nil {
		return nil, err
	}

	var result types.ServiceSummaryList

	for _, service := range services {
		tasks, err := cli.TaskList(ctx, dockertypes.TaskListOptions{Filters:f})
		if err != nil {
			return nil, err
		}
		var s = types.ServiceSummary{
			Name:     service.Spec.Name,
			Image:    strings.SplitN(service.Spec.TaskTemplate.ContainerSpec.Image, ":", 2)[0],
			Tag:      strings.SplitN(service.Spec.TaskTemplate.ContainerSpec.Image, ":", 2)[1],
			Mode:     service.Spec.Mode,
			TaskList: make([]types.TaskSummary, 0),
		}

		for _, task := range tasks {
			s.TaskList = append(s.TaskList, types.TaskSummary{
				Id:           task.ID,
				Status:       task.Status,
				DesiredState: task.DesiredState,
			})
		}
		result = append(result, s)
	}

	return result, nil
}

func ListServicesByName(name string) (types.ServiceList, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	f := filters.NewArgs()
	f.Add("name", name)

	services, err := cli.ServiceList(
		ctx,
		dockertypes.ServiceListOptions{
			Filters: f,
		},
	)

	if err != nil {
		return nil, err
	}

	var result types.ServiceList

	for _, service := range services {
		var envVars = make(map[string]string, 20)
		for _, e := range service.Spec.TaskTemplate.ContainerSpec.Env {
			var v = strings.SplitN(e, "=", 2)
			envVars[v[0]] = v[1]
		}
		var image = strings.SplitN(service.Spec.TaskTemplate.ContainerSpec.Image, ":", 2)
		result = append(
			result,
			types.Service{
				Name:        service.Spec.Name,
				Image:       image[0],
				Tag:         image[1],
				Labels:      service.Spec.Labels,
				Environment: envVars,
			},
		)
	}

	return result, nil
}
