package vspherex

import (
	"context"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

// VsphereClient implements Machine interface to Create/Update/Delete VM's.
type VsphereClient struct {
	kubeadmToken   string
	ctx            context.Context
}

//
type MachineActuatorParams struct {
	// vsphere api coordinates

	KubeadmToken string
	//TODO Add parameters
}

//
func NewMachineActuator(params MachineActuatorParams) (*VsphereClient, error) {
	//TODO prepare stuff to access vsphere api
	return &VsphereClient{
		kubeadmToken:   params.KubeadmToken,
		ctx:            context.Background(),
	}, nil
}

func (azure *VsphereClient) Create(cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
/*	_, err := azure.createOrUpdateGroup(cluster)
	if err != nil {
		return err
	}
	_, err = azure.createOrUpdateDeployment(cluster, machine)
	if err != nil {
		return err
	}*/
	//Get the Login info from the VMs
	/*
		_, _, err = azure.getLogin(cluster, machine)
		if err != nil {
			return err
		}
	*/
	//Set up Kubernetes
	return nil
}

func (azure *VsphereClient) Update(cluster *clusterv1.Cluster, goalMachine *clusterv1.Machine) error {
/*	//Parse in configurations
	var goalMachineConfig azureconfigv1.AzureMachineProviderConfig
	err := azure.decodeMachineProviderConfig(goalMachine.Spec.ProviderConfig, &goalMachineConfig)
	if err != nil {
		return err
	}
	var clusterConfig azureconfigv1.AzureClusterProviderConfig
	err = azure.decodeClusterProviderConfig(cluster.Spec.ProviderConfig, &clusterConfig)
	if err != nil {
		return err
	}
	_, err = azure.vmIfExists(cluster, goalMachine)
	if err != nil {
		return err
	}*/
	return nil
}

func (azure *VsphereClient) Delete(cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
/*	//Parse in configurations
	var machineConfig azureconfigv1.AzureMachineProviderConfig
	err := azure.decodeMachineProviderConfig(machine.Spec.ProviderConfig, &machineConfig)
	if err != nil {
		return err
	}
	var clusterConfig azureconfigv1.AzureClusterProviderConfig
	err = azure.decodeClusterProviderConfig(cluster.Spec.ProviderConfig, &clusterConfig)
	if err != nil {
		return err
	}
	//Check if the machine exists
	vm, err := azure.vmIfExists(cluster, machine)
	if err != nil {
		return err
	}
	if vm == nil {
		//Skip deleting if we couldn't find anything to delete
		return nil
	}
	*//*
		TODO: See if this is the last remaining machine, and if so,
		delete the resource group, which will automatically delete
		all associated resources
	*//*
	groupsClient := resources.NewGroupsClient(azure.SubscriptionID)
	groupsClient.Authorizer = azure.Authorizer
	groupsDeleteFuture, err := groupsClient.Delete(azure.ctx, clusterConfig.ResourceGroup)
	if err != nil {
		return err
	}
	return groupsDeleteFuture.Future.WaitForCompletion(azure.ctx, groupsClient.BaseClient.Client)*/
	return nil
}

func (azure *VsphereClient) Exists(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (bool, error) {
/*	vm, err := azure.vmIfExists(cluster, machine)
	if err != nil {
		return false, err
	}
	return (vm != nil), nil*/
	return false, nil
}


func (azure *VsphereClient) GetKubeConfig(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (string, error) {
	return "", nil
}
