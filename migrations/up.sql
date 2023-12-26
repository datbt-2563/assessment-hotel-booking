CREATE TABLE customer (
    customer_id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(10) NOT NULL,
    email VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    PRIMARY KEY (customer_id)
);

CREATE TABLE room (
    room_id INT NOT NULL AUTO_INCREMENT,
    room_type VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    max_occupancy INT NOT NULL,
    status VARCHAR(255) NOT NULL,
    PRIMARY KEY (room_id)
);

CREATE TABLE reservation (
    reservation_id INT NOT NULL AUTO_INCREMENT,
    customer_id INT NOT NULL,
    room_id INT NOT NULL,
    check_in_date DATE NOT NULL,
    check_out_date DATE NOT NULL,
    occupancy INT NOT NULL,
    total_price INT NOT NULL,
    status  VARCHAR(255) NOT NULL,
    PRIMARY KEY (reservation_id),
    FOREIGN KEY (customer_id) REFERENCES customer (customer_id),
    FOREIGN KEY (room_id) REFERENCES room (room_id)
);
