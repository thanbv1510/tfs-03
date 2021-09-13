package scheduler

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/robfig/cron/v3"
	"rabbit-demo/rabbitmq"
	"rabbit-demo/repository"
)

type Scheduler struct {
	db       *sql.DB
	cron     *cron.Cron
	ctx      context.Context
	producer *rabbitmq.Producer
}

func NewScheduler(ctx context.Context, db *sql.DB, producer *rabbitmq.Producer) *Scheduler {
	return &Scheduler{
		db:       db,
		cron:     cron.New(cron.WithSeconds()),
		ctx:      ctx,
		producer: producer,
	}
}

func (scheduler *Scheduler) Start() {
	scheduler.cron.AddFunc("30 * * * * *", scheduler.schedulerJob)
	scheduler.producer.Start()
	scheduler.cron.Start()
}

func (scheduler *Scheduler) schedulerJob() {
	fmt.Println("Scan to email....")
	repo := repository.NewOrderRepository(scheduler.db)
	emails := repo.GetUnSendEmailsWithLimit(10)
	for _, v := range emails {
		scheduler.producer.Public(v)
	}

}
