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
)
