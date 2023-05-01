//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status displaying:
//  - number of servers

func serverStatus(s map[string]int) {
	fmt.Println("\nThe number of servers are:", len(s))

	//  - number of servers for each status (Online, Offline, Maintenance, Retired)
	numberofStatuses := make(map[int]int)

	for _, status := range s {
		switch status {
		case 0:
			numberofStatuses[0] += 1
		case 1:
			numberofStatuses[1] += 1
		case 2:
			numberofStatuses[2] += 1
		case 3:
			numberofStatuses[3] += 1
		default:
			panic("Server staus unknown")
		}
	}
	fmt.Printf("Online: %d Offline: %d Maintenance: %d Retired: %d\n\n", numberofStatuses[0], numberofStatuses[1], numberofStatuses[2], numberofStatuses[3])

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverList := make(map[string]int)

	//* Create a map using the server names as the key and the server status
	for _, v := range servers {
		serverList[v] = Online
	}
	serverStatus(serverList)
	//  - change server status of `darkstar` to `Retired`
	//  - change server status of `aiur` to `Offline`
	//  - call display server info function
	serverList["darkstar"] = Retired
	serverList["aiur"] = Offline
	serverStatus(serverList)
	//  - change server status of all servers to `Maintenance`
	//  - call display server info function
	for k := range serverList {
		serverList[k] = Maintenance
	}
	serverStatus(serverList)

}
