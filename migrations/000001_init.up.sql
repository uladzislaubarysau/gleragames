create table requests_history (
    id BIGSERIAL primary key,
    status_code TEXT not null,
    created_at TIMESTAMP default now()
);