**_ Info _**

https://www.youtube.com/watch?v=a6G5-LUlFO4&t=3228&ab_channel=AkhilSharma

protoc --go_out=. --go-grpc_out=. proto/greet.proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

**_ Create User _**

CREATE TABLE User (
id INT AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL,
password VARCHAR(255) NOT NULL,
role VARCHAR(255) NOT NULL,
is_active BOOLEAN NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL
);

**_ Dump Users _**

INSERT INTO User (name, email, password, role, is_active, created_at, updated_at) VALUES
('John Doe', 'john@example.com', 'password1', 'admin', true, '2024-03-26 09:00:00', '2024-03-26 09:00:00'),
('Jane Smith', 'jane@example.com', 'password2', 'user', true, '2024-03-26 10:00:00', '2024-03-26 10:00:00'),
('Alice Johnson', 'alice@example.com', 'password3', 'user', true, '2024-03-26 11:00:00', '2024-03-26 11:00:00'),
('Bob Brown', 'bob@example.com', 'password4', 'user', true, '2024-03-26 12:00:00', '2024-03-26 12:00:00'),
('Eve Wilson', 'eve@example.com', 'password5', 'user', true, '2024-03-26 13:00:00', '2024-03-26 13:00:00'),
('Michael Lee', 'michael@example.com', 'password6', 'user', true, '2024-03-26 14:00:00', '2024-03-26 14:00:00'),
('Emily Davis', 'emily@example.com', 'password7', 'user', true, '2024-03-26 15:00:00', '2024-03-26 15:00:00'),
('David Rodriguez', 'david@example.com', 'password8', 'user', true, '2024-03-26 16:00:00', '2024-03-26 16:00:00'),
('Sarah Martinez', 'sarah@example.com', 'password9', 'user', true, '2024-03-26 17:00:00', '2024-03-26 17:00:00'),
('Ryan Hernandez', 'ryan@example.com', 'password10', 'user', true, '2024-03-26 18:00:00', '2024-03-26 18:00:00'),
('Olivia Walker', 'olivia@example.com', 'password11', 'user', true, '2024-03-26 19:00:00', '2024-03-26 19:00:00'),
('Daniel Young', 'daniel@example.com', 'password12', 'user', true, '2024-03-26 20:00:00', '2024-03-26 20:00:00'),
('Sophia Allen', 'sophia@example.com', 'password13', 'user', true, '2024-03-26 21:00:00', '2024-03-26 21:00:00'),
('Matthew Scott', 'matthew@example.com', 'password14', 'user', true, '2024-03-26 22:00:00', '2024-03-26 22:00:00'),
('Emma King', 'emma@example.com', 'password15', 'user', true, '2024-03-26 23:00:00', '2024-03-26 23:00:00'),
('Liam Green', 'liam@example.com', 'password16', 'user', true, '2024-03-27 00:00:00', '2024-03-27 00:00:00'),
('Ava Hall', 'ava@example.com', 'password17', 'user', true, '2024-03-27 01:00:00', '2024-03-27 01:00:00'),
('Noah Adams', 'noah@example.com', 'password18', 'user', true, '2024-03-27 02:00:00', '2024-03-27 02:00:00'),
('Mia Baker', 'mia@example.com', 'password19', 'user', true, '2024-03-27 03:00:00', '2024-03-27 03:00:00'),
('James Cook', 'james@example.com', 'password20', 'user', true, '2024-03-27 04:00:00', '2024-03-27 04:00:00');
