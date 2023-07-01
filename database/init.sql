
CREATE TABLE IF NOT EXISTS Products (
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
);

