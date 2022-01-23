# orchestrator-service Documentation

## Instructions to run the Orchestration Service: </br>

To run the service navigate to the folder containing the repository and follow these instructions:
1) From the root directory containing the repo navigate to the server folder. Then run the following two command in the given order </br>
        **`go run OrchestratorServer1.go 9000`** </br>
        **`go run OrchestratorServer1.go 9001`**
2) From the root folder navigate to ./DataMock/DataMockService and run the command </br>
    **`go run DataMockService.go`**.
3) From the root folder navigate to client folder and run the command </br>
    **`go run client.go "name"`**     where "name" will be the name of the user that you want tio fetch
    
</br>
</br>

### Other Information
1) The operating system used is MacOs
2) GOPATH -> `/Users/rakshit/go`
3) repository path on local -> `/Users/rakshit/Desktop/orchestrator/orchestrator-service`


### Note to reader
As mentioned in Part 4, the logic folder is not created because the orchestration logic itself resides in the **OrchestratorServer.go** file in the function **GetUserByName**
