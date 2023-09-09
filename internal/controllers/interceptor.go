package controllers

import "sigs.k8s.io/controller-runtime/pkg/client"

type CapsuleInterceptor struct {
	Client client.Client
}
