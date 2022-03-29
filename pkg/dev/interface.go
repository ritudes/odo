package dev

import (
	"io"

	"github.com/redhat-developer/odo/pkg/devfile/adapters/common"

	"github.com/redhat-developer/odo/pkg/watch"

	"github.com/devfile/library/pkg/devfile/parser"
	"github.com/redhat-developer/odo/pkg/devfile/adapters/kubernetes"
)

type Client interface {
	// Start the resources in devfileObj on the platformContext. It then pushes the files in path to the container.
	Start(devfileObj parser.DevfileObj, platformContext kubernetes.KubernetesContext, path string) error

	// Watch watches for any changes to the files under path while ignoring the files/directories in ignorePaths.
	// It logs messages to out and uses the Handler h to perform push operation when anything changes in path.
	// It uses devfileObj to notify user to restart odo dev if they change endpoint information in the devfile.
	Watch(devfileObj parser.DevfileObj, path string, ignorePaths []string, out io.Writer, h Handler) error

	// Cleanup cleans the resources created by Start
	Cleanup() error
}

type Handler interface {
	RegenerateAdapterAndPush(common.PushParameters, watch.WatchParameters) error
}