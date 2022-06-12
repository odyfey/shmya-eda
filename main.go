package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jszwec/csvutil"
)

type Order struct {
	Date    OrderDate `csv:"date"`
	Cutlery int       `csv:"cutlery"`
	Tips    int       `csv:"tips"`
	Price   int       `csv:"order_price"`
	UserID  int       `csv:"uid"`
	ID      int       `csv:"order_id"`
}

type OrderDate struct {
	time.Time
}

const format = "2006-01-02 15:04:05"

func (t *OrderDate) UnmarshalCSV(data []byte) error {
	tt, err := time.Parse(format, string(data))
	if err != nil {
		return err
	}
	*t = OrderDate{Time: tt}
	return nil
}

const (
	shmyaSourceCSV = "shmya_final_version.csv"
	shmyaOutputCSV = "shmya_users_segment.csv"
)

func main() {
	res, err := os.ReadFile(shmyaSourceCSV)
	if err != nil {
		log.Fatalf("failed to read csv source file, %v", err)
	}

	var orders []Order
	if err := csvutil.Unmarshal(res, &orders); err != nil {
		log.Fatalf("failed to unmarshal orders to struct, %v", err)
	}

	if isGreater := tipsHypothesisForHolidays(orders); isGreater {
		log.Print("users with more cutlery (>2) leave more tips on holidays\n\n")
	} else {
		log.Print("users with fewer cutlery (<=2) leave more tips on holidays\n\n")
	}

	if isGreater := tipsHypothesisForAnotherDays(orders); isGreater {
		log.Print("users with more cutlery (>2) leave more tips on another days\n\n")
	} else {
		log.Print("users with fewer cutlery (<=2) leave more tips on another days\n\n")
	}

	usersIDs := usersSegment(orders)

	f, err := os.Create(shmyaOutputCSV)
	if err != nil {
		log.Fatalf("failed to create output file, %v", err)
	}

	w := csv.NewWriter(f)
	w.Write([]string{"uid"})

	for _, userID := range usersIDs {
		line := []string{strconv.Itoa(userID)}
		if err := w.Write(line); err != nil {
			log.Printf("can't write line to output file, userID = %d\n", userID)
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatalf("failed to write buffered data, %v", err)
	}

	log.Printf("users segment saved to %s\n", shmyaOutputCSV)
	log.Printf("users segment count = %d", len(usersIDs))
}
