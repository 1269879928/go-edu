package dao

import (
	"fmt"
	kgorm "github.com/heyuanlong/go-tools/db/gorm"
	"go-edu/work/base/inits"
	"go-edu/work/entity"
)
func TestInsert() bool  {
	transaction := kgorm.NewTransaction()
	tx := transaction.Begin(inits.Gorm)
	defer transaction.Defer()
	test1 := &entity.Test1{
		T1Name: "t1",
		T1Age:  10,
	}
	err := tx.Create(test1).Error
	if err != nil {
		_ = transaction.Rollback()
		fmt.Println("err:", err)
		return false
	}
	test2 := &entity.Test2{
		T2Name: "t2",
		T1Id:  test1.Id,
	}
	err = tx.Create(test2).Error
	if err != nil {
		_ = transaction.Rollback()
		fmt.Println("err:", err)
		return false
	}
	//_ = transaction.Commit()
	return true
}
