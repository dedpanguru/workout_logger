CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name varchar UNIQUE,
  password varchar,
  createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tokens (
  userID int PRIMARY KEY,
  token varchar,
  createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user FOREIGN KEY (userID) REFERENCES users(id)
);

CREATE TABLE workouts(
  userID int PRIMARY KEY,
  id SERIAL UNIQUE,
  createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  day date,
  CONSTRAINT fk_user FOREIGN KEY (userID) REFERENCES users(id)
);

CREATE TABLE exercises (
  exerciseName varchar PRIMARY KEY,
  exerciseType int
);

CREATE TABLE actuals (
  workoutID int PRIMARY KEY,
  exercise varchar, 
  resistance varchar,
  numSets int,
  reps int,
  mins int,
  distance varchar,
  createdAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_workout FOREIGN KEY (workoutID) REFERENCES workouts(id),
  CONSTRAINT fk_exercise FOREIGN KEY (exercise) REFERENCES exercises(exerciseName)
);
