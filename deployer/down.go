package deployer

import (
	"fmt"

	"k8s.io/klog/v2"
	"sigs.k8s.io/kubetest2/pkg/exec"
)

// Down should tear down the test cluster if any
func (ad *aksDeployer) Down() error {
	up, err := ad.IsUp()
	if err != nil {
		return fmt.Errorf("failed to get cluster state: %w", err)
	}

	if !up {
		klog.Info("No cluster, skipping down.")
		return nil
	}

	klog.Infof("Deleting Resources: %+v", *ad)
	if err := deleteResource(ad.clusterResourceID); err != nil {
		return fmt.Errorf("failed to delete resource: %w", err)
	}
	klog.Info("Deleted resources")
	return nil
}

func deleteResource(resourceID string) error {
	_, err := runWithErrorOutput(exec.Command("az", "resource", "delete", "--ids", resourceID))
	return err
}
