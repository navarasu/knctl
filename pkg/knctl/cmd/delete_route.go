/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"

	"github.com/cppforlife/go-cli-ui/ui"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeleteRouteOptions struct {
	ui          ui.UI
	depsFactory DepsFactory

	RouteFlags RouteFlags
}

func NewDeleteRouteOptions(ui ui.UI, depsFactory DepsFactory) *DeleteRouteOptions {
	return &DeleteRouteOptions{ui: ui, depsFactory: depsFactory}
}

func NewDeleteRouteCmd(o *DeleteRouteOptions, flagsFactory FlagsFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "route",
		Short: "Delete route",
		Example: `
  # Delete route 'route1' in namespace 'ns1'
  knctl delete route --route route1 -n ns1`,
		RunE: func(_ *cobra.Command, _ []string) error { return o.Run() },
	}
	o.RouteFlags.Set(cmd, flagsFactory)
	return cmd
}

func (o *DeleteRouteOptions) Run() error {
	servingClient, err := o.depsFactory.ServingClient()
	if err != nil {
		return err
	}

	err = servingClient.ServingV1alpha1().Routes(o.RouteFlags.NamespaceFlags.Name).Delete(o.RouteFlags.Name, &metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("Deleting route: %s", err)
	}

	// TODO idempotent?

	return nil
}
