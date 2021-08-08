CREATE TYPE status_enum AS ENUM ('preparing', 'baking', 'done');

CREATE TABLE IF NOT EXISTS bake (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  name TEXT NOT NULL,
  bake_status status_enum
);

CREATE TABLE IF NOT EXISTS toppings (
  id UUID PRIMARY KEY
);

CREATE Table bake_toppings(
  bake_id UUID,
  topping_id UUID,
  PRIMARY KEY (bake_id, topping_id),
  FOREIGN KEY (bake_id) REFERENCES bake(id),
  FOREIGN KEY (topping_id) REFERENCES toppings(id) 
);