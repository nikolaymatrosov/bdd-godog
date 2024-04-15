package steps

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type DiskFeature struct {
	imageFamily string
	imageFolder string
	imageId     string
	diskName    string
	diskFolder  string
	diskId      string

	sdk *ycsdk.SDK
}

func (df *DiskFeature) anImageWithFamilyInFolder(family, folder string) error {
	df.imageFamily = family
	df.imageFolder = folder

	image, err := df.sdk.Compute().Image().GetLatestByFamily(context.Background(), &compute.GetImageLatestByFamilyRequest{
		Family:   family,
		FolderId: folder,
	})
	if err != nil {
		return err
	}

	df.imageId = image.Id

	return nil
}

func (df *DiskFeature) iCreateADiskFromItWithNameInTheFolder(name, folder string) error {
	df.diskName = name
	df.diskFolder = folder

	operation, err := df.sdk.WrapOperation(df.sdk.Compute().Disk().Create(context.Background(), &compute.CreateDiskRequest{
		FolderId: folder,
		Name:     name,
		Source:   &compute.CreateDiskRequest_ImageId{ImageId: df.imageId},
		ZoneId:   "ru-central1-a",
		Size:     10 * 1024 * 1024 * 1024,
	}))
	if err != nil {
		return err
	}
	err = operation.Wait(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	resp, err := operation.Response()
	if err != nil {
		log.Fatal(err)
	}
	disk := resp.(*compute.Disk)

	df.diskId = disk.Id

	return nil
}

func (df *DiskFeature) iShouldSeeTheDiskInTheFolder() error {

	disk, err := df.sdk.Compute().Disk().List(context.Background(), &compute.ListDisksRequest{
		FolderId: df.diskFolder,
		Filter:   `name = "` + df.diskName + `"`,
	})

	if err != nil {
		return err
	}
	if len(disk.Disks) == 0 {
		return fmt.Errorf("disk %s not found in folder %s", df.diskName, df.diskFolder)
	}

	return nil
}

func (df *DiskFeature) iWantToDeleteTheDisk() error {
	_, err := df.sdk.WrapOperation(df.sdk.Compute().Disk().Delete(context.Background(), &compute.DeleteDiskRequest{
		DiskId: df.diskId,
	}))
	if err != nil {
		return err
	}

	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{".."},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	df := &DiskFeature{}
	ctx := context.Background()
	token := os.Getenv("YC_OAUTH_TOKEN")
	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: ycsdk.OAuthToken(token),
	})
	if err != nil {
		log.Fatal(err)
	}
	df.sdk = sdk

	sc.Step(`^an image with family "([^"]*)" in the "([^"]*)" folder$`, df.anImageWithFamilyInFolder)
	sc.Step(`^I create a disk from it with name "([^"]*)" in the "([^"]*)" folder$`, df.iCreateADiskFromItWithNameInTheFolder)
	sc.Step(`^I should see the disk in the folder$`, df.iShouldSeeTheDiskInTheFolder)
	sc.Step(`^I want to delete the disk$`, df.iWantToDeleteTheDisk)
}
