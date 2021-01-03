package window

import (
	"sort"
	"time"
)

type RollingNumber struct {
	timeInMilliseconds int
	numberOfBuckets    int
	buckets            *CycleQueue
}

func NewRollingNumber(timeInMilliseconds, numberOfBuckets int) *RollingNumber {
	return &RollingNumber{
		timeInMilliseconds: timeInMilliseconds,
		numberOfBuckets:    numberOfBuckets,
		buckets:            NewCycleQueue(numberOfBuckets),
	}
}

/// 当前Bucket 自加1
func (r *RollingNumber) Increment(event Event) {
	r.GetCurrentBucket().Increment(event)
}

/// 当前Bucket 加上指定值
func (r *RollingNumber) Add(event Event, value int) {
	r.GetCurrentBucket().Add(event, value)
}

// 更新当前maxUpdater，保留最大值
func (r *RollingNumber) UpdateRollingMax(event Event, value int) {
	r.GetCurrentBucket().UpdateMaxUpdater(event, value)
}

// 清空数据
func (r *RollingNumber) Reset() {

}

//根据event type 获取所有Bucket 某index 总和
func (r *RollingNumber) GetRollingSum(event Event) int {
	if r.GetCurrentBucket() == nil {
		return 0
	}

	sum := 0
	for _, b := range r.buckets.data {
		bucket := b.(*Bucket)
		sum += bucket.GetAdder(event)
	}
	return sum

}

// 获取最后一个bucket 值
func (r *RollingNumber) GetValueOfLatestBucket(event Event) int {

	return r.buckets.getLast().(*Bucket).GetAdder(event)
}

// 获取所有bucket 某一个索引的所有值
func (r *RollingNumber) GetValues(event Event) []int {

	result := make([]int, r.buckets.curSize())
	for idx, b := range r.buckets.data {
		bucket := b.(*Bucket)
		result[idx] += bucket.GetAdder(event)
	}
	return result
}

// getValues 结果的最大值
func (r *RollingNumber) GetRollingMaxValue(event Event) int {

	result := r.GetValues(event)
	sort.Ints(result)
	return result[len(result)-1]
}

// 获取当前bucket
func (r *RollingNumber) GetCurrentBucket() *Bucket {
	var bucket *Bucket
	if r.buckets.getLast() == nil {
		// 如果为空，重新生成一个
		currentTime := time.Now().Unix()
		bucket = NewBucket(currentTime)
		r.buckets.Push(bucket)
	} else {
		bucket = r.buckets.getLast().(*Bucket)
	}
	return bucket
}
