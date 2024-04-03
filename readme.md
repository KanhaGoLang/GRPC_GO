**_ Info _**

https://www.youtube.com/watch?v=a6G5-LUlFO4&t=3228&ab_channel=AkhilSharma

protoc --go_out=. --go-grpc_out=. proto/proto.proto && go mod tidy
protoc --go_out=. --go-grpc_out=. proto/user.proto
protoc --go_out=. --go-grpc_out=. proto/post.proto
protoc --go_out=. --go-grpc_out=. proto/post.proto

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

**_ Create Post Table _**
CREATE TABLE posts (
id INT AUTO_INCREMENT PRIMARY KEY,
title VARCHAR(255),
description TEXT,
is_active BOOLEAN,
user_id INT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO posts (title, description, is_active, user_id, created_at, updated_at)
VALUES
('Title 1', 'Description 1', 1, 22, NOW(), NOW()),
('Title 2', 'Description 2', 0, 23, NOW(), NOW()),
('Title 3', 'Description 3', 1, 24, NOW(), NOW()),
('Title 4', 'Description 4', 0, 25, NOW(), NOW()),
('Title 5', 'Description 5', 1, 22, NOW(), NOW()),
('Title 6', 'Description 6', 0, 23, NOW(), NOW()),
('Title 7', 'Description 7', 1, 24, NOW(), NOW()),
('Title 8', 'Description 8', 0, 25, NOW(), NOW()),
('Title 9', 'Description 9', 1, 22, NOW(), NOW()),
('Title 10', 'Description 10', 0, 23, NOW(), NOW()),
('Title 11', 'Description 11', 1, 24, NOW(), NOW()),
('Title 12', 'Description 12', 0, 25, NOW(), NOW()),
('Title 13', 'Description 13', 1, 22, NOW(), NOW()),
('Title 14', 'Description 14', 0, 23, NOW(), NOW()),
('Title 15', 'Description 15', 1, 24, NOW(), NOW()),
('Title 16', 'Description 16', 0, 25, NOW(), NOW()),
('Title 17', 'Description 17', 1, 22, NOW(), NOW()),
('Title 18', 'Description 18', 0, 23, NOW(), NOW()),
('Title 19', 'Description 19', 1, 24, NOW(), NOW()),
('Title 20', 'Description 20', 0, 25, NOW(), NOW());
