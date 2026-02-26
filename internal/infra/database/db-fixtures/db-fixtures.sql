DROP DATABASE IF EXISTS duondb;

CREATE DATABASE duondb;

INSERT INTO usuarios (nome, email, senha_hash) VALUES
('Alice', 'alice@example.com', '1234'),
('Bob', 'bob@example.com', 'abcd');
