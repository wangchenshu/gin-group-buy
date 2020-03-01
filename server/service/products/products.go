package products

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"

	"net/http"

	"gin-group-buy/server/db"
	"gin-group-buy/server/enum"
	"gin-group-buy/server/model"
	"gin-group-buy/server/service/mylinebot"
)

var myBot = mylinebot.Init()

const (
	ADD_TO_CART        enum.CartEnum = enum.ADD_TO_CART
	ADD_CART_SUCCESS   enum.CartEnum = enum.ADD_CART_SUCCESS
	ADD_CART_FAIL      enum.CartEnum = enum.ADD_CART_FAIL
	EMPTY_CART         enum.CartEnum = enum.EMPTY_CART
	CHECK_CART         enum.CartEnum = enum.CHECK_CART
	CLEAR_CART         enum.CartEnum = enum.CLEAR_CART
	CLEAR_CART_SUCCESS enum.CartEnum = enum.CLEAR_CART_SUCCESS
	CURRENT_CART       enum.CartEnum = enum.CURRENT_CART
	CHECKOUT           enum.CartEnum = enum.CHECKOUT

	INPUT_KEYWORDS enum.ProductEnum = enum.INPUT_KEYWORDS

	GROUP_BUY_PRODUCT   enum.OrderEnum = enum.GROUP_BUY_PRODUCT
	PRICE               enum.OrderEnum = enum.PRICE
	QTY                 enum.OrderEnum = enum.QTY
	ORDER_DETAIL        enum.OrderEnum = enum.ORDER_DETAIL
	TRANSFER_BANK_NUM   enum.OrderEnum = enum.TRANSFER_BANK_NUM
	TRANSFER_ACCOUNT    enum.OrderEnum = enum.TRANSFER_ACCOUNT
	TRANSFER_AMOUNT     enum.OrderEnum = enum.TRANSFER_AMOUNT
	MONEY_TRANSFER      enum.OrderEnum = enum.MONEY_TRANSFER
	MONEY_TRANSFER_TIPS enum.OrderEnum = enum.MONEY_TRANSFER_TIPS
	TOTAL_PRICE         enum.OrderEnum = enum.TOTAL_PRICE

	TEST_BANK_NUM     = "007"
	TEST_BANK_ACCOUNT = "001234567899999"
	IMG_URL_OPTION_1  = "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E6%B5%B7%E8%8B%94%E7%A6%AE%E7%9B%92.jpg?alt=media&token=4e1e859f-fae6-41de-86f4-94a506c3a2a9"
	IMG_URL_OPTION_2  = "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b"
	IMG_URL_OPTION_3  = "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b"
	IMG_URL_OPTION_4  = "https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b"
)

// PostHandler -
func PostHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		events, err := myBot.ParseRequest(context.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				context.JSON(http.StatusBadRequest, nil)
			} else {
				context.JSON(http.StatusInternalServerError, nil)
			}
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				profile, err := myBot.GetProfile(event.Source.UserID).Do()
				if err != nil {
					log.Print(err)
				}

				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					matchText := message.Text
					buyProductName := ""

					if strings.Contains(matchText, ADD_TO_CART.String()) {
						buyProductName = strings.Split(matchText, ",")[1]
						matchText = strings.Split(matchText, ",")[0]
					}

					switch matchText {
					case ADD_TO_CART.String():
						replyMsg := ADD_CART_FAIL.String()

						if addCarts(profile.DisplayName, buyProductName, 1) {
							replyMsg = " " + buyProductName + " " + ADD_CART_SUCCESS.String()
						}

						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage(profile.DisplayName+replyMsg),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case GROUP_BUY_PRODUCT.String():
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							myMenuTemplate(),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case CHECK_CART.String():
						carts := getCartsByUsername(profile.DisplayName)
						repMsg := CURRENT_CART.String() + ": \n\n"
						cartIsEmpty := EMPTY_CART.String()
						totalPrice := 0

						for _, cart := range carts {
							cartQtyStr := fmt.Sprintf("%d", cart.Qty)
							repMsg += cart.ProductName + ", "
							repMsg += PRICE.String() + ": $" + fmt.Sprintf("%d", cart.Price) + ", "
							repMsg += QTY.String() + ": " + cartQtyStr + "\n"
							totalPrice += cart.Price * cart.Qty
						}
						repMsg += "\n" + TOTAL_PRICE.String() + ": $ " + fmt.Sprintf("%d", totalPrice) + "\n"

						if totalPrice <= 0 {
							repMsg = cartIsEmpty
						}

						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage(repMsg),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case CLEAR_CART.String():
						if clearCarts(profile.DisplayName) {
							repMsg := CLEAR_CART_SUCCESS.String()

							if _, err := myBot.ReplyMessage(
								event.ReplyToken,
								linebot.NewTextMessage(repMsg),
								myQuickReply(),
							).Do(); err != nil {
								log.Print(err)
							}
						}
					case CHECKOUT.String():
						carts := getCartsByUsername(profile.DisplayName)
						repMsg := ORDER_DETAIL.String() + ": \n\n"
						cartIsEmpty := EMPTY_CART.String()
						totalPrice := 0
						testBankNum := TEST_BANK_NUM
						testAccount := TEST_BANK_ACCOUNT

						for _, cart := range carts {
							cartQtyStr := fmt.Sprintf("%d", cart.Qty)
							repMsg += cart.ProductName + ", "
							repMsg += PRICE.String() + ": $" + fmt.Sprintf("%d", cart.Price) + ", "
							repMsg += QTY.String() + ": " + cartQtyStr + "\n"
							totalPrice += cart.Price * cart.Qty
						}

						if totalPrice <= 0 {
							repMsg = cartIsEmpty
						} else {
							totalPriceStr := fmt.Sprintf("%d", totalPrice)
							repMsg += "\n" + TOTAL_PRICE.String() + ": $ " + totalPriceStr + "\n\n"
							repMsg += TRANSFER_BANK_NUM.String() + ": " + testBankNum + "\n"
							repMsg += TRANSFER_ACCOUNT.String() + ": " + testAccount + "\n"
							repMsg += TRANSFER_AMOUNT.String() + ": $ " + totalPriceStr + "\n"
							repMsg += "\n" + MONEY_TRANSFER_TIPS.String() + "\n"
						}

						clearCarts(profile.DisplayName)

						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage(repMsg),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					default:
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							myMenuTemplate(),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					}
				case *linebot.ImageMessage:
					log.Print(message)
				case *linebot.VideoMessage:
					log.Print(message)
				case *linebot.AudioMessage:
					log.Print(message)
				case *linebot.FileMessage:
					log.Print(message)
				case *linebot.LocationMessage:
					log.Print(message)
				case *linebot.StickerMessage:
					log.Print(message)
				default:
					log.Printf("Unknown message: %v", message)
				}
			default:
				log.Printf("Unknown event: %v", event)
			}
		}
		context.JSON(http.StatusOK, gin.H{
			"success": events,
		})
	}
}

