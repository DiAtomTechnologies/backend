CREATE TABLE CAREERS (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  location VARCHAR(20) CHECK (location IN ('chennai', 'bangalore','remote','online')),
  worktype VARCHAR(20) CHECK (workType IN ('job', 'internship', 'event')),
  description TEXT NOT NULL,
  duration INT DEFAULT 0,
  durationType VARCHAR(20) CHECK (durationType IN ('month', 'months', 'day', 'days', 'week', 'weeks')) ,
  start_date TIMESTAMP WITH TIME ZONE NOT NULL,
  end_date TIMESTAMP WITH TIME ZONE NOT NULL,
  application_time INT NOT NULL
);

CREATE TABLE BLOG(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    imgurl VARCHAR(255) NOT NULL,
    heading VARCHAR(255) NOT NULL,
    tag VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at := CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_blog_timestamp
BEFORE UPDATE ON BLOG
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();
