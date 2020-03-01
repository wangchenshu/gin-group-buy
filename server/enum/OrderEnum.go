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

const (
	TEST_BANK_NUM                    = "007"
	TEST_BANK_ACCOUNT                = "001234567899999"
	IMG_URL_OPTION_GROUP_BUY_PRODUCT = "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E6%B5%B7%E8%8B%94%E7%A6%AE%E7%9B%92.jpg?alt=media&token=4e1e859f-fae6-41de-86f4-94a506c3a2a9"
	IMG_URL_OPTION_CHECK_CART        = "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b"
	IMG_URL_OPTION_CHECKOUT          = "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b"
	IMG_URL_OPTION_CLEAR_CART        = "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b"
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
