/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// kubeCmd represents the kube command
var kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runKube,
}

func DefaultConnect() (*kubernetes.Clientset, error) {
	var (
		client     *kubernetes.Clientset
		kubeconfig []byte
		restConf   *rest.Config
		err        error
	)
	home := homedir.HomeDir()
	if kubeconfig, err = ioutil.ReadFile(home + "/.kube/config"); err != nil {
		goto FAIL
	}
	// 生成rest client配置
	if restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeconfig); err != nil {
		goto FAIL
	}
	if client, err = kubernetes.NewForConfig(restConf); err != nil {
		goto FAIL
	}
	return client, nil
FAIL:
	return nil, err
}

func runKube(cmd *cobra.Command, args []string) {
	//fmt.Println("kube called")
	var client *kubernetes.Clientset

	var b []byte
	var err error
	var deploymentRes *appsv1.Deployment

	client ,err= DefaultConnect()
	if b, err = yaml2.ToJSON([]byte(MysqlDeployTemplate)); err != nil {
		fmt.Println(err)
	}
	var deploymentReq appsv1.Deployment

	if err = json.Unmarshal(b, &deploymentReq); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("deployment: ",deploymentReq.Name,"创建中.....")

		deploymentRes,err = client.AppsV1().Deployments("default").Create(context.Background(),&deploymentReq,metav1.CreateOptions{})
		if err != nil{
			fmt.Println(err)
		}else {
			fmt.Println("deployment: ",deploymentRes.Name,"资源创建成功")
		}
	}

}

func init() {
	rootCmd.AddCommand(kubeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
