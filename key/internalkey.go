package key

// 值类型
type ValueType uint8

const (
	// 表示删除操作
	ValueTypeDeletion ValueType = 0
	// 表示值操作
	ValueTypeValue ValueType = 1
	// 表示查询操作的时候的值类型
	// 因为序列号排序都是降序的，并且值类型被嵌入为序列号的低八位，因此需要最大的值类型（用于搜索时起始Key）
	ValueTypeForSeek = ValueTypeValue
)

// 最大序列号
const MaxSequenceNumber = (uint64(1) << 56) - 1

// 内部键
type InternalKey []byte

func NewInternalKey(userKey []byte, seqNum uint64, valueType ValueType) InternalKey {

}

func (k *InternalKey) UserKey() {

}
