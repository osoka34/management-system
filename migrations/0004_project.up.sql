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
)
