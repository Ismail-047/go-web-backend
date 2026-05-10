-- 1. Add the column with a completely unusable, fake hash as the default
ALTER TABLE users 
ADD COLUMN password VARCHAR(60) NOT NULL DEFAULT 'unusable_fake_hash_needs_password_reset';

-- 2. Immediately drop the default behavior
ALTER TABLE users 
ALTER COLUMN password DROP DEFAULT;