package main

type User struct {
	Name string
	Title string
	Role string
	ID string
}

func (u *User) Authorised() {
	return true
}

func GetUser() *User {
	return &User{
		Name: "Unlucky Reporter",
		Title: "Systems Administrator",
		Role: "REPORTER",
		ID: "12345",
	}
}

type Report struct {
	EventDescription string
	ReportedAt time.Time
	Status string
	ReporterName string
	ReporterTitle string
	ReporterID string
	Severity int
}

func populateUserData(in, out chan Report) {
	defer wg.Done()
	defer close(out)

	for {
		select {
		case <- ctx.Done():
			return
		case report, closed := <-in:
			if closed {
				return
			}

			user := GetUser()
			if user.Authorised() {
				report.ReporterName = user.Name
				report.ReporterTitle = user.Title
				report.ReporterID = user.ID
				out <- report
			} else {
				log.Println("access by unauthorised account, discarding report")
			}
		}
	}
}

func populateMetadata(in, out chan Report) {
	defer wg.Done()
	defer close(out)

	for {
		select {
		case <- ctx.Done():
			return
		case report, closed := <-in:
			if closed {
				return
			}

			// Populate other metadata
			report.ReportedAt = time.Now()
			report.Status = "UNRESOLVED"

			out <- report
		}
	}
}

func dispatchAlert(in, out chan Report) {
	defer wg.Done()
	defer close(out)

	for {
		select {
		case <- ctx.Done():
			return
		case report, closed := <-in:
			if closed {
				return
			}

			if report.Severity > 2 {
				log.Println("Pushing report w/ UserID to queue for notification service")
			}

			out <- report
		}
	}
}

func persistToDatabase(in, out chan Report) {
	defer wg.Done()
	defer close(out)

	for {
		select {
		case <- ctx.Done():
			return
		case report, closed := <-in:
			if closed {
				return
			}

			log.Println("Persisting report to the database")
			out <- report
		}
	}
}

func distributeToConnectedClients(in chan Report) {
	defer wg.Done()
	defer close(out)

	for {
		select {
		case <- ctx.Done():
			return
		case report, closed := <-in:
			if closed {
				return
			}

			log.Println("Sending alert to all distributed clients via ws connection")
		}
	}
}
