package main

// JobSchedule represents the job.
type JobSchedule struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CronSchedule string `json:"cronSchedule"`
	Command      string `json:"command"`
}

// Database represents the datastore.
type Database struct {
	bag map[int]JobSchedule
}

func NewDatabase() Database {
	return Database{
		bag: make(map[int]JobSchedule),
	}
}

func (d *Database) Add(job JobSchedule) {
	d.bag[job.ID] = job
}
