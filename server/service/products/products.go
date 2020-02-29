package products

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"

	"net/http"

	"gin-group-buy/server/db"
	"gin-group-buy/server/model"
	"gin-group-buy/server/service/mylinebot"
)

var myBot = mylinebot.Init()

// PostHandler -
func PostHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		events, err := myBot.ParseRequest(context.Request)
		fmt.Println(events)

		if err != nil {
			if err == linebot.ErrInvalidSignature {
				context.JSON(http.StatusBadRequest, nil)
			} else {
				context.JSON(500, nil)
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
					if strings.Contains(matchText, "+1") {
						buyProductName = strings.Split(matchText, ",")[1]
						matchText = strings.Split(matchText, ",")[0]
					}

					switch matchText {
					case "+1":
						replyMsg := "+1 失敗"
						if addCarts(profile.DisplayName, buyProductName, 1) {
							replyMsg = " " + buyProductName + " +1 成功"
						}
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							linebot.NewTextMessage(profile.DisplayName+replyMsg),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case "團購商品":
						if _, err := myBot.ReplyMessage(
							event.ReplyToken,
							myMenuTemplate(),
							myQuickReply(),
						).Do(); err != nil {
							log.Print(err)
						}
					case "查看購物車":
						carts := getCartsByUsername(profile.DisplayName)
						repMsg := "目前購物車有: \n\n"
						cartIsEmpty := "購物車是空的"
						totalPrice := 0

						for _, cart := range carts {
							cartQtyStr := fmt.Sprintf("%d", cart.Qty)
							repMsg += cart.ProductName + ", 價格: $" + fmt.Sprintf("%d", cart.Price) + ", 數量: " + cartQtyStr + "\n"

							totalPrice += cart.Price * cart.Qty
						}
						repMsg += "\n總計: $ " + fmt.Sprintf("%d", totalPrice) + "\n"

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
					case "清除購物車":
						if clearCarts(profile.DisplayName) {
							repMsg := "清除成功"
							if _, err := myBot.ReplyMessage(
								event.ReplyToken,
								linebot.NewTextMessage(repMsg),
								myQuickReply(),
							).Do(); err != nil {
								log.Print(err)
							}
						}
					case "結帳":
						carts := getCartsByUsername(profile.DisplayName)
						repMsg := "明細如下: \n\n"
						cartIsEmpty := "購物車是空的"
						totalPrice := 0
						testBankNum := "007"
						testAccount := "001234567899999"

						for _, cart := range carts {
							cartQtyStr := fmt.Sprintf("%d", cart.Qty)
							repMsg += cart.ProductName + ", 價格: $" + fmt.Sprintf("%d", cart.Price) + ", 數量: " + cartQtyStr + "\n"

							totalPrice += cart.Price * cart.Qty
						}

						if totalPrice <= 0 {
							repMsg = cartIsEmpty
						} else {
							totalPriceStr := fmt.Sprintf("%d", totalPrice)
							repMsg += "\n總計: $ " + totalPriceStr + "\n\n"
							repMsg += "銀行代號: " + testBankNum + "\n"
							repMsg += "匯款帳戶: " + testAccount + "\n"
							repMsg += "匯款金額: $ " + totalPriceStr + "\n"
							repMsg += "\n請於2日內，匯款至以下指定帳戶，我們收到後會儘快為您出貨，謝謝您的配合。\n"
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
	wannaBuyStr := "我想+1"
	altText := "團購人氣商品"
	arouselColumns := []*linebot.CarouselColumn{}

	for _, product := range products {
		actionLabel := fmt.Sprintf("$ %d", product.Price)
		arouselColumns = append(arouselColumns,
			newCarouselColumn(
				product.PicUrl,
				product.Name,
				actionLabel,
				wannaBuyStr,
				"+1,"+product.Name,
			),
		)
	}
	template := linebot.NewCarouselTemplate(arouselColumns...)
	templateMessage := linebot.NewTemplateMessage(altText, template)

	return templateMessage
}

func myQuickReply() linebot.SendingMessage {
	content := "快速選單或輸入商品關鍵字"
	imageURLs := []string{
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E6%B5%B7%E8%8B%94%E7%A6%AE%E7%9B%92.jpg?alt=media&token=4e1e859f-fae6-41de-86f4-94a506c3a2a9",
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b",
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b",
		"https://firebasestorage.googleapis.com/v0/b/atomy-bot.appspot.com/o/%E8%89%BE%E5%A4%9A%E7%BE%8E%20%E7%89%A9%E7%90%86%E6%80%A7%E9%98%B2%E6%9B%AC%E8%86%8F.jpg?alt=media&token=e659398b-c5a5-4e0e-ae91-614633d2355b",
	}
	labels := []string{"團購商品", "查看購物車", "結帳", "清除購物車"}
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
