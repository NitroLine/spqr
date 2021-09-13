package qrouter

import (
	"strconv"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/pg-sharding/spqr/internal/config"
	"github.com/pg-sharding/spqr/internal/qdb"
	"github.com/pg-sharding/spqr/internal/qdb/etcdcl"
	"github.com/pg-sharding/spqr/yacc/spqrparser"
	"github.com/pkg/errors"
	"github.com/wal-g/tracelog"
)

type ShardQrouter struct {
	ColumnMapping map[string]struct{}

	LocalTables map[string]struct{}

	Ranges map[string]qdb.KeyRange

	ShardCfgs map[string]*config.ShardCfg

	qdb qdb.QrouterDB
}

var _ Qrouter = &ShardQrouter{}

func NewShardQrouter() (*ShardQrouter, error) {
	db, err := etcdcl.NewQDBETCD()
	if err != nil {
		return nil, err
	}

	return &ShardQrouter{
		ColumnMapping: map[string]struct{}{},
		LocalTables:   map[string]struct{}{},
		Ranges:        map[string]qdb.KeyRange{},
		ShardCfgs:     map[string]*config.ShardCfg{},
		qdb:           db,
	}, nil
}

func (qr *ShardQrouter) Subscribe(krid string, krst qdb.KeyRangeStatus, noitfyio chan<- interface{}) error {
	return qr.qdb.Watch(krid, krst, noitfyio)
}

func (qr *ShardQrouter) Unite(req *spqrparser.UniteKeyRange) error {
	panic("implement me")
}

func (qr *ShardQrouter) Split(req *spqrparser.SplitKeyRange) error {
	if err := qr.qdb.Begin(); err != nil {
		return err
	}

	defer func() { _ = qr.qdb.Commit() }()

	krOld := qr.Ranges[req.KeyRangeFromID]
	krNew := qdb.KeyRange{
		From:       req.Border,
		To:         krOld.To,
		KeyRangeID: req.KeyRangeID,
	}

	_ = qr.qdb.Add(krNew)
	krOld.To = req.Border
	_ = qr.qdb.Update(krOld)

	qr.Ranges[krOld.KeyRangeID] = krOld
	qr.Ranges[krNew.KeyRangeID] = krNew

	return nil
}

func (qr *ShardQrouter) Lock(krid string) error {
	var kr qdb.KeyRange
	var ok bool

	if kr, ok = qr.Ranges[krid]; !ok {
		return errors.Errorf("key range with id %v not found", krid)
	}

	return qr.qdb.Lock(kr)
}

func (qr *ShardQrouter) UnLock(krid string) error {
	var kr qdb.KeyRange
	var ok bool

	if kr, ok = qr.Ranges[krid]; !ok {
		return errors.Errorf("key range with id %v not found", krid)
	}

	return qr.qdb.UnLock(kr)
}

func (qr *ShardQrouter) AddShard(name string, cfg *config.ShardCfg) error {

	tracelog.InfoLogger.Printf("adding node %s", name)
	qr.ShardCfgs[name] = cfg

	return nil
}

func (qr *ShardQrouter) Shards() []string {

	var ret []string

	for name := range qr.ShardCfgs {
		ret = append(ret, name)
	}

	return ret
}

func (qr *ShardQrouter) KeyRanges() []qdb.KeyRange {

	var ret []qdb.KeyRange

	for _, kr := range qr.Ranges {
		ret = append(ret, qdb.KeyRange{
			KeyRangeID: kr.KeyRangeID,
			ShardID:    kr.ShardID,
			To:         kr.To,
			From:       kr.From,
		})
	}

	return ret
}

func (qr *ShardQrouter) AddShardingColumn(col string) error {
	qr.ColumnMapping[col] = struct{}{}
	return nil
}

func (qr *ShardQrouter) AddLocalTable(tname string) error {
	qr.LocalTables[tname] = struct{}{}
	return nil
}

func (qr *ShardQrouter) AddKeyRange(kr qdb.KeyRange) error {
	if _, ok := qr.Ranges[kr.KeyRangeID]; ok {
		return errors.Errorf("key range with ID already defined", kr.KeyRangeID)
	}

	qr.Ranges[kr.KeyRangeID] = kr
	return nil
}

