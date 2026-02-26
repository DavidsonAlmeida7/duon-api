CREATE EXTENSION "uuid-ossp";

CREATE TABLE usuarios (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nome VARCHAR(120) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    senha_hash VARCHAR(255) NOT NULL,
    criado_em TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_usuarios_email ON usuarios(email);

CREATE TYPE tipo_divisao_enum AS ENUM ('IGUAL', 'PROPORCIONAL');

CREATE TABLE casais (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    usuario1_id UUID NOT NULL REFERENCES usuarios(id),
    usuario2_id UUID NOT NULL REFERENCES usuarios(id),
    tipo_divisao tipo_divisao_enum NOT NULL,
    percentual_usuario1 NUMERIC(5,2),
    percentual_usuario2 NUMERIC(5,2),
    criado_em TIMESTAMP NOT NULL DEFAULT NOW(),
    CHECK (usuario1_id <> usuario2_id)
);

CREATE INDEX idx_casal_usuario1 ON casais(usuario1_id);
CREATE INDEX idx_casal_usuario2 ON casais(usuario2_id);

CREATE TABLE despesas (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    casal_id UUID NOT NULL REFERENCES casais(id) ON DELETE CASCADE,
    pago_por_id UUID NOT NULL REFERENCES usuarios(id),
    valor NUMERIC(12,2) NOT NULL CHECK (valor > 0),
    categoria VARCHAR(80),
    data_despesa DATE NOT NULL,
    criado_em TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_despesas_casal ON despesas(casal_id);
CREATE INDEX idx_despesas_data ON despesas(data_despesa);
CREATE INDEX idx_despesas_casal_data ON despesas(casal_id, data_despesa);

CREATE TABLE metas (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    casal_id UUID NOT NULL REFERENCES casais(id) ON DELETE CASCADE,
    nome VARCHAR(120) NOT NULL,
    valor_objetivo NUMERIC(12,2) NOT NULL CHECK (valor_objetivo > 0),
    valor_atual NUMERIC(12,2) NOT NULL DEFAULT 0 CHECK (valor_atual >= 0),
    data_limite DATE,
    criado_em TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE equidade_mensal (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    casal_id UUID NOT NULL REFERENCES casais(id) ON DELETE CASCADE,
    mes_referencia CHAR(7) NOT NULL,
    percentual_usuario1 NUMERIC(5,2),
    percentual_usuario2 NUMERIC(5,2),
    nivel VARCHAR(50),
    calculado_em TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (casal_id, mes_referencia)
);

-- (OPCIONAL FUTURO)
CREATE TABLE progresso_casal (
    casal_id UUID PRIMARY KEY REFERENCES casais(id) ON DELETE CASCADE,
    pontuacao INT NOT NULL DEFAULT 0 CHECK (pontuacao >= 0),
    nivel_atual VARCHAR(50) NOT NULL,
    atualizado_em TIMESTAMP NOT NULL DEFAULT NOW()
);
