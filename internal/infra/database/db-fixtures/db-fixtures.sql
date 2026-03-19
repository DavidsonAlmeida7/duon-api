DROP DATABASE IF EXISTS duondb;

CREATE DATABASE duondb;

INSERT INTO usuarios (nome_completo, email, whatsapp, data_nascimento, senha_hash) VALUES
('Alice Ferreira', 'alice@example.com', '31987665555', '1234'),
('Bob da Silva', 'bob@example.com', '33987878787', 'abcd');