func (qr *ShardQrouter) routeByIndx(i int) qdb.KeyRange {

	for _, kr := range qr.Ranges {
		if kr.From <= i && kr.To >= i {
			return kr
		}
	}

	return qdb.KeyRange{
		ShardID: NOSHARD,
	}
}

func (qr *ShardQrouter) matchShkey(expr sqlparser.Expr) bool {
	switch texpr := expr.(type) {
	case sqlparser.ValTuple:
		for _, val := range texpr {
			if qr.matchShkey(val) {
				return true
			}
		}
	case *sqlparser.ColName:
		_, ok := qr.ColumnMapping[texpr.Name.String()]
		return ok
	default:
	}

	return false
}

func (qr *ShardQrouter) routeByExpr(expr sqlparser.Expr) ShardRoute {
	switch texpr := expr.(type) {
	case *sqlparser.AndExpr:
		lft := qr.routeByExpr(texpr.Left)
		if lft.Shkey.Name == NOSHARD {
			return qr.routeByExpr(texpr.Right)
		}
		return lft
	case *sqlparser.ComparisonExpr:
		if qr.matchShkey(texpr.Left) {
			shindx := qr.routeByExpr(texpr.Right)
			return shindx
		}
	case *sqlparser.SQLVal:
		valInt, err := strconv.Atoi(string(texpr.Val))
		if err != nil {
			return ShardRoute{
				Shkey: qdb.ShardKey{
					Name: NOSHARD,
				},
			}
		}
		kr := qr.routeByIndx(valInt)
		rw := qr.qdb.Check(kr)

		return ShardRoute{Shkey: qdb.ShardKey{
			Name: kr.ShardID,
			RW:   rw,
		},
			Matchedkr: kr,
		}
	default:
		//tracelog.InfoLogger.Println("typ is %T\n", expr)
	}

	return ShardRoute{
		Shkey: qdb.ShardKey{
			Name: NOSHARD,
		},
	}
}

func (qr *ShardQrouter) isLocalTbl(from sqlparser.TableExprs) bool {
	for _, texpr := range from {
		switch tbltype := texpr.(type) {
		case *sqlparser.ParenTableExpr:
		case *sqlparser.JoinTableExpr:
		case *sqlparser.AliasedTableExpr:

			switch tname := tbltype.Expr.(type) {
			case sqlparser.TableName:
				if _, ok := qr.LocalTables[tname.Name.String()]; ok {
					return true
				}
			case *sqlparser.Subquery:
			default:
			}
		}

	}
	return false
}

func (qr *ShardQrouter) matchShards(sql string) []ShardRoute {
	parsedStmt, err := sqlparser.Parse(sql)
	if err != nil {
		return nil
	}

	tracelog.InfoLogger.Printf("parsed qtype %T", parsedStmt)

	switch stmt := parsedStmt.(type) {
	case *sqlparser.Select:
		if qr.isLocalTbl(stmt.From) {
			return nil
		}
		if stmt.Where != nil {
			shroute := qr.routeByExpr(stmt.Where.Expr)
			return []ShardRoute{shroute}
		}
		return nil

	case *sqlparser.Insert:
		for i, c := range stmt.Columns {

			if _, ok := qr.ColumnMapping[c.String()]; ok {

				switch vals := stmt.Rows.(type) {
				case sqlparser.Values:
					valTyp := vals[0]
					shroute := qr.routeByExpr(valTyp[i])
					return []ShardRoute{shroute}
				}
			}
		}
	case *sqlparser.Update:
		if stmt.Where != nil {
			shroute := qr.routeByExpr(stmt.Where.Expr)
			return []ShardRoute{shroute}
		}
		return nil
	case *sqlparser.CreateTable:
		tracelog.InfoLogger.Printf("ddl routing excpands to every shard")
		// route ddl to every shard
		shrds := qr.Shards()
		var ret []ShardRoute
		for _, sh := range shrds {
			ret = append(ret,
				ShardRoute{Shkey: qdb.ShardKey{
					Name: sh,
					RW:   true,
				},
				})
		}

		return ret
	}

	return nil
}

func (qr *ShardQrouter) Route(q string) []ShardRoute {
	return qr.matchShards(q)
}
