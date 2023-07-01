
CREATE TABLE IF NOT EXISTS Products (
                                        id VARBINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
                                        name VARCHAR(50) NOT NULL,
                                        price DECIMAL(10,4) NOT NULL DEFAULT 10,
                                        stock INT NOT NULL DEFAULT 10,
                                        shortDescription VARCHAR(100) NOT NULL,
                                        longDescription VARCHAR(255) NOT NULL,
                                        createdAt DATETIME NOT NULL DEFAULT NOW(),
                                        updatedAt DATETIME NOT NULL DEFAULT NOW() ON UPDATE now(),
                                        PRIMARY KEY (id)
);

