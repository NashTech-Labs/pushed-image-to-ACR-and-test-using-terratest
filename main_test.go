package main


import (

    "testing"
    "fmt"
    "os/exec"
    "strings"
	"github.com/stretchr/testify/assert"

)

var (
	acrURL = "testacrmk.azurecr.io" // ACR login server URL
	acrName = "testacrmk"  // ACR name
    testImage = "hello-world:latest" // local test image
	expectedPushedResponse = true  // for test the status of docker image pushed or not
)


func TestTerraformOutputs(t *testing.T) {

	tag := "1.1.0"
	tagImageName := fmt.Sprintf("%s/hello-world:%s", acrURL, tag)
	fmt.Println(tagImageName)
	TagImageWithAcr(testImage, tagImageName, t)

	actualPushedResponse := PushImageToAcr(tagImageName)
	// fmt.Println(statusCode)

	t.Run(fmt.Sprintf("Checking the images is pushed or not to ACR: %s", acrName), func(t *testing.T) {

        assert.Equal(t, expectedPushedResponse, actualPushedResponse, "publicNetworkAccess mismatch")

    })
    


}


func TagImageWithAcr(testImage, tagImageName string, t *testing.T) {
	acrTagCmd := exec.Command("docker", "tag", testImage, tagImageName)
	err := acrTagCmd.Run()
	if err != nil {
		t.Errorf("error Tagging image with ACR Repo: %v", err)
		return
	}
	fmt.Printf("Image %s Tagged with %s Successfully\n", testImage, tagImageName)
}



func PushImageToAcr(acrImageName string) bool {
	acrPushCmd := exec.Command("docker", "push", acrImageName)
	output, err := acrPushCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error pushing image to ACR: %v\n", err)
		return false
	}

	successMessage := fmt.Sprintf("Image %s Pushed to ACR Successfully", acrImageName)
	if strings.Contains(string(output), successMessage) {
		fmt.Printf("%s [Muzakkir]\n", successMessage)
		return true
	} else if strings.Contains(string(output), "digest") {
		fmt.Printf("Image %s Pushed to ACR Successfully [Muzakkir]\n", acrImageName)
		return true
	}

	return false
}





