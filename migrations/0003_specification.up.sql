create table if not exists specifications (
    id UUID primary key,
    title varchar(255) not null,
    description text not null,
    status_id bigint not null,
    -- 0 - undefined,
    -- 1 - created,
    -- 2 - closed
    project_id UUID not null,
    created_at timestamp not null,
    updated_at timestamp not null
);


insert into specifications (id, title, description, status_id, project_id, created_at, updated_at)
    values
('37c54952-78fc-4978-b355-6ee8c23da7ed', 'спецификация 1', 'описание тут', 1, 'eea67bc4-ed69-4d48-b40a-2f3bea9d7582', '2025-01-16 19:10:06.830530', '2025-01-16 20:56:35.928839');
