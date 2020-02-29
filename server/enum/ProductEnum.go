package enum

type ProductEnum int

const (
	MENU           = iota // 選單
	INPUT_KEYWORDS        // 快速選單或輸入商品關鍵字
)

func (action ProductEnum) String() string {
	return [...]string{"選單", "快速選單或輸入商品關鍵字"}[action]
}
