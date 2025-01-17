create table if not exists requirements (
    id UUID primary key,
    title varchar(255) not null,
    status_id bigint not null,
    -- 0 - undefined,
    -- 1 - create,
    -- 2 - closed,
    -- 3 - in progress,
    -- 4 - done,
    description text not null,
    executor_id UUID not null,
    project_id UUID not null,
    specification_id UUID default null,
    created_at timestamp not null,
    updated_at timestamp not null
    );

insert into requirements (id, title, status_id, description, executor_id, project_id, specification_id, created_at, updated_at)
values
    ('85969871-2015-4c2a-9a47-01b7e698558e', 'требование 100', 1, 'какой-то странный деспкрипшен', 'd869a52b-037c-4f92-802e-d935b7dfb56b', 'eea67bc4-ed69-4d48-b40a-2f3bea9d7582', null, '2025-01-16 16:46:18.723070', '2025-01-16 20:11:41.171369'),
    ('08d81ed0-0999-4aaf-b599-66c85a233955', 'требование 1', 1, 'какой-то странный деспкрипшен', '52d23550-0072-48ed-b3a6-aff7eb3521a8', 'eea67bc4-ed69-4d48-b40a-2f3bea9d7582', null, '2025-01-16 19:01:32.696476', '2025-01-16 20:29:01.282124'),
    ('ed7a8343-eafd-4e08-ab70-45f8451adeee', 'новое название вместо 1', 1, 'какой-то странный деспкрипшен', 'bf250fb9-6be8-4639-bb52-caa032b3291c', 'eea67bc4-ed69-4d48-b40a-2f3bea9d7582', '37c54952-78fc-4978-b355-6ee8c23da7ed', '2025-01-16 18:58:45.841098', '2025-01-16 20:40:11.671892'),
    ('43586418-c862-4cfc-8ac5-6345c889d534', 'требование 2', 1, 'какой-то странный деспкрипшен', 'd869a52b-037c-4f92-802e-d935b7dfb56b', 'eea67bc4-ed69-4d48-b40a-2f3bea9d7582', null, '2025-01-16 19:01:47.313801', '2025-01-16 20:49:54.019935'),
    ('f3587497-85e5-4641-9afb-778426c663b6', 'требование 13', 1, 'какой-то странный деспкрипшен', 'd869a52b-037c-4f92-802e-d935b7dfb56b', 'eea67bc4-ed69-4d48-b40a-2f3bea9d7582', '37c54952-78fc-4978-b355-6ee8c23da7ed', '2025-01-16 20:51:27.523715', '2025-01-16 20:53:02.172753'),
    ('26bd21e8-668a-48a8-8ccf-2c3749f4d6e8', 'требование 13', 1, 'какой-то странный деспкрипшен', 'd869a52b-037c-4f92-802e-d935b7dfb56b', 'eea67bc4-ed69-4d48-b40a-2f3bea9d7582', '37c54952-78fc-4978-b355-6ee8c23da7ed', '2025-01-16 21:11:18.454322', '2025-01-16 21:11:43.086693');
