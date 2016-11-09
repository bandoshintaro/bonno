package controllers


// modelsはあとで作成します
import (
  "database/sql"
  "gopkg.in/gorp.v1"
  _ "github.com/mattn/go-sqlite3"
  "github.com/revel/revel"
  "bonno/app/models"
)


// DbMapという変数にgorpのポインタを持たせる
var (
  DbMap *gorp.DbMap
)


//テーブルの初期化処理

func InitDB() {
  dbdir := revel.Config.StringDefault("bonno.dbdir","./app.db")
  db, err := sql.Open("sqlite3", dbdir)
  if err != nil {
    panic(err.Error())
  }
  DbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}


  // ここより下に、テーブルの作成の処理を書いていく
  t := DbMap.AddTable(models.Movie{}).SetKeys(true, "Id")
  t.ColMap("Name").MaxSize = 50

  DbMap.CreateTables()
}


// GorpContrllerの定義
type GorpController struct {
  *revel.Controller
  Transaction *gorp.Transaction
}


// これより以下はトランザクション処理
// 必ず定義する

// Begin()でトランザクションを開始して、Commit()でDBに反映、Rollback()で処理を元に戻す処理を書いていく


func (c *GorpController) Begin() revel.Result {
  txn, err := DbMap.Begin()
  if err != nil {
    panic(err)
  }
  c.Transaction = txn
  return nil
}

func (c *GorpController) Commit() revel.Result {
  if c.Transaction == nil {
    return nil
  }
  err := c.Transaction.Commit() // SQLによる変更をDBに反映
  if err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Transaction = nil // 正常終了した場合はROLLBACK処理に入らないようにする
  return nil
}


func (c *GorpController) Rollback() revel.Result {
  if c.Transaction == nil {
    return nil
  }
  err := c.Transaction.Rollback() // 問題があった場合変更前の状態に戻す
  if err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Transaction = nil
  return nil
}
