# Bookstore

### Task

- [ ] User
    + [x] Register `localhost:8901/signup`
    + [ ] Login with JWT
    + [ ] Forgot password
    + [ ] Rerender Token
- [ ] Product (book)
    + [x] Create Product `localhost:8901/books`
    + [x] Edit Product `localhost:8901/books`
    + [x] Delete Product `localhost:8901/books`
    + [x] Get product by ID `localhost:8901/books/{id}`
    + [x] Get All Product `localhost:8901/books`
- [ ] Shopping Cart
    + [x] Add product to Shopping Cart `localhost:8901/cart/items`
    + [ ] Remove product from Shopping Cart
    + [x] Get info Shopping Cart `localhost:8901/carts`
    + [ ] Payment Shopping Cart **[50%] Using Stripe** `localhost:8901/payment`
- [ ] Test case

### Setup environment

- run docker file:

```bash
$ docker-compose up -d
```

### Information

- Website: awesomebook
- Database: MySQL
    - database name: awesomebook
    - username: root
    - password: 123456Aa@

### ERD

![ERD](https://github.com/thanbv1510/tfs-03/blob/master/lec-06/awesomebook/resources/ERD.png)