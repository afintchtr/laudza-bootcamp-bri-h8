# laudza-bootcamp-bri-h8
 Golang and JavaScript Path

## Penjelasan Assignment Golang 1
Berikut adalah daftar endpoint yang ada pada service tersebut.
- Create order\
  
- Fetch all orders
- Fetch single order by Id
- Fetch all items
- Fetch single item by Id
- Update single order
- Delete single order

```
r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.FetchAllOrderWithItems)
	r.GET("/orders/:id", controllers.FetchOrderWithItemsById)
	r.GET("/orders/items/", controllers.FetchAllItems)
	r.GET("/orders/items/:id", controllers.FetchItemById)
	r.PATCH("/orders/:id", controllers.UpdateOrderById)
	r.DELETE("/orders/:id", controllers.DeleteOrderById)
```
