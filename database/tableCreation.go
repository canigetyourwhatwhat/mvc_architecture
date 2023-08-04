package database

const createProductsTable = `CREATE TABLE IF NOT EXISTS products (
    id                       BIGINT AUTO_INCREMENT NOT NULL,
    code                     VARCHAR(4) NOT NULL,
    name                     VARCHAR(50) NOT NULL,
    price                    DECIMAL(10,4) NOT NULL DEFAULT 10,
    stock                    INT NOT NULL DEFAULT 10,
    short_description        VARCHAR(100) NOT NULL,
    long_description         VARCHAR(255) NOT NULL,
    created_at               DATETIME NOT NULL DEFAULT NOW(),
    updated_at               DATETIME NOT NULL DEFAULT NOW() ON UPDATE now(),
                             PRIMARY KEY (id),
                             UNIQUE KEY (code)
);`

const createUserTable = `CREATE TABLE IF NOT EXISTS users ( 
    id            BIGINT AUTO_INCREMENT NOT NULL,
    firstName     VARCHAR(50)   NOT NULL,
    lastName      VARCHAR(50)   NOT NULL,
    username      VARCHAR(50)   NOT NULL,
    password      VARCHAR(250)  NOT NULL,
    createdAt     DATETIME      NOT NULL DEFAULT NOW(),
    updatedAt     DATETIME      NOT NULL DEFAULT NOW() ON UPDATE now(),
                  PRIMARY KEY (ID),
                  UNIQUE KEY(username)
);`

const createSessionTable = `CREATE TABLE IF NOT EXISTS sessions (
    id          VARCHAR(100) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
    userId      BIGINT  NOT NULL,
    expiresAt   DATETIME     NOT NULL,
                PRIMARY KEY (id),
                UNIQUE KEY(userId),
                CONSTRAINT session_to_user foreign key (userId) references users(id)
                
);`

const createCartItemTable = `CREATE TABLE IF NOT EXISTS cartItems (
    id           BIGINT AUTO_INCREMENT NOT NULL,
    productCode  VARCHAR(4)     NOT NULL,
    cartId       BIGINT         NOT NULL,
    quantity     INT            NOT NULL,
    totalPrice   DECIMAL(10, 4) NOT NULL,
    taxPrice     DECIMAL(10, 4) NOT NULL,
    netPrice     DECIMAL(10, 4) NOT NULL,
    createdAt    DATETIME       NOT NULL DEFAULT NOW(),
    updatedAt    DATETIME       NOT NULL DEFAULT NOW() ON UPDATE NOW(),
                 PRIMARY KEY (id),
                 CONSTRAINT cart_item_to_cart foreign key (cartId) references carts(id)
);`

const createCartTable = `CREATE TABLE IF NOT EXISTS carts (
    id         BIGINT          NOT NULL AUTO_INCREMENT,
    userId     BIGINT     NOT NULL,
    totalPrice DECIMAL(10, 4)  NOT NULL DEFAULT 0,
    taxPrice   DECIMAL(10, 4)  NOT NULL DEFAULT 0,
    netPrice   DECIMAL(10, 4)  NOT NULL DEFAULT 0,
    status     INT             NOT NULL DEFAULT 0,
    createdAt  DATETIME        NOT NULL DEFAULT NOW(),
    updatedAt  DATETIME        NOT NULL DEFAULT NOW() ON UPDATE NOW(),
               PRIMARY KEY (id),
               CONSTRAINT cart_to_user foreign key (userId) references users(id)
);`

const createPaymentTable = `CREATE TABLE IF NOT EXISTS payments (
                                     id         BIGINT          NOT NULL AUTO_INCREMENT,
                                     amount     DECIMAL(10, 4)  NOT NULL,
                                     userId     BIGINT          NOT NULL,
                                     method     INT             NOT NULL DEFAULT 0,
                                     createdAt  DATETIME        NOT NULL DEFAULT NOW(),
                                     updatedAt  DATETIME        NOT NULL DEFAULT NOW() ON UPDATE NOW(),
                                     PRIMARY KEY (id),
                                     CONSTRAINT payment_to_user foreign key (userId) references users(id)
);`

const createOrderTable = `CREATE TABLE IF NOT EXISTS orders (
                                        id         bigint           NOT NULL AUTO_INCREMENT,
                                        userId     BIGINT           NOT NULL,
                                        cartId     bigint           NOT NULL,
                                        paymentId  bigint           NOT NULL,
                                        createdAt  DATETIME         NOT NULL DEFAULT NOW(),
                                        PRIMARY KEY (id),
                                        CONSTRAINT order_to_user    foreign key (userId) references users(id),
                                        CONSTRAINT order_to_cart    foreign key (cartId) references carts(id),
                                        CONSTRAINT order_to_payment foreign key (paymentId) references payments(id)
);`
