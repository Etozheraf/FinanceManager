CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id                        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login                     varchar(255) null,
    budget                    integer      not null,
    current_balance           integer      not null,
    calculate_expiration_date time         not null
);

CREATE TABLE IF NOT EXISTS transactions
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sum INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS user_transaction(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) NOT NULL,
    transaction_id UUID REFERENCES transactions(id) NOT NULL
);