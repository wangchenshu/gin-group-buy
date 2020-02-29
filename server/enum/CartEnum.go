package enum

type CartEnum int

const (
	CHECK_CART         = iota // 查看購物車
	CLEAR_CART                // 清除購物車
	CHECKOUT                  // 結帳
	EMPTY_CART                // 購物車無商品
	ADD_TO_CART               // 加入購物車
	ADD_CART_SUCCESS          // 加入購物車成功
	ADD_CART_FAIL             // 加入購物車失敗
	CLEAR_CART_SUCCESS        // 清除購物車成功
	CLEAR_CART_FAIL           // 清除購物車失敗
	CURRENT_CART              // 目前購物車有
)

func (action CartEnum) String() string {
	return [...]string{
		"查看購物車",
		"清除購物車",
		"結帳",
		"購物車內無商品",
		"加入購物車",
		"加入購物車成功",
		"加入購物車失敗",
		"清除購物車成功",
		"清除購物車失敗",
		"目前購物車有",
	}[action]
}
