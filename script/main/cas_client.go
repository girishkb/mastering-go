package main

import (
	casv1 "bitbucket.org/swigy/protorepo/location-platform/central-address-service/v1"
	"google.golang.org/grpc"
)

func NewCasClient(endpoint string) (casv1.CentralAddressAPIClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return casv1.NewCentralAddressAPIClient(conn), nil
}
