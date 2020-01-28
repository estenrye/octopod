package stack

import (
	"context"
	"github.com/docker/docker/client"
	dockertypes "github.com/docker/docker/api/types"
	"octopod/internal/types"
)

func ListServices(string) ([]types.Service, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	services, err := cli.ServiceList(
		ctx,
		dockertypes.ServiceListOptions{},
	)
	var result = []types.Service{}

	for _, service := range services {
		labels := []types.KeyValue{}

		for name, value := range service.Spec.Labels {
			labels = append(labels, types.KeyValue{
				Name: name,
				Value: value,
			})
		}

		result = append(result, types.Service{
			Name:service.Spec.Name,
			Labels: labels,
		},
		)
	}

	return result, nil
}