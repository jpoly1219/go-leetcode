CREATE EXTENSION [IF NOT EXISTS] "uuid-ossp";

SET TIMEZONE='KST'

CREATE TABLE [IF NOT EXISTS] attempts (
    attempt_id uuid DEFAULT uuid_generate_v4(),
    username -- foreign key to users,
    slug -- foreign key to problems,
    lang VARCHAR(8) NOT NULL,
    code TEXT NOT NULL,
    result TEXT NOT NULL,
    output TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)

CREATE TABLE [IF NOT EXISTS] templates (
    template_id uuid DEFAULT uuid_generate_v4(),
    slug -- foreign key to problems,
    template TEXT UNIQUE NOT NULL
)

CREATE TABLE [IF NOT EXISTS] testcases (
    testcase_id uuid DEFAULT uuid_generate_v4(),
    slug -- foreign key to problems,
    testcase TEXT UNIQUE NOT NULL
)

CREATE TABLE [IF NOT EXISTS] users (
    user_id uuid DEFAULT uuid_generate_v4(),
    username VARCHAR(16) UNIQUE NOT NULL,
    fullname VARCHAR(128) NOT NULL,
    email VARCHAR(128) UNIQUE NOT NULL,
    password VARCHAR(128) NOT NULL,
    profile_pic TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)

CREATE TABLE [IF NOT EXISTS] problems (
    problem_id uuid DEFAULT uuid_generate_v4(),
    title VARCHAR(128) UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    difficulty VARCHAR(8) NOT NULL,
    description TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)

CREATE TABLE [IF NOT EXISTS] solutions (
    solution_id uuid DEFAULT uuid_generate_v4(),
    slug -- foreign key to problems,
    solution TEXT UNIQUE NOT NULL
)

CREATE TABLE [IF NOT EXISTS] discussions (
    discussion_id uuid DEFAULT uuid_generate_v4(),
    author -- foreign key to users,
    slug -- foreign key to problems,
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)

CREATE TABLE [IF NOT EXISTS] comments (
    comment_id uuid DEFAULT uuid_generate_v4(),
    author -- foreign key to users,
    discussion_id -- foreign key to discussions,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)