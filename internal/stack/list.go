package stack

import (
	"context"
	"encoding/json"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/estenrye/octopod/internal/types"
)

func ListServices() ([]types.Service, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	services, err := cli.ServiceList(
		ctx,
		dockertypes.ServiceListOptions{},
	)

	if err != nil {
		return nil, err
	}

	var result []types.Service

	for _, service := range services {
		result = append(
			result,
			types.Service{
				Name:        service.Spec.Name,
				Image:       service.Spec.TaskTemplate.ContainerSpec.Image,
				Labels:      service.Spec.Labels,
				Environment: service.Spec.TaskTemplate.ContainerSpec.Env,
			},
		)
	}

	return result, nil
}

func ListServicesJson() (string, error) {
	var services, err = ListServices()
	if err != nil {
		return "[]", err
	}
	formatted, err := json.MarshalIndent(services, "", "  ")
	if err != nil {
		return "[]", err
	}
	return string(formatted), nil
}
