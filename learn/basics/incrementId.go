//分布式id递增

package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"sync"
	"time"
)

func init() {
	orm.RegisterModel(new(IdInfo))
}

//模块id信息
type IdInfo struct {
	Id       int    //模块id
	Max      uint64 //最大值
	Interval uint   //区间
}

//id生成器
type IdGenerator struct {
	moduleId      int         //模块id
	curr          uint64      //当前值
	max           uint64      //最大值
	thresholdRate uint8       //阈值百分比(1-100)，用于计算threshold
	threshold     uint64      //阈值，curr等于阈值时则开始获取下一区间值
	dbAlia        string      //数据库别名
	nextChan      chan IdInfo //下一区间信息通道
	sync.Mutex
}

func NewIdGenerator(moduleId int, dbAlia string) (*IdGenerator, error) {
	ig := &IdGenerator{
		moduleId:      moduleId,
		dbAlia:        dbAlia,
		thresholdRate: 20,
		nextChan:      make(chan IdInfo, 1),
	}
	err := ig.getNext()
	if err != nil {
		//logger.Error(err)
		return nil, err
	}
	return ig, nil
}

//获取自增id
func (ig *IdGenerator) GenId() uint64 {
	ig.Lock()
	defer ig.Unlock()
	ig.curr++
	if ig.curr > ig.max {
		ig.next()
		ig.curr++
	}

	if ig.curr > ig.max {
		//logger.Error("gen error")
		return 0
	}

	if ig.curr == ig.threshold {
		ig.asyncGetNext()
	}

	return ig.curr

}

//跳到下一区间
func (ig *IdGenerator) next() {
	select {
	case idInfo, ok := <-ig.nextChan:
		if ok {
			ig.curr = idInfo.Max - uint64(idInfo.Interval)
			ig.max = idInfo.Max
			ig.threshold = ig.curr + uint64(ig.thresholdRate)*uint64(idInfo.Interval)/100
		}
	}
}

//异步获取下一区间值
func (ig *IdGenerator) asyncGetNext() {
	go func() {
		retryTimes := time.Duration(1)
		for {
			err := ig.getNext()
			//失败重试
			if err != nil {
				time.Sleep(retryTimes * time.Second)
				retryTimes++
				ig.Lock()
				ig.nextChan = make(chan IdInfo, 1)
				ig.Unlock()
				continue
			}
			break
		}
	}()
}

//获取下一区间值
func (ig *IdGenerator) getNext() (err error) {
	idInfo := &IdInfo{Id: ig.moduleId}
	o := orm.NewOrm()
	o.Using(ig.dbAlia)
	o.Begin()
	defer func() {
		if x := recover(); x != nil {
			//logger.Error(x)
			err = fmt.Errorf("%v", x)
		}
		if err != nil {
			//logger.Error(err)
			o.Rollback()
			close(ig.nextChan)
		} else {
			o.Commit()
			ig.nextChan <- *idInfo
		}
	}()
	err = o.ReadForUpdate(idInfo)
	if err != nil {
		return
	}
	idInfo.Max += uint64(idInfo.Interval)
	_, err = o.Update(idInfo)
	if err != nil {
		return
	}
	return
}
