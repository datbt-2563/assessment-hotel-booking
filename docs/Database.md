# Database

## Customer
| Field          | Type           | Description  |
| :------------- | :------------- | :----------- |
| `customer_id`  | `int`          | Customer ID  |
| `name`         | `varchar(255)` | Name         |
| `phone_number` | `varchar(10)`  | Phone number |
| `email`        | `varchar(255)` | Email        |
| `address`      | `varchar(255)` | Address      |

## Room

| Field           | Type           | Description                                                  |
| :-------------- | :------------- | :----------------------------------------------------------- |
| `room_id`       | `int`          | Room ID                                                      |
| `room_type`     | `varchar(255)` | Room type                                                    |
| `price`         | `int`          | Price                                                        |
| `max_occupancy` | `int`          | Max occupancy                                                |
| `status`        | `varchar(255)` | Status<br />Available, Occupied, Cleaning, Maintenance<br />Reserved: booked but empty now<br/>Hold: payment is being made<br />Canceled: Payment failed or canceled. |

## Reservation
| Field            | Type         | Description    |
| :--------------- | :----------- | :------------- |
| `reservation_id` | `int`        | Reservation ID |
| `customer_id`    | `int`        | Customer ID    |
| `room_id`        | `int`        | Room ID        |
| `check_in_date`  | `datetime`   | Check-in date  |
| `check_out_date` | `datetime`   | Check-out date |
| `occupancy`      | `int`        | Occupancy      |
| `total_price`    | `int`        | Total price    |
| status           | varchar(255) | Status         |
