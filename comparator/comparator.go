package comparator

// 键值比较器
type Comparer interface {
	// 返回值可以是`[小于0, 0, 大于0]`表示`[小于, 等于, 大于]`的关系。
	Compare(a, b []byte) int
	// 比较器的名字，以"leveldb."起始是保留字段，不应该使用
	Name() string
	// 为了节省存储空间，比如在IndexBlock的Key
	// 要求返回一个最短字节序列s，其大小在[start，limit)之间，且如果s==start则应该返回nil
	// 也就是s长度不应该大于start，但是大小关系要大于等于start
	FindShortestSeparator(start, limit []byte) []byte
	// 为了节省存储空间，比如在IndexBlock的Key
	// 要求返回一个最短字节序列s，其大小在[start, +∞)之间，且如果s==start则应该返回nil
	// 也就是s长度不应该大于start，但是大小关系要大于等于start
	FindShortSuccessor(start []byte) []byte
}
