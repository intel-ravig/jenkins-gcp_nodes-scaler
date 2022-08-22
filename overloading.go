func autoScaling() {
	for {
		queueSize := fetchQueueSize()
		queueSize = adjustQueueSizeDependingWhetherJobRequiringAllNodesIsRunning(queueSize)
		//log.Printf("%d jobs waiting to be executed\n", queueSize)
		if queueSize > 0 {
			log.Printf("%d jobs waiting to be executed\n", queueSize)
			enableMoreNodes(queueSize)
		} else if queueSize == 0 {
			log.Println("No jobs in the queue")
			disableUnnecessaryBuildBoxes()
		}

		log.Println("Iteration finished")
		fmt.Println("")
		time.Sleep(time.Second * 8)
	}
}