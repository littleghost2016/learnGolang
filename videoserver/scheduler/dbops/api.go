package dbops

/*
1. user->api service->delete video
2. api service->scheduler->writer video deletion record
3. timer
4. timer->runner->read wvdr(writer video deletion record)->exec->delete video from folder
*/

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddVideoDeletionRecord(vid string) error {
	smtmIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = smtmIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer smtmIns.Close()
	return nil
}