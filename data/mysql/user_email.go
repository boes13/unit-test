package mysql

import (
	"database/sql"
	"fmt"

	"github.com/boes13/unit-test/common"
)

func GetUserEmail(userID int64) (string, bool, error) {
	db := common.GetDBObject()
	query := fmt.Sprintf("select email from user where id=%d", userID)
	row := db.QueryRow(query)

	var email string
	err := row.Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false, nil
		}
		return "", false, err
	}

	return email, true, nil
}
