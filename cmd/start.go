// Copyright © 2018 TFGCo backend@tfgco.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/topfreegames/resources-check/model"
	"k8s.io/client-go/kubernetes"
)

var incluster bool
var context string
var kubeconfig string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts worker",
	Long: `Initializes Kubernetes API and starts worker that at every checks deployments, statefulsets and daemonsets
	on the cluster. If there is some with not specified an event is sent to monitoring systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		log := configureLogger().WithFields(logrus.Fields{
			"source":    "worker",
			"operation": "start",
		})

		log.Info("starting resources-check worker")

		var kubernetesClient kubernetes.Interface
		worker, err := model.NewWorker(
			config,
			kubernetesClient,
			log,
			incluster,
			kubeconfig,
		)
		if err != nil {
			log.Fatal(err)
		}

		worker.Start()
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVar(
		&incluster, "incluster", false, "incluster mode (for running on kubernetes)")
	startCmd.Flags().StringVar(
		&context, "context", "", "kubeconfig context")
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	startCmd.Flags().StringVar(
		&kubeconfig, "kubeconfig",
		fmt.Sprintf("%s/.kube/config", home),
		"path to the kubeconfig file (not needed if using --incluster)")
}
