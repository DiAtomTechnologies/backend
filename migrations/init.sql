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
