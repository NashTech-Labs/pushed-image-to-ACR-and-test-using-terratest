# Push the docker image to the ACR and run the test case to check the image is pushed or not.

### When you create a ACR and push the image to ACR then you need to validate that image is pushed or not. You can test this by using terratest. 

### Follow the below steps to run the terratest code:


You need to define these values in your code before running:-

                acrURL = "testacrmk.azurecr.io" // ACR login server URL
                acrName = "testacrmk"  // ACR name
                testImage = "hello-world:latest" // local test image
                expectedPushedResponse = true  // for test the status of docker image pushed or not

Step 1:- Run the go initialization command:

            go mod init < name >

Step 2:- Run the tidy command to install the packages:-

            go mod tidy

Step 3:- Run the test command:-

            go test -v
