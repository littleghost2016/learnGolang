package dbops

import (
	"log"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")

	// stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ?") //这是第一次错误的地方
	if err != nil {
		log.Printf("Delete user error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info
	(id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime:ctime}

	defer stmtIns.Close()
	return res, nil
}