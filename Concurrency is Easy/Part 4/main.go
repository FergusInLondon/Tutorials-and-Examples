package main

func main() {
	ctx, cancel := 
	wg := sync.WaitGroup{}

	input := make(chan Report)
	go runPipeline(ctx, &wg, input)

	// Input example reports

	input <- Report{

	}
	
	// We're going to rely upon channel closure to cascade down
	close

	// As a demonstration, we don't want a long running pipeline, so we will
	// cancel it after one second.
	<- time.After(1 * time.Second)
	cancel()
	wg.Wait()
}

func runPipeline(ctx context.Context, wg *sync.WaitGroup, input chan Report) {
	metadata := make(chan Report)
	dispatcher := make(chan Report)
	database := make(chan Report)
	distribute := make(chan Report)

	go populateUserData(ctx, wg, input, metadata)
	go populateMetadata(ctx, wg, metadata, dispatcher)
	go dispatchAlert(ctx, wg, dispatcher, database)
	go persistToDatabase(ctx, wg, database, distribute)
	go distributeToConnectedClients(ctx, wg, distribute)

	wg.Wait()
}