func newCarouselColumn(imageURL, title, text string, actionLabel string, actionText string) *linebot.CarouselColumn {
	return linebot.NewCarouselColumn(
		imageURL, title, text,
		linebot.NewMessageAction(actionLabel, actionText),
	)
}

func getAllProducts() []model.Product {
	products := []model.Product{}
	db.Db.Where("active = ?", "1").Find(&products)
	return products
}

func getProductsLike(name string, limit int) []model.Product {
	products := []model.Product{}
	db.Db.Where("name LIKE ?", "%"+name+"%").Limit(limit).Find(&products)
	return products
}

func getProductLike(name string) model.Product {
	product := model.Product{}
	db.Db.Where("name LIKE ?", "%"+name+"%").First(&product)
	return product
}

func newMessageAction(label string, text string) *linebot.MessageAction {
	return linebot.NewMessageAction(label, text)
}

func myMenuTemplate() *linebot.TemplateMessage {
	products := getAllProducts()
	wannaBuyStr := ADD_TO_CART.String()
	altText := GROUP_BUY_PRODUCT.String()
	arouselColumns := []*linebot.CarouselColumn{}

	for _, product := range products {
		actionLabel := fmt.Sprintf("$ %d", product.Price)
		arouselColumns = append(arouselColumns,
			newCarouselColumn(
				product.PicUrl,
				product.Name,
				actionLabel,
				wannaBuyStr,
				ADD_TO_CART.String()+","+product.Name,
			),
		)
	}

	template := linebot.NewCarouselTemplate(arouselColumns...)
	templateMessage := linebot.NewTemplateMessage(altText, template)

	return templateMessage
}

func myQuickReply() linebot.SendingMessage {
	content := INPUT_KEYWORDS.String()
	imageURLs := []string{
		IMG_URL_OPTION_1,
		IMG_URL_OPTION_2,
		IMG_URL_OPTION_3,
		IMG_URL_OPTION_4,
	}
	labels := []string{
		GROUP_BUY_PRODUCT.String(),
		CHECK_CART.String(),
		CHECKOUT.String(),
		CLEAR_CART.String(),
	}
	quickReplyButtons := []*linebot.QuickReplyButton{}

	for k, v := range labels {
		quickReplyButtons = append(quickReplyButtons, linebot.NewQuickReplyButton(
			imageURLs[k], linebot.NewMessageAction(v, v),
		))
	}
	quickReply := linebot.NewTextMessage(content).
		WithQuickReplies(linebot.NewQuickReplyItems(quickReplyButtons...))

	return quickReply
}

func addCarts(username string, productName string, qty int) bool {
	carts := model.Cart{}
	product := getProductLike(productName)

	if product.ID != 0 {
		carts.Username = username
		carts.ProductName = product.Name
		carts.Qty = qty
		carts.Price = product.Price

		db.Db.Create(&carts)
		return true
	}

	return false
}

func getCartsByUsername(username string) []model.Cart {
	carts := []model.Cart{}
	db.Db.Where("username = ?", username).Find(&carts)

	return carts
}

func clearCarts(username string) bool {
	db.Db.Delete(model.Cart{}, "username = ?", username)
	return true
}

// GetProducts -
func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, getAllProducts())
	}
}

// GetProductsLike -
func GetProductsLike(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, getProductsLike(name, 100))
	}
}
