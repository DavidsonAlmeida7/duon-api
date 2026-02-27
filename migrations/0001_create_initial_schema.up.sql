CREATE EXTENSION "uuid-ossp";

CREATE TABLE usuarios (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nome VARCHAR(120) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    senha_hash VARCHAR(255) NOT NULL,
    criado_em TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_usuarios_email ON usuarios(email);

-- CREATE TYPE tipo_divisao_enum AS ENUM ('IGUAL', 'PROPORCIONAL');
-- CREATE TYPE tipo_grupo_enum AS ENUM ('INDIVIDUAL', 'CASAL');

CREATE TABLE grupos_financeiros (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tipo ENUM('INDIVIDUAL', 'CASAL', 'FAMILIA') NOT NULL,
    tipo_divisao ENUM('IGUAL', 'PROPORCIONAL') NULL,
    nome VARCHAR(120),
    criado_em TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Se grupo_finaceiro for tipo = individual, percentual_responsabilidade = 100.00
CREATE TABLE grupo_usuarios (
    grupo_id UUID NOT NULL REFERENCES grupos_financeiros(id) ON DELETE CASCADE,
    usuario_id UUID NOT NULL REFERENCES usuarios(id) ON DELETE CASCADE,
    percentual_responsabilidade NUMERIC(5,2) CHECK (percentual_responsabilidade BETWEEN 0 AND 100),
    PRIMARY KEY (grupo_id, usuario_id)
);

CREATE TABLE despesas (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    grupo_id UUID NOT NULL REFERENCES grupos_financeiros(id) ON DELETE CASCADE,
    pago_por_id UUID NOT NULL REFERENCES usuarios(id),
    valor NUMERIC(12,2) NOT NULL CHECK (valor > 0),
    categoria VARCHAR(80),
    data_despesa DATE NOT NULL,
    criado_em TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_despesas_casal ON despesas(grupo_id);
CREATE INDEX idx_despesas_data ON despesas(data_despesa);
CREATE INDEX idx_despesas_grupo_data ON despesas(grupo_id, data_despesa);

CREATE TABLE metas (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    grupo_id UUID NOT NULL REFERENCES grupos_financeiros(id) ON DELETE CASCADE,
    nome VARCHAR(120) NOT NULL,
    valor_objetivo NUMERIC(12,2) NOT NULL CHECK (valor_objetivo > 0),
    valor_atual NUMERIC(12,2) NOT NULL DEFAULT 0 CHECK (valor_atual >= 0),
    data_limite DATE,
    criado_em TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE equidade_mensal (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    grupo_id UUID NOT NULL REFERENCES grupos_financeiros(id) ON DELETE CASCADE,
    mes_referencia CHAR(7) NOT NULL,
    percentual_usuario1 NUMERIC(5,2),
    percentual_usuario2 NUMERIC(5,2),
    nivel VARCHAR(50),
    calculado_em TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (grupo_id, mes_referencia)
);

-- (OPCIONAL FUTURO)
CREATE TABLE progresso_grupo (
    grupo_id UUID PRIMARY KEY REFERENCES grupos_financeiros(id) ON DELETE CASCADE,
    pontuacao INT NOT NULL DEFAULT 0 CHECK (pontuacao >= 0),
    nivel_atual VARCHAR(50) NOT NULL,
    atualizado_em TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE historico_nivel_grupo (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    grupo_id UUID NOT NULL REFERENCES grupos_financeiros(id) ON DELETE CASCADE,
    nivel VARCHAR(50) NOT NULL,
    atingido_em TIMESTAMP NOT NULL DEFAULT NOW()
);