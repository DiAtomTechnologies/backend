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
    tag VARCHAR(100) CHECK (tag IN ('TECHNOLOGY', 'INNOVATION')),
    description TEXT NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE USERS (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  email VARCHAR(225) UNIQUE NOT NULL,
  password TEXT NOT NULL,
  role VARCHAR(20) CHECK (role IN ('ADMIN' , 'MANAGER'))
);

CREATE TABLE job_applications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    address TEXT NOT NULL,
    work_experience INT NOT NULL,
    job_id UUID REFERENCES careers(id) ON DELETE CASCADE,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
