create table if not exists users (
    id UUID primary key,
    login varchar(255) not null,
    password_hash varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

insert into users (id, login, password_hash, created_at, updated_at)
    values
('d869a52b-037c-4f92-802e-d935b7dfb56b', 'test', '36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80', '2025-01-16 16:45:29.624003', '2025-01-16 16:45:29.624003');
