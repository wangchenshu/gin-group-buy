package enum

type OrderEnum int

// 狀態類別
const (
	GROUP_BUY_PRODUCT   = iota // 團購商品
	CHECKOUT_SUCCESS           // 結帳完成
	CHECKOUT_FAIL              // 結帳失敗
	PRICE                      // 價格
	QTY                        // 數量
	TOTAL_PRICE                // 總計
	MONEY_TRANSFER             // 匯款
	MONEY_TRANSFER_TIPS        //
	TRANSFER_BANK_NUM          // 銀行代號
	TRANSFER_ACCOUNT           // 帳戶
	TRANSFER_AMOUNT            // 匯款金額
	ORDER_DETAIL               // 明細如下
)

func (action OrderEnum) String() string {
	return [...]string{
		"團購商品",
		"結帳完成",
		"結帳失敗",
		"價格",
		"數量",
		"總計",
		"匯款",
		"請於2日內，匯款至以下指定帳戶，我們收到後會儘快為您出貨，謝謝您的配合。",
		"銀行代號",
		"帳戶",
		"匯款金額",
		"明細如下",
	}[action]
}
