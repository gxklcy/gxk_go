package main

import (
	"database/sql"
	"fmt"
	"log"
)

package main

import (
    "database/sql"
    "github.com/pkg/errors"
)

func main() {
	names, err := DaoQueryNames(20)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(names)
}
func DaoQueryNames(age int) (ret []string, err error)  {
	err =db.QueryRow("select name from users where age = ?", age).Scan(&ret)
	if err != nil {
		if err == sql.ErrNoRows {
			// 如果无数据，我认为不需要再往上抛error了
			return []string{}, nil
		} else {
			return []string{}, err
		}
	}
	return
}

