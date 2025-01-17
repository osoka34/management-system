create table if not exists projects (
    id UUID primary key,
    title varchar(255) not null,
    status_id bigint not null,
    -- 0 - undefined,
    -- 1 - created,
    -- 2 - closed
    
    creator_id UUID not null,
    description text not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

insert into projects (id, title, status_id, creator_id, description, created_at, updated_at)
    values
('36411f72-4393-4490-a89c-3b76d83db7ce', 't2', 2, 'd869a52b-037c-4f92-802e-d935b7dfb56b', 't2', '2025-01-16 16:45:41.294176', '2025-01-16 20:57:44.191451'),
('eea67bc4-ed69-4d48-b40a-2f3bea9d7582', 'Мой первый проект', 1, 'd869a52b-037c-4f92-802e-d935b7dfb56b', 'попробуем сделать что-то интересное, вдруг что-то получится, кто знает', '2025-01-16 16:45:36.153119', '2025-01-16 20:57:47.428470');
