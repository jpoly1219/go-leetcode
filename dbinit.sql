CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE='Asia/Seoul';

CREATE TABLE IF NOT EXISTS users (
    user_id uuid DEFAULT uuid_generate_v4(),
    username VARCHAR(16) UNIQUE NOT NULL,
    fullname VARCHAR(128) NOT NULL,
    email VARCHAR(128) UNIQUE NOT NULL,
    password VARCHAR(128) NOT NULL,
    profile_pic TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS problems (
    problem_id uuid DEFAULT uuid_generate_v4(),
    title VARCHAR(128) UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    difficulty VARCHAR(8) NOT NULL,
    description TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS attempts (
    attempt_id uuid DEFAULT uuid_generate_v4(),
    username VARCHAR(16) UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    lang VARCHAR(8) NOT NULL,
    code TEXT NOT NULL,
    result TEXT NOT NULL,
    output TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_user
        FOREIGN KEY(username)
            REFERENCES users(username)
            ON DELETE CASCADE,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS templates (
    template_id uuid DEFAULT uuid_generate_v4(),
    slug VARCHAR(128) UNIQUE NOT NULL,
    template TEXT UNIQUE NOT NULL,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS testcases (
    testcase_id uuid DEFAULT uuid_generate_v4(),
    testcase TEXT UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS solutions (
    solution_id uuid DEFAULT uuid_generate_v4(),
    slug VARCHAR(128) UNIQUE NOT NULL,
    solution TEXT UNIQUE NOT NULL,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS discussions (
    discussion_id uuid DEFAULT uuid_generate_v4(),
    username VARCHAR(16) UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_user
        FOREIGN KEY(author)
            REFERENCES users(username)
            ON DELETE CASCADE,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id uuid DEFAULT uuid_generate_v4(),
    username VARCHAR(16) UNIQUE NOT NULL,
    discussion_id uuid,
    CONSTRAINT fk_user
        FOREIGN KEY(author)
            REFERENCES users(username)
            ON DELETE CASCADE,
    CONSTRAINT fk_discussion
        FOREIGN KEY(discussion_id)
            REFERENCES discussions(discussion_id)
            ON DELETE CASCADE,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);