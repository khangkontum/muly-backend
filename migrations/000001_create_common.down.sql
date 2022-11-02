CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY NOT NULL DEFAULT gen_random_uuid(),
    username text NOT NULL,
    email text NOT NULL,
    phoneNumber text NOT NULL,
    password text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updatedAt timestamp,
)

CREATE TABLE IF NOT EXISTS Posts (
    id UUID PRIMARY NOT NULL DEFAULT gen_random_uuid(),
    userId text NOT NULL,
    

    FOREIGN KEY userId REFERENCES users(id)
)

CREATE TABLE IF NOT EXISTS Comments(
    id UUID PRIMARY NOT NULL DEFAULT gen_random_uuid(),
    userId text NOT NULL,
    postId text NOT NULL,

    FOREIGN KEY userId REFERENCES users(id)
    FOREIGN KEY postId REFERENCES posts(id) 
)



