package database

const createProductsTable = `CREATE TABLE IF NOT EXISTS products 
    (
                                        id VARBINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
                                        code VARCHAR(4) NOT NULL,
                                        name VARCHAR(50) NOT NULL,
                                        price DECIMAL(10,4) NOT NULL DEFAULT 10,
                                        stock INT NOT NULL DEFAULT 10,
                                        short_description VARCHAR(100) NOT NULL,
                                        long_description VARCHAR(255) NOT NULL,
                                        created_at DATETIME NOT NULL DEFAULT NOW(),
                                        updated_at DATETIME NOT NULL DEFAULT NOW() ON UPDATE now(),
                                        PRIMARY KEY (id),
                                        unique key (code)
    );`

const createUserTable = `CREATE TABLE IF NOT EXISTS users
    (
                                    id        VARCHAR(16) NOT NULL,
                                    firstName VARCHAR(50) NOT NULL,
                                    lastName  VARCHAR(50) NOT NULL,
                                    username  VARCHAR(50) NOT NULL,
                                    password  VARCHAR(250) NOT NULL,
                                    createdAt DATETIME NOT NULL DEFAULT NOW(),
                                    updatedAt DATETIME NOT NULL DEFAULT NOW() ON UPDATE now(),
                                    primary key (ID),
                                    UNIQUE KEY(username)
    );`

const createSessionTable = `CREATE TABLE IF NOT EXISTS sessions
    (
                                    id         VARCHAR(100) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
                                    userId     VARCHAR(50) NOT NULL,
                                    expiresAt  DATETIME NOT NULL,
                                    primary key (id),
                                    UNIQUE KEY(userId)
    );`

const createCartItemTable = `CREATE TABLE IF NOT EXISTS carItems
    (
                                    id         VARCHAR(100) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
                                    userId     VARCHAR(50) NOT NULL,
                                    expiresAt  DATETIME NOT NULL,
                                    primary key (id),
                                    UNIQUE KEY(userId)
    );`